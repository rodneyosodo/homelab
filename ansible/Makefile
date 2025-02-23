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

.PHONY: setup-tana
setup-tana:
	ansible-playbook playbooks/tana.yaml

.PHONY: help
help:
	@echo "This Makefile provides a set of commands to manage ansible services."
	@echo "It allows you to setup, install, and manage various services on your ansible infrastructure."
	@echo ""
	@echo "Usage:"
	@echo "  make install       - Install ansible roles"
	@echo "  make ping          - Ping all hosts"
	@echo "  make setup-odin    - Setup odin"
	@echo "  make setup-bohr    - Setup bohr"
	@echo "  make install-zsh   - Install zsh"
	@echo "  make setup-homelab - Setup homelab"
	@echo "  make setup-tana    - Setup tana"
