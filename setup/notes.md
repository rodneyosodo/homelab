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

Use dotfiles found [here](https://github.com/rodneyosodo/dotfiles/tree/main/config)

```bash
mkdir ~/.ssh
```

copy ssh key

```bash
scp ~/.ssh/github rodneyosodo@bohr:/home/rodneyosodo/.ssh/
scp ~/.ssh/github.pub rodneyosodo@bohr:/home/rodneyosodo/.ssh/
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

## Setup Syncthing On Remote Host

Port forward to localhost:1111

```bash
ssh rodneyosodo@tana -L 1111:localhost:8384
```

### Clear disk storage

```bash
fdisk /dev/sda
```

```bash
 Command (m for help): p
 Command (m for help): d
 Command (m for help): w
```

## Passthrough USB nic to VM:

```bash
usb-devices
```

```bash
T:  Bus=04 Lev=01 Prnt=01 Port=00 Cnt=01 Dev#=  5 Spd=5000 MxCh= 0
D:  Ver= 3.00 Cls=00(>ifc ) Sub=00 Prot=00 MxPS= 9 #Cfgs=  2
P:  Vendor=0bda ProdID=8153 Rev=31.00
S:  Manufacturer=Realtek
S:  Product=USB 10/100/1000 LAN
S:  SerialNumber=001000001
C:  #Ifs= 1 Cfg#= 1 Atr=a0 MxPwr=288mA
I:  If#= 0 Alt= 0 #EPs= 3 Cls=ff(vend.) Sub=ff Prot=00 Driver=r8152
E:  Ad=02(O) Atr=02(Bulk) MxPS=1024 Ivl=0ms
E:  Ad=81(I) Atr=02(Bulk) MxPS=1024 Ivl=0ms
E:  Ad=83(I) Atr=03(Int.) MxPS=   2 Ivl=16ms
```

```bash
qm set 110 -usb0 host=0bda:8153,usb3=yes
```

## Create a new user and add it to the sudo group

```bash
adduser rodneyosodo
usermod -aG sudo rodneyosodo
su - rodneyosodo
```
