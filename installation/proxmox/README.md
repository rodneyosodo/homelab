# Proxmox Installation

## Requirements

- [Proxmox VE](https://www.proxmox.com/en/downloads) 8.1 or higher

## Installation

1. Edit `/etc/apt/sources.list` and add the following line:

   ```bash
   # not for production use
   deb http://download.proxmox.com/debian bookworm pve-no-subscription
   ```

2. Edit `/etc/apt/sources.list.d/pve-enterprise.list` and comment out the following line:

   ```bash
   # deb https://enterprise.proxmox.com/debian/pve bookworm pve-enterprise
   ```

3. Edit `/etc/apt/sources.list.d/ceph.list` and comment out the following line:

   ```bash
   # deb https://enterprise.proxmox.com/debian/ceph-quincy bookworm enterprise
   ```

4. Update the package lists and upgrade the packages:

   ```bash
   apt update && apt upgrade -y
   ```

5. Clear disk storage:

   ```bash
   fdisk /dev/sda
   ```

   ```bash
    Command (m for help): p
    Command (m for help): d
    Command (m for help): w
   ```

6. Enable ZFS, done in the Proxmox VE web interface.
7. Enable smart monitoring:

   ```bash
   smartctl -a /dev/sda
   ```

8. Turn [IOMMU (PCI Passthrough)](https://pve.proxmox.com/wiki/PCI_Passthrough)

   Edit `/etc/default/grub` and add the following line:

   ```bash
   GRUB_CMDLINE_LINUX_DEFAULT="quiet amd_iommu=on"
   ```

   Update grub:

   ```bash
   update-grub
   ```

   Edit `/etc/modules` and add the following line:

   ```bash
   vfio
   vfio_iommu_type1
   vfio_pci
   vfio_virqfd
   ```

9. Make proxmox VLAN aware on UI:
10. Adding microcode updates:

    ```bash
    bash -c "$(wget -qLO - https://github.com/tteck/Proxmox/raw/main/misc/microcode.sh)"
    ```

11. Passthrough USB nic to VM:

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
