---
title: dlx lxc storage volume rename
description: dlx lxc storage volume rename
lead: dlx lxc storage volume rename
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
## dlx lxc storage volume rename

Rename storage volumes and storage volume snapshots

### Synopsis

Description:
  Rename storage volumes



```
dlx lxc storage volume rename [<remote>:]<pool> <old name>[/<old snapshot name>] <new name>[/<new snapshot name>] [flags]
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

