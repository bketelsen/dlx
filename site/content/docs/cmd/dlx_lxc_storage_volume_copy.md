---
title: dlx lxc storage volume copy
description: dlx lxc storage volume copy
lead: dlx lxc storage volume copy
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
## dlx lxc storage volume copy

Copy storage volumes

### Synopsis

Description:
  Copy storage volumes



```
dlx lxc storage volume copy <pool>/<volume>[/<snapshot>] <pool>/<volume> [flags]
```

### Options

```
      --mode             Transfer mode. One of pull (default), push or relay. (default "pull")
      --refresh          Refresh and update the existing storage volume copies
      --target           Cluster member name
      --target-project   Copy to a project different from the source
      --volume-only      Copy the volume without its snapshots
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

