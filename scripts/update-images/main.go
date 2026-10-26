// Command update-images bumps container image tags in docker-compose files,
// Dependabot style.
//
// For every `image: registry/repo:tag` line it looks up the registry's tags and
// picks the highest version with the same shape as the current tag (same
// prefix such as `v`, same number of version components, same suffix such as
// `-alpine`). Like Renovate's minimumReleaseAge, a candidate is only accepted
// once it has been published for at least -min-age days, so a freshly pushed
// (and possibly compromised) release is never pulled in straight away.
//
// Release age comes from Docker Hub's tag_last_pushed for docker.io images and
// from the image config's created timestamp everywhere else. A tag whose age
// cannot be determined is never selected.
//
// Dependencies listed under `ignore:` in .github/dependabot.yaml are skipped,
// so both tools share one ignore list.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	imageLine      = regexp.MustCompile(`^(\s*image:\s*["']?)([^"'\s#]+)(.*)$`)
	tagShape       = regexp.MustCompile(`^([A-Za-z]*)(\d+(?:\.\d+)*)(.*)$`)
	ignoreName     = regexp.MustCompile(`dependency-name:\s*["']?([^"'\s]+)`)
	bearerParam    = regexp.MustCompile(`(\w+)="([^"]*)"`)
	linkNext       = regexp.MustCompile(`<([^>]+)>;\s*rel="next"`)
	skipTags       = map[string]bool{"latest": true, "stable": true, "main": true, "master": true, "nightly": true, "edge": true, "dev": true}
	minTrustedDate = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // reproducible builds pin created to the epoch
	manifestAccept = strings.Join([]string{
		"application/vnd.oci.image.index.v1+json",
		"application/vnd.docker.distribution.manifest.list.v2+json",
		"application/vnd.oci.image.manifest.v1+json",
		"application/vnd.docker.distribution.manifest.v2+json",
	}, ", ")
)

// sem bounds the number of concurrent HTTP requests across all goroutines.
var sem chan struct{}

var client = &http.Client{Timeout: 60 * time.Second}

type version struct {
	prefix string
	parts  []int
	suffix string
}

func parseTag(tag string) (version, bool) {
	m := tagShape.FindStringSubmatch(tag)
	if m == nil {
		return version{}, false
	}
	var parts []int
	for _, p := range strings.Split(m[2], ".") {
		n, err := strconv.Atoi(p)
		if err != nil {
			return version{}, false
		}
		parts = append(parts, n)
	}
	return version{m[1], parts, m[3]}, true
}

func (v version) sameShape(o version) bool {
	return v.prefix == o.prefix && v.suffix == o.suffix && len(v.parts) == len(o.parts)
}

func (v version) less(o version) bool {
	for i := range v.parts {
		if v.parts[i] != o.parts[i] {
			return v.parts[i] < o.parts[i]
		}
	}
	return false
}

type occurrence struct {
	path string
	line int
}

type result struct {
	name, tag, newTag, note string
	err                     error
}

// get performs a rate-limited GET, retrying on 429 and 5xx responses.
func get(rawURL string, header http.Header) (*http.Response, []byte, error) {
	sem <- struct{}{}
	defer func() { <-sem }()

	var lastErr error
	for attempt := 0; attempt < 4; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(math.Pow(2, float64(attempt))) * time.Second)
		}
		req, err := http.NewRequest(http.MethodGet, rawURL, nil)
		if err != nil {
			return nil, nil, err
		}
		for k, v := range header {
			req.Header[k] = v
		}
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if cerr := resp.Body.Close(); err == nil {
			err = cerr
		}
		if err != nil {
			lastErr = err
			continue
		}
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("GET %s: %s", rawURL, resp.Status)
			if secs, err := strconv.Atoi(resp.Header.Get("Retry-After")); err == nil && secs > 0 && secs <= 60 {
				time.Sleep(time.Duration(secs) * time.Second)
			}
			continue
		}
		return resp, body, nil
	}
	return nil, nil, lastErr
}

func getJSON(rawURL string, header http.Header, out any) (*http.Response, error) {
	resp, body, err := get(rawURL, header)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("GET %s: %s", rawURL, resp.Status)
	}
	return resp, json.Unmarshal(body, out)
}

func parseTime(s string) (time.Time, bool) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil || t.Before(minTrustedDate) {
		return time.Time{}, false
	}
	return t, true
}

