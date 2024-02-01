# Rodney Osodo's Homelab

Hello, I'm [Rodney Osodo][website] [@blackd0t][twitter] on Twitter. I'm a software engineer and homelab. I use it to learn new technologies and to host myprojects and other services. This repository contains the configuration files for my homelab.

## Goals

- [ ] De-google my life.
- [ ] To eventually provide a highly available and scalable infrastructure for my projects and services with no single point of failure.
- [x] Learn new technologies.
- [x] Have fun.

## Hardware

I have a single server running Proxmox VE 6.5.11-7-pve with the following specs:

- CPU: AMD Ryzen 7 4800H with Radeon Graphics (16) @ 1.4GHz - 2.9GHz
- RAM: 32 GB DDR4-3200 Memory
- SSD: 1x 1 TB NVMe SSD
- Disk: 1x 2 TB SATA SSD

I have a 10 Mbps uplink and a 10 Mbps downlink internet connection from Safaricom (planning to upgrade depending on upload usage). I don't have a static IP address so I use [cloudflare-tunnel][cft] to expose my services to the internet.

I also have a Raspberry Pi 4 Model B with 4 GB RAM which is not running since I have a few issues with running proxmox backup server on it. I had a backup 4TB HDD connected to it but it had 4096 bytes per sector and proxmox backup server only supports 512 bytes per sector. I'm planning to get a new HDD and try again with a mini PC instead of the Raspberry Pi. Another issue is that the proxmox backup server doesn't support ARM64 architecture.

Some photos of my homelab:

![computers][computers]

![odin-server][odin-image]

More photos can be found in this [Immich-photo-album][Immich-photo-album].

## Software

I use Proxmox VE as my hypervisor. I currently have 1 VM running Debian 12 (Bookworm) which is my main server. Based on some viewpoints from the [Linux Unplugged][linux-unplugged] podcast, I'm planning to move to Debian as my main OS because proxmos is based on Debian and I can use the same tools on both the host and the VMs or use Arch Linux as my main OS because of the rolling release model because I use it on my workstation. I'm still undecided on which one to use.

I use ZFS for my storage. I have a single pool made up of 1x 2 TB SSD. I'm planning to add another pool for my backups. Currently, backups are stored on the SSD which hosts the OS.

![proxmox-dashboard][proxmox-dashboard]

## Services

I use the following services:

![services][services]

### portainer

Portainer is a lightweight management UI that allows me to easily manage my different Docker environments (Docker hosts or containers). It is meant to be as simple to deploy as it is to use.

My portainer configuration files can be found in the [portainer-compose][portainer-compose] directory.

![portainer dashboard][portainer-dashboard]

### uptime-kuma

uptime-kuma is a fancy self-hosted monitoring tool. It monitors the uptime of my websites and the status of my services. It uses a beautiful dashboard to display the status of your monitored sites.

My uptime-kuma configuration files can be found in the [uptime-kuma-compose][uptime-kuma-compose] directory.

![uptime-kuma dashboard][uptime-kuma-dashboard]

### heimdall

Heimdall is a dashboard for all my web applications. It is a way to organize all the applications that I use daily.

My heimdall configuration files can be found in the [heimdall-compose][heimdall-compose] directory.

![heimdall dashboard][heimdall-dashboard]

### postgres

Postgres is a relational database management system. I use it to store data for my applications. I use it for my [nextcloud][nextcloud-compose].

My postgres configuration files can be found in the [postgres-compose][postgres-compose] directory.

### nextcloud

Nextcloud is a suite of client-server software for creating and using file hosting services. It is functionally similar to Dropbox, although Nextcloud is free and open-source, allowing anyone to install and operate it on a private server.

My nextcloud configuration files can be found in the [nextcloud-compose][nextcloud-compose] directory.

![nextcloud dashboard][nextcloud-dashboard]

### littlelink

Littlelink is a lightweight DIY alternative to services like Linktree, Retriever, and Linkin.bio. It is a single page that you can host on your server listing all your important links.

My littlelink configuration files can be found in the [littlelink-compose][littlelink-compose] directory.

### ntp

NTP is a networking protocol for clock synchronization between computer systems over packet-switched, variable-latency data networks. In operation since before 1985, NTP is one of the oldest Internet protocols in current use. I use it to synchronize the time on my local network.

My ntp configuration files can be found in the [ntp-compose][ntp-compose] directory.

### cloudflared

Cloudflared is a lightweight tunnel daemon that proxies any localhost HTTP traffic through the Cloudflare network. I use it to expose my services to the internet.

My cloudflared configuration files can be found in the [cloudflared-compose][cloudflared-compose] directory.

### pihole

Pi-hole is a Linux network-level advertisement and Internet tracker blocking application that acts as a DNS sinkhole and optionally a DHCP server, intended for use on a private network. I use it to block ads on my local network. I have running alongside unbound to provide DNS. Unbound is a validating, recursive, and caching DNS resolver. I use it to provide DNS resolution for my services.

My pihole configuration files can be found in the [pihole-compose][pihole-compose] directory.

![pihole dashboard][pihole-dashboard]

### swagger-editor

