---
title: dlx lxc storage volume snapshot
description: dlx lxc storage volume snapshot
lead: dlx lxc storage volume snapshot
date: 2021-10-12T10:37:58Z
lastmod: 2021-10-12T10:37:58Z
draft: false
images: []
menu:
  docs:
    parent: "cli"
weight: 100
toc: true
---
## dlx lxc storage volume snapshot

Snapshot storage volumes

### Synopsis

Description:
  Snapshot storage volumes



```
dlx lxc storage volume snapshot [<remote>:]<pool> <volume> [<snapshot>] [flags]
```

### Options

```
      --no-expiry   Ignore any configured auto-expiry for the storage volume
      --reuse       If the snapshot name already exists, delete and create a new one
      --target      Cluster member name
```

### Options inherited from parent commands

```
      --debug         Show all debug messages
      --force-local   Force using the local unix socket
  -h, --help          Print help
      --project       Override the source project
  -q, --quiet         Don't show progress information
  -v, --verbose       Show all information messages
      --version       Print version number
```

### SEE ALSO

* [dlx lxc storage volume](/docs/cmd/dlx_lxc_storage_volume)	 - Manage storage volumes