// splitName returns the registry host and repository path of an image name.
func splitName(name string) (string, string) {
	first, rest, found := strings.Cut(name, "/")
	if !found || (!strings.ContainsAny(first, ".:") && first != "localhost") {
		return "docker.io", name
	}
	return first, rest
}

// tagLister returns the tags of a repository and a function resolving a tag's
// release date. A zero time means the date is unknown.
type tagLister interface {
	tags(current string) ([]string, error)
	released(tag string) time.Time
}

// dockerHub uses the Hub API, which returns push dates alongside the tags.
// Anonymous clients may only read the first 10 pages, so tags are fetched
// newest first and paging stops once the current tag is seen: anything newer
// was pushed after it.
type dockerHub struct {
	repo  string
	dates map[string]time.Time
}

type hubPage struct {
	Next    string `json:"next"`
	Results []struct {
		Name          string `json:"name"`
		TagLastPushed string `json:"tag_last_pushed"`
		LastUpdated   string `json:"last_updated"`
	} `json:"results"`
}

func (h *dockerHub) tags(current string) ([]string, error) {
	const maxPages = 10
	h.dates = map[string]time.Time{}
	var tags []string
	next := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags?page_size=100&ordering=last_updated", h.repo)
	for page := 0; next != "" && page < maxPages; page++ {
		var p hubPage
		if _, err := getJSON(next, nil, &p); err != nil {
			return nil, err
		}
		next = p.Next
		for _, r := range p.Results {
			tags = append(tags, r.Name)
			if t, ok := parseTime(r.TagLastPushed); ok {
				h.dates[r.Name] = t
			} else if t, ok := parseTime(r.LastUpdated); ok {
				h.dates[r.Name] = t
			}
			if r.Name == current {
				next = ""
			}
		}
	}
	return tags, nil
}

func (h *dockerHub) released(tag string) time.Time {
	return h.dates[tag]
}

// registry speaks the OCI distribution API (ghcr.io, quay.io, lscr.io, ...).
type registry struct {
	host, repo string
	once       sync.Once
	auth       http.Header
	authErr    error
}

func (r *registry) header(accept string) (http.Header, error) {
	r.once.Do(func() { r.auth, r.authErr = r.token() })
	if r.authErr != nil {
		return nil, r.authErr
	}
	h := r.auth.Clone()
	if accept != "" {
		h.Set("Accept", accept)
	}
	return h, nil
}

// token performs the anonymous bearer token handshake if the registry asks for it.
func (r *registry) token() (http.Header, error) {
	resp, _, err := get("https://"+r.host+"/v2/", nil)
	if err != nil {
		return nil, err
	}
	h := http.Header{}
	challenge := resp.Header.Get("WWW-Authenticate")
	if resp.StatusCode != http.StatusUnauthorized || !strings.HasPrefix(strings.ToLower(challenge), "bearer ") {
		return h, nil
	}
	params := map[string]string{}
	for _, m := range bearerParam.FindAllStringSubmatch(challenge, -1) {
		params[m[1]] = m[2]
	}
	q := url.Values{"scope": {"repository:" + r.repo + ":pull"}}
	if s := params["service"]; s != "" {
		q.Set("service", s)
	}
	var tok struct {
		Token       string `json:"token"`
		AccessToken string `json:"access_token"`
	}
	if _, err := getJSON(params["realm"]+"?"+q.Encode(), nil, &tok); err != nil {
		return nil, err
	}
	if tok.Token == "" {
		tok.Token = tok.AccessToken
	}
	h.Set("Authorization", "Bearer "+tok.Token)
	return h, nil
}

// tags lists the repository's tags. Some registries (ghcr.io) page in push
// order, so a `last=<current>` cursor returns only tags pushed after the
// current one, skipping thousands of old CI tags. Registries that page in
// lexical order would skip newer versions such as v1.10 after v1.9, so when
// the cursor page comes back empty or sorted the full listing is used instead.
func (r *registry) tags(current string) ([]string, error) {
	h, err := r.header("")
	if err != nil {
		return nil, err
	}
	base := fmt.Sprintf("https://%s/v2/%s/tags/list?n=10000", r.host, r.repo)
	if tags, err := r.list(base+"&last="+url.QueryEscape(current), h, true); err == nil && tags != nil {
		return tags, nil
	}
	return r.list(base, h, false)
}

