# Odin Proxmox Setup

## Requirements

- [Proxmox VE](https://www.proxmox.com/en/downloads) 8.2 or higher

## Installation

1. Flash the Odin Proxmox image to a USB drive.

```bash
# writing an ISO image to a USB stick using the dd command
sudo dd if=proxmox-ve_8.2-2.iso of=/dev/sdc bs=1M conv=sync status=progress
```

2. Boot the Odin Proxmox server.
3. Configure the server with.
   - DISK: btrfs raid0 (raid0)
4. Reboot the server.

## Post Installation

1. Login to the server UI.
2. Enable ZFS, done in the Proxmox VE web interface.
3. Make proxmox VLAN aware on UI.
4. Change DNS to `1.1.1.1` and `8.8.8.8`
5. Run the following commands:

   ```bash
   curl -fsSL https://tailscale.com/install.sh | sh
   tailscale up --ssh --accept-routes --advertise-exit-node
   zpool import -f yatta
   ```

6. Login to the tailscale UI.
7. Run ansible playbooks.

   ```bash
   cd ~/homelab/ansible
   make setup-odin
   ```

8. Enable smart monitoring:

   ```bash
   smartctl -a /dev/sda
   ```

9. Add cloud images to proxmox UI.
   - [Ubuntu 24.04 LTS (noble-server-amd64)](https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img)
