# Heimdall Proxmox Backup Server Setup

## Requirements

- [Proxmox Backup Server](https://proxmox.com/en/downloads/proxmox-backup-server) 3.3 or higher

## Installation

1. Flash the Heimdall Proxmox Backup Server image to a USB drive.

```bash
# writing an ISO image to a USB stick using the dd command
sudo dd if=proxmox-backup-server_3.3-1.iso of=/dev/sda1 bs=1M conv=sync status=progress
```

2. Boot the Heimdall Proxmox Backup Server.
3. Configure the server with.
   - DISK: zfs raid0 (raid0)
4. Reboot the server.

## Post Installation

1. Connect to ethernet port.
2. Login to the server UI.
3. Run ansible playbooks.

```bash
cd ~/homelab/ansible
make setup-heimdall
```

6. Run smart monitoring:

   ```bash
   smartctl -a /dev/sda
   ```

7. Run the following commands:

   ```bash
   curl -fsSL https://tailscale.com/install.sh | sh
   tailscale up --ssh --accept-routes --advertise-exit-node
   zpool import -f nyika
   ```

8. Login to the tailscale UI.