// list follows the Link pagination from start. With pushOrder set it gives up
// (returning nil) unless the first page is non-empty and not lexically sorted.
func (r *registry) list(start string, h http.Header, pushOrder bool) ([]string, error) {
	var tags []string
	next := start
	for first := true; next != ""; first = false {
		var page struct {
			Tags []string `json:"tags"`
		}
		resp, err := getJSON(next, h, &page)
		if err != nil {
			return nil, err
		}
		if first && pushOrder && (len(page.Tags) < 2 || sort.StringsAreSorted(page.Tags)) {
			return nil, nil
		}
		tags = append(tags, page.Tags...)
		next = ""
		if m := linkNext.FindStringSubmatch(resp.Header.Get("Link")); m != nil {
			u, err := resp.Request.URL.Parse(m[1])
			if err != nil {
				return nil, err
			}
			next = u.String()
		}
	}
	return tags, nil
}

type manifest struct {
	Manifests []struct {
		Digest   string `json:"digest"`
		Platform struct {
			OS           string `json:"os"`
			Architecture string `json:"architecture"`
		} `json:"platform"`
	} `json:"manifests"`
	Config struct {
		Digest string `json:"digest"`
	} `json:"config"`
}

func (r *registry) released(tag string) time.Time {
	h, err := r.header(manifestAccept)
	if err != nil {
		return time.Time{}
	}
	base := fmt.Sprintf("https://%s/v2/%s/", r.host, r.repo)

	var m manifest
	if _, err := getJSON(base+"manifests/"+tag, h, &m); err != nil {
		return time.Time{}
	}
	if len(m.Manifests) > 0 {
		digest := ""
		for _, d := range m.Manifests {
			if d.Platform.OS == "linux" && d.Platform.Architecture == "amd64" {
				digest = d.Digest
				break
			}
			if digest == "" && d.Platform.OS != "unknown" {
				digest = d.Digest
			}
		}
		if digest == "" {
			return time.Time{}
		}
		m = manifest{}
		if _, err := getJSON(base+"manifests/"+digest, h, &m); err != nil {
			return time.Time{}
		}
	}
	if m.Config.Digest == "" {
		return time.Time{}
	}

	var cfg struct {
		Created string `json:"created"`
		Config  struct {
			Labels map[string]string `json:"Labels"`
		} `json:"config"`
	}
	if _, err := getJSON(base+"blobs/"+m.Config.Digest, h, &cfg); err != nil {
		return time.Time{}
	}
	if t, ok := parseTime(cfg.Created); ok {
		return t
	}
	t, _ := parseTime(cfg.Config.Labels["org.opencontainers.image.created"])
	return t
}

func newLister(name string) tagLister {
	host, repo := splitName(name)
	switch host {
	case "docker.io", "index.docker.io", "registry-1.docker.io":
		if !strings.Contains(repo, "/") {
			repo = "library/" + repo
		}
		return &dockerHub{repo: repo}
	}
	return &registry{host: host, repo: repo}
}

func resolve(name, tag string, minAge time.Duration, now time.Time) result {
	res := result{name: name, tag: tag}
	current, ok := parseTag(tag)
	if !ok {
		res.note = "unversioned tag"
		return res
	}

	lister := newLister(name)
	tags, err := lister.tags(tag)
	if err != nil {
		res.err = fmt.Errorf("cannot list tags: %w", err)
		return res
	}

	type candidate struct {
		tag  string
		v    version
		date time.Time
	}
	var candidates []*candidate
	for _, t := range tags {
		v, ok := parseTag(t)
		if ok && v.sameShape(current) && current.less(v) {
			candidates = append(candidates, &candidate{tag: t, v: v})
		}
	}
	if len(candidates) == 0 {
		return res
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[j].v.less(candidates[i].v) })

	var wg sync.WaitGroup
	for _, c := range candidates {
		wg.Add(1)
		go func(c *candidate) {
			defer wg.Done()
			c.date = lister.released(c.tag)
		}(c)
	}
	wg.Wait()

	var held []string
	for _, c := range candidates {
		switch {
		case c.date.IsZero():
			held = append(held, c.tag+" (unknown age)")
		case now.Sub(c.date) < minAge:
			held = append(held, fmt.Sprintf("%s (%dd old)", c.tag, int(now.Sub(c.date).Hours()/24)))
		default:
			res.newTag = c.tag
			if len(held) > 0 {
				res.note = "held back: " + strings.Join(held, ", ")
			}
			return res
		}
	}
	res.note = "held back: " + strings.Join(held, ", ")
	return res
}

func splitList(s string) map[string]bool {
	set := map[string]bool{}
	for _, item := range strings.Split(s, ",") {
		if item = strings.TrimSpace(item); item != "" {
			set[item] = true
		}
	}
	return set
}

