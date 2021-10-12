---
title: dlx lxc storage volume export
description: dlx lxc storage volume export
lead: dlx lxc storage volume export
date: 2021-10-12T10:25:15Z
lastmod: 2021-10-12T10:25:15Z
draft: false
images: []
menu:
  docs:
    parent: "cli"
weight: 100
toc: true
---
## dlx lxc storage volume export

Export custom storage volume

### Synopsis

Description:
  Export custom storage volume



```
dlx lxc storage volume export [<remote>:]<pool> <volume> [<path>] [flags]
```

### Options

```
      --compression         Define a compression algorithm: for backup or none
      --optimized-storage   Use storage driver optimized format (can only be restored on a similar pool)
      --target              Cluster member name
      --volume-only         Export the volume without its snapshots
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

