terraform {
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.66.3"
    }
  }
}

variable "proxmox_url" {
  type = string
  description = "Proxmox URL"
}

variable "proxmox_username" {
  type = string
  description = "Proxmox username"
}

variable "proxmox_password" {
  type = string
  description = "Proxmox password for the user"
}

provider "proxmox" {
  endpoint = var.proxmox_url
  insecure = true
  username = var.proxmox_username
  password = var.proxmox_password
}
