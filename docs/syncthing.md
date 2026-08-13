# Syncthing Setup

This document describes the Syncthing setup used to sync files across the PC, laptop, phone and the `tana` backup server.

## Overview

- **Local instance:** `syncthing@rodneyosodo.service` on Arch Linux
- **Web UI / API:** `http://127.0.0.1:8384`
- **Sync ports:** TCP/UDP `22000` (QUIC `22000`, Relay fallback via `relays.syncthing.net`)
- **Discovery:** global discovery + IPv4/IPv6 local discovery

## Devices

| Name  | Device ID (first 7 chars) | Role                                     |
| ----- | ------------------------- | ---------------------------------------- |
| elgon | `RB4UPYS`                 | Laptop, Arch Linux                       |
| thor  | `RZYEDOC`                 | PC                                       |
| tana  | `MYWK6S5`                 | Proxmox VM, dedicated backup + sync node |
| nzoia | `FTRETWW`                 | Mobile phone                             |

## Folders

| Folder ID     | Label    | Path                   | Type        | Devices           | Versioning           |
| ------------- | -------- | ---------------------- | ----------- | ----------------- | -------------------- |
| `qggat-6vvnd` | Obsidian | `~/Documents/obsidian` | sendreceive | all 4             | simple (keep 5, 30d) |
| `ytecf-pkben` | code     | `~/code`               | sendreceive | elgon, thor, tana | none                 |
| `yythq-ewhw4` | ssh      | `~/.ssh`               | sendreceive | elgon, thor       | simple (keep 5, 30d) |

All folders are `sendreceive` with filesystem watchers enabled (`rescanIntervalS: 3600`, `fsWatcherEnabled: true`).

## The `.stignore` file

The `~/code` folder ignores patterns via `~/code/.stignore`, which is a symlink:

```
~/code/.stignore -> ~/Downloads/dotfiles/.stignore
```

So the ignore rules live in the [dotfiles repo](https://github.com/rodneyosodo/dotfiles) and are shared/versioned.

### Why a code-specific ignore file

`~/code` is a large directory (~160 GB on disk across several repositories). Without ignores, Syncthing would sync build artifacts, dependency trees, toolchain downloads and caches to every peer. The ignore file keeps the synced set to source code only.
