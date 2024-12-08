# Odin Proxmox Setup

## Requirements

- [Proxmox VE](https://www.proxmox.com/en/downloads) 8.2 or higher

## Installation

1. Flash the Odin Proxmox image to a USB drive.
2. Boot the Odin Proxmox server.
3. Configure the server with.
   - DISK: btrfs raid0 (raid0)
4. Reboot the server.

## Post Installation

1. Login to the server UI.
2. Diable `ceph` and `pve-enterprise` repositories.
3. Change DNS to ```1.1.1.1``` and ```8.8.8.8```
4. Run the following commands:

   ```bash
   apt update
   curl -fsSL https://tailscale.com/install.sh | sh
   tailscale up --ssh --accept-routes --advertise-exit-node
   zpool import -f yatta
   ```

5. Login to the tailscale UI.
6. Run ansible playbooks.

   ```bash
   cd ~/homelab/ansible
   make setup-odin
   ```

7. Add cloud images to proxmox UI.
   - [Ubuntu 24.04 LTS (noble-server-amd64)](https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img)
