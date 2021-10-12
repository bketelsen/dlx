---
title: dlx lxc storage volume edit
description: dlx lxc storage volume edit
lead: dlx lxc storage volume edit
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
## dlx lxc storage volume edit

Edit storage volume configurations as YAML

### Synopsis

Description:
  Edit storage volume configurations as YAML



```
dlx lxc storage volume edit [<remote>:]<pool> <volume>[/<snapshot>] [flags]
```

### Examples

```
  lxc storage volume edit [<remote>:]<pool> <volume> < volume.yaml
      Update a storage volume using the content of pool.yaml.
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

