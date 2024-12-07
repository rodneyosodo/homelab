# Installation

## Award Superuser rights

```bash
su - root
```

```bash
usermod -aG sudo rodneyosodo
```

## Disable DVD/ISO CD-ROM Package Repository

This is on debian

```bash
nano /etc/apt/sources.list
```

comment line

```list
# deb cdrom:[Debian GNU/Linux 12.4.0 _Bookworm_ - Official amd64 DVD Binary-1 with firmware 20231210-17:57]/ bookworm main non-free-firmware
```

## Install software

```bash
sudo nala install ssh openssh-server git vim bpytop p7zip tar curl wget make thefuck
```

## Install docker

```bash
sudo bash -c "$(wget -qLO - https://get.docker.com)"
```

post installation

```bash
sudo usermod -aG docker $USER
sudo systemctl enable docker.service
sudo systemctl enable containerd.service
```

## Disable GRUB delay

```bash
sudo vim /etc/default/grub
sudo update-grub
```

## Setup git

```bash
git config --global user.email "socials@rodneyosodo.com"
git config --global user.name "Rodney Osodo"
git config --global push.autoSetupRemote true
git config --global init.defaultBranch main
git config --global commit.gpgsign true
git config --global gpg.format ssh
git config --global core.editor nvim
```

```bash
mkdir ~/.ssh
```

copy ssh key

```bash
scp ~/.ssh/github rodneyosodo@192.168.100.32:/home/rodneyosodo/.ssh/
scp ~/.ssh/github.pub rodneyosodo@192.168.100.32:/home/rodneyosodo/.ssh/
```

```bash
eval `ssh-agent -s`
ssh-add ~/.ssh/github
ssh -vT git@github.com
```

## Setup SSH

```bash
sudo systemctl start sshd
sudo systemctl enable sshd
```

copy SSH key (GitHub and personal)

```bash
scp ~/.ssh/homelab-pc.pub rodneyosodo@192.168.100.32:/home/rodneyosodo/.ssh/
```

on the remote system add the key to the authorized keys

```bash
cat ~/.ssh/homelab-pc.pub >> ~/.ssh/authorized_keys
```

Backup config

```bash
sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.bak
```

Edit config

```bash
sudo vim /etc/ssh/sshd_config
```

- Enable Protocol 2

  ```bash
  Protocol 2
  ```

- Change Port

  ```bash
  Port 365
  ```

- Disable root login

  ```bash
  PermitRootLogin no
  ```

- Use public key authentication

  ```bash
  PubkeyAuthentication yes
  Password Authentication no
  ```

- Disable empty passwords

  ```bash
  PermitEmptyPasswords no
  ```

```bash
sudo systemctl restart sshd
```

```bash
ssh -2 rodneyosodo@192.168.100.32
```

## Install tailscale

```bash
curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up --ssh
```
