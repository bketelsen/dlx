---
title: dlx lxc storage volume import
description: dlx lxc storage volume import
lead: dlx lxc storage volume import
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
## dlx lxc storage volume import

Import custom storage volumes

### Synopsis

Description:
  Import backups of custom volumes including their snapshots.



```
dlx lxc storage volume import [<remote>:]<pool> <backup file> [<volume name>] [flags]
```

### Examples

```
  lxc storage volume import default backup0.tar.gz
  		Create a new custom volume using backup0.tar.gz as the source.
```

### Options

```
      --target   Cluster member name
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