func dependabotIgnores(root string) map[string]bool {
	set := map[string]bool{}
	for _, name := range []string{"dependabot.yaml", "dependabot.yml"} {
		data, err := os.ReadFile(filepath.Join(root, ".github", name))
		if err != nil {
			continue
		}
		for _, m := range ignoreName.FindAllStringSubmatch(string(data), -1) {
			set[m[1]] = true
		}
	}
	return set
}

func composeFiles(dir string) ([]string, error) {
	var files []string
	for _, pattern := range []string{"*compose*.y*ml", "*/*compose*.y*ml"} {
		matches, err := filepath.Glob(filepath.Join(dir, pattern))
		if err != nil {
			return nil, err
		}
		files = append(files, matches...)
	}
	sort.Strings(files)
	return files, nil
}

func main() {
	root := flag.String("root", ".", "repository root")
	minAgeDays := flag.Int("min-age", 7, "minimum release age in days")
	dryRun := flag.Bool("dry-run", false, "only report, do not modify files")
	ignore := flag.String("ignore", "", "extra comma separated image names to skip")
	only := flag.String("only", "", "comma separated image names to check (default: all)")
	jobs := flag.Int("jobs", 32, "maximum concurrent HTTP requests")
	flag.Parse()

	sem = make(chan struct{}, *jobs)
	ignores := dependabotIgnores(*root)
	for name := range splitList(*ignore) {
		ignores[name] = true
	}
	onlySet := splitList(*only)

	files, err := composeFiles(filepath.Join(*root, "docker-compose"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	type ref struct{ name, tag string }
	occurrences := map[ref][]occurrence{}
	contents := map[string][]string{}
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		lines := strings.SplitAfter(string(data), "\n")
		contents[path] = lines
		for i, line := range lines {
			m := imageLine.FindStringSubmatch(strings.TrimRight(line, "\r\n"))
			if m == nil || strings.Contains(m[2], "@") {
				continue
			}
			idx := strings.LastIndex(m[2], ":")
			if idx < 0 || strings.Contains(m[2][idx+1:], "/") {
				continue
			}
			r := ref{m[2][:idx], m[2][idx+1:]}
			if ignores[r.name] || (len(onlySet) > 0 && !onlySet[r.name]) || skipTags[r.tag] {
				continue
			}
			occurrences[r] = append(occurrences[r], occurrence{path, i})
		}
	}

	now := time.Now().UTC()
	minAge := time.Duration(*minAgeDays) * 24 * time.Hour
	results := make(chan result, len(occurrences))
	var wg sync.WaitGroup
	for r := range occurrences {
		wg.Add(1)
		go func(r ref) {
			defer wg.Done()
			results <- resolve(r.name, r.tag, minAge, now)
		}(r)
	}
	wg.Wait()
	close(results)

	var sorted []result
	for res := range results {
		sorted = append(sorted, res)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].name != sorted[j].name {
			return sorted[i].name < sorted[j].name
		}
		return sorted[i].tag < sorted[j].tag
	})

	changed := map[string]bool{}
	updates, failures := 0, 0
	for _, res := range sorted {
		switch {
		case res.err != nil:
			failures++
			fmt.Fprintf(os.Stderr, "  error   %s:%s  [%v]\n", res.name, res.tag, res.err)
		case res.newTag != "":
			note := ""
			if res.note != "" {
				note = "  [" + res.note + "]"
			}
			fmt.Fprintf(os.Stderr, "  update  %s: %s -> %s%s\n", res.name, res.tag, res.newTag, note)
			for _, o := range occurrences[ref{res.name, res.tag}] {
				old, repl := res.name+":"+res.tag, res.name+":"+res.newTag
				contents[o.path][o.line] = strings.Replace(contents[o.path][o.line], old, repl, 1)
				changed[o.path] = true
				updates++
			}
		case res.note != "":
			fmt.Fprintf(os.Stderr, "  skip    %s:%s  [%s]\n", res.name, res.tag, res.note)
		}
	}

	switch {
	case updates == 0:
		fmt.Fprintln(os.Stderr, "No image updates available.")
	case *dryRun:
		fmt.Fprintf(os.Stderr, "Dry run: %d image line(s) would change.\n", updates)
	default:
		for path := range changed {
			if err := os.WriteFile(path, []byte(strings.Join(contents[path], "")), 0o644); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "wrote %s\n", path)
		}
	}
	if failures > 0 {
		os.Exit(1)
	}
}
