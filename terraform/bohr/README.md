# Bohr

This directory contains the Terraform configuration files for my homelab server.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/install)

## Environment Variables

Create a copy of the [`terraform.tfvars.example`](./terraform.tfvars.example) file and rename it to `terraform.tfvars`.

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit the `terraform.tfvars` file and update the values.

## Usage

### Plan

```bash
terraform plan
```

### Apply

```bash
terraform apply
```

### Destroy

```bash
terraform destroy
```