Swagger Editor lets you edit OpenAPI specifications in YAML inside your browser and preview documentation in real-time. I use it to edit my OpenAPI specifications. Since I am a backend developer, I use it to document my APIs.

My swagger-editor configuration files can be found in the [swagger-editor][swagger-editor] directory.

### immich

Immich is a photo album that I use to store all my photos. I use it to store all my photos. Immich is divided into several services, which are run as individual docker containers.

1. `immich-server` - Handle and respond to REST API requests
2. `immich-microservices` - Execute background jobs (thumbnail generation, metadata extraction, transcoding, etc.)
3. `immich-machine-learning` - Execute machine-learning models
4. `postgres` - Persistent data storage
5. `redis`- Queue management for immich-microservices

My immich configuration files can be found in the [immich-compose][immich-compose] directory.

![immich dashboard][immich-dashboard]

### redis

Redis is an in-memory data structure store, used as a distributed, in-memory key–value database, cache and message broker, with optional durability. I use it to store data for my applications. I use it for my [immich][immich-compose] application.

My redis configuration files can be found in the [redis-compose][redis-compose] directory.

### vaultwarden

Vaultwarden is a lightweight implementation of the Bitwarden API, written in Rust, supports password management and generation, and can be self-hosted. I use it to store my passwords.

My vaultwarden configuration files can be found in the [vaultwarden-compose][vaultwarden-compose] directory.

### opengist

OpenGist is a lightweight implementation of the GitHub Gist API, written in Rust, supports creating, editing, deleting, and listing gists, and can be self-hosted. I use it to store my code snippets.

My opengist configuration files can be found in the [opengist-compose][opengist-compose] directory.

### speedtest-tracker

Speedtest Tracker is a self-hosted internet performance tracking application that runs speedtest checks against Ookla's Speedtest service. I use it to track my internet speeds.

My speedtest-tracker configuration files can be found in the [speedtest-tracker-compose][speedtest-tracker-compose] directory.

![speedtest-tracker dashboard][speedtest-tracker-dashboard]

### kavita

Kavita is a fast, feature rich, cross platform reading server. Built with a focus for manga and the goal of being a full solution for all your reading needs. Setup your own server and share your reading collection with your friends and family.

My kavita configuration files can be found in the [kavita-compose][kavita-compose] directory.

## endlessh

endlessh is an SSH tarpit that slowly sends an endless banner. It keeps SSH clients locked up for hours or even days at a time. The purpose is to put your real SSH server on another port and then let the script kiddies get stuck in this tarpit instead of bothering a real server.

My endlessh configuration files can be found in the [endlessh-compose][endlessh-compose] directory.

## dozzle

dozzle is a simple container log viewer for Docker. It is designed to be very minimalistic and fit into the Docker ecosystem without extra dependencies.

My dozzle configuration files can be found in the [dozzle-compose][dozzle-compose] directory.

![dozzle dashboard][dozzle-dashboard]

[website]: https://rodneyosodo.com
[twitter]: https://twitter.com/b1ackd0t
[cft]: https://www.cloudflare.com/products/tunnel/
[odin-image]: ./assets/odin.jpg
[computers]: ./assets/computers.jpg
[Immich-photo-album]: https://immich.rodneyosodo.com/share/dgJE3wNLnS0ntsFlABuRwvkzCGHZeMBueTNo2NmtVKT-3mM1SYaK--p-ENXRGxph0oY
[linux-unplugged]: https://linuxunplugged.com/
[proxmox-dashboard]: ./assets/proxmox.png
[services]: ./assets/homelab-arch.png
[portainer-compose]: ./docker-compose/portainer/
[portainer-dashboard]: ./assets/portainer.png
[uptime-kuma-compose]: ./docker-compose/uptime-kuma/
[uptime-kuma-dashboard]: ./assets/uptime-kuma.png
[heimdall-compose]: ./docker-compose/heimdall/
[heimdall-dashboard]: ./assets/heimdall.png
[postgres-compose]: ./docker-compose/postgres/
[nextcloud-compose]: ./docker-compose/nextcloud/
[nextcloud-dashboard]: ./assets/nextcloud.png
[littlelink-compose]: ./docker-compose/littlelink/
[ntp-compose]: ./docker-compose/ntp/
[cloudflared-compose]: ./docker-compose/cloudflared/
[pihole-compose]: ./docker-compose/pihole/
[pihole-dashboard]: ./assets/pihole.png
[swagger-editor]: ./docker-compose/swagger-editor/
[immich-compose]: ./docker-compose/immich/
[immich-dashboard]: ./assets/immich.png
[redis-compose]: ./docker-compose/redis/
[vaultwarden-compose]: ./docker-compose/vaultwarden/
[opengist-compose]: ./docker-compose/opengist/
[speedtest-tracker-compose]: ./docker-compose/speedtest-tracker/
[speedtest-tracker-dashboard]: ./assets/speedtest.png
[kavita-compose]: ./docker-compose/kavita/
[endlessh-compose]: ./docker-compose/endlessh/
[dozzle-compose]: ./docker-compose/dozzle/
[dozzle-dashboard]: ./assets/dozzle.png
