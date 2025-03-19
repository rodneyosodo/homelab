# Kubernetes

## Requirements

- [Galana](terraform/galana)
- [Turkwel](terraform/turkwel)
- [Yala](terraform/yala)

## Setup

### Galana

```bash
ssh rodneyosodo@galana
curl -sfL https://get.k3s.io | sh -
sudo scp /etc/rancher/k3s/k3s.yaml rodneyosodo@thor:/home/rodneyosodo/Downloads/k3s-config
sudo cat /var/lib/rancher/k3s/server/node-token
```

Change server address to `https://galana:6443`

### Turkwel & Yala

```bash
ssh rodneyosodo@turkwel
ssh rodneyosodo@yala
curl -sfL https://get.k3s.io | K3S_URL=https://galana:6443 K3S_TOKEN=mynodetoken sh -
```
