---
title: dlx lxc storage volume show
description: dlx lxc storage volume show
lead: dlx lxc storage volume show
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
## dlx lxc storage volume show

Show storage volume configurations

### Synopsis

Description:
  Show storage volume configurations



```
dlx lxc storage volume show [<remote>:]<pool> <volume>[/<snapshot>] [flags]
```

### Examples

```
  lxc storage volume show default data
      Will show the properties of a custom volume called "data" in the "default" pool.

  lxc storage volume show default container/data
      Will show the properties of the filesystem for a container called "data" in the "default" pool.
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

