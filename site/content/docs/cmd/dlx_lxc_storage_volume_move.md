---
title: dlx lxc storage volume move
description: dlx lxc storage volume move
lead: dlx lxc storage volume move
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
## dlx lxc storage volume move

Move storage volumes between pools

### Synopsis

Description:
  Move storage volumes between pools



```
dlx lxc storage volume move <pool>/<volume> <pool>/<volume> [flags]
```

### Options

```
      --mode             Transfer mode, one of pull (default), push or relay (default "pull")
      --target           Cluster member name
      --target-project   Move to a project different from the source
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

