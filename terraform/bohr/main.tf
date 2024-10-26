resource "proxmox_virtual_environment_vm" "debian_vm" {
  name      = "bohr"
  node_name = "odin"
  vm_id     = 300
  on_boot   = true
  tags      = ["docker", "ubuntu", "production"]
  bios      = "ovmf"

  operating_system {
    type = "l26"
  }

  efi_disk {
    datastore_id = "yatta"
    file_format  = "raw"
    type         = "4m"
  }

  tpm_state {
    datastore_id = "yatta"
    version      = "v2.0"
  }

  agent {
    enabled = true
  }

  disk {
    datastore_id = "yatta"
    file_id      = "local-btrfs:iso/noble-server-cloudimg-amd64.img"
    size         = 500
    interface    = "scsi0"
  }

  cpu {
    architecture = "x86_64"
    cores        = 8
    flags        = ["+aes"]
    sockets      = 1
    type         = "x86-64-v2-AES"
  }

  memory {
    dedicated = 20480
    floating  = 20480
  }

  network_device {
    bridge = "vmbr0"
  }

  initialization {
    datastore_id      = "yatta"
    user_data_file_id = proxmox_virtual_environment_file.cloud_init_config.id

    ip_config {
      ipv4 {
        address = "192.168.100.32/24"
        gateway = "192.168.100.1"
      }
    }
  }

  keyboard_layout = "en-us"
}

variable "vm_username" {
  type        = string
  description = "VM username"
}

variable "vm_password" {
  type        = string
  description = "VM password for the user"
}

variable "vm_username_gecos" {
  type        = string
  description = "VM username gecos"
}

variable "vm_hostname" {
  type        = string
  description = "VM hostname"
}

variable "vm_fqdn" {
  type        = string
  description = "VM fqdn"
}

variable "tailscale_auth_key" {
  type        = string
  description = "Tailscale auth key"
}

resource "proxmox_virtual_environment_file" "cloud_init_config" {
  content_type = "snippets"
  datastore_id = "local-btrfs"
  node_name    = "odin"

  source_raw {
    data      = templatefile("cloudinit.tfpl", { username = var.vm_username, vm_username_gecos = var.vm_username_gecos, password = var.vm_password, hostname = var.vm_hostname, fqdn = var.vm_fqdn, tailscale_auth_key = var.tailscale_auth_key })
    file_name = "cloud-init-config.yaml"
  }
}
