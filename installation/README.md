# Installation

## Install softwares

```bash
sudo apt install git vim htop neofetch p7zip unrar tar flac curl wget make thefuck python-pip
```

## Install bpytop

```bash
pip install bpytop --break-system-packages
```

## Install docker

```bash
sudo bash -c "$(wget -qLO - https://get.docker.com)"
```

post installation

```bash
sudo groupadd docker
sudo usermod -aG docker $USER
sudo systemctl enable docker.service
sudo systemctl enable containerd.service
```

## Install vscode server

```bash
curl -fsSL https://code-server.dev/install.sh | sh
```
