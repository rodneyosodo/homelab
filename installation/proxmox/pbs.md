# Proxmox Backup Server Installation

## Requirements

- [Proxmox Backup Server](https://www.proxmox.com/en/proxmox-backup-server/overview) 3.1 or higher

## Installation

### Edit `/etc/apt/sources.list` and add the following line

```bash
# not for production use
deb http://download.proxmox.com/debian bookworm pve-no-subscription
```

### Edit `/etc/apt/sources.list.d/pbs-enterprise.list` and comment out the following line

```bash
# deb https://enterprise.proxmox.com/debian/pve bookworm pve-enterprise
```

### Update the package lists and upgrade the packages

```bash
apt update && apt upgrade -y
```

### Install nala (apt package manager)

```bash
apt install -y nala sudo
sudo nala fetch
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

### Enable ZFS, done in the Proxmox VE web interface

### Enable smart monitoring

```bash
smartctl -a /dev/sda
```

### Turn [IOMMU (PCI Passthrough)](https://pve.proxmox.com/wiki/PCI_Passthrough)

Edit `/etc/default/grub` and add the following line:

```bash
GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"
```

Edit `/etc/modules` and add the following line:

```bash
vfio
vfio_iommu_type1
vfio_pci
vfio_virqfd
```

### Disable GRUB delay

```bash
sudo nano /etc/default/grub
```

### Update grub

```bash
update-grub
```

### Setup dhcp

```bash
sudo nano /etc/network/interfaces
```

```bash
auto enp1s0
iface enp1s0 inet dhcp

iface wlo1 inet dhcp
```

### Create a new user and add it to the sudo group

```bash
adduser rodneyosodo
usermod -aG sudo rodneyosodo
su - rodneyosodo
```

### Install software

```bash
sudo nala install ssh openssh-server git vim htop bpytop neofetch p7zip tar curl wget make thefuck python3-pip
```

### Setup SSH

Have a look at [debian installation](../debian/README.md)

### Setp tailscale

Have a look at [debian installation](../debian/README.md)

### Setup NTP Update

Install nptdate

```bash
sudo nala install nptdate
```

Test

```bash
sudo ntpdate bohr
```

Copy systemd service

```bash
sudo vim /etc/systemd/system/ntp-update.service
sudo systemctl start ntp-update.service
sudo systemctl enable ntp-update.service
sudo systemctl status ntp-update.service
```

### Setup Network Manager

Install network-manager

```bash
sudo nala install network-manager
```

Configure interfaces to be managed

```bash
sudo vim /etc/NetworkManager/NetworkManager.conf
sudo systemctl restart NetworkManager
```

```bash
mcli d
sudo nmcli r wifi on
sudo nmcli d wifi list
```

```bash
nmcli d wifi connect <SSID> password <PASSWORD>
```

Copy systemd service

```bash
whereis nmcli
sudo vim /etc/systemd/system/wifi-connect.service
sudo systemctl start wifi-connect.service
sudo systemctl enable wifi-connect.service
sudo systemctl status wifi-connect.service
sudo nmcli d wifi list
```
