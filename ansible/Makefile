.PHONY: install
install:
	ansible-galaxy role install ctorgalson.nerdfonts

.PHONY: ping
ping:
	ansible-playbook playbooks/ping.yaml

.PHONY: setup-odin
setup-odin:
	ansible-playbook playbooks/odin.yaml

.PHONY: setup-bohr
setup-bohr:
	ansible-playbook playbooks/bohr.yaml

.PHONY: install-zsh
install-zsh:
	ansible-playbook playbooks/zsh.yaml

.PHONY: setup-homelab
setup-homelab:
	ansible-playbook playbooks/homelab.yaml
