---
title: dlx lxc storage edit
description: dlx lxc storage edit
lead: dlx lxc storage edit
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
## dlx lxc storage edit

Edit storage pool configurations as YAML

### Synopsis

Description:
  Edit storage pool configurations as YAML



```
dlx lxc storage edit [<remote>:]<pool> [flags]
```

### Examples

```
  lxc storage edit [<remote>:]<pool> < pool.yaml
      Update a storage pool using the content of pool.yaml.
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

* [dlx lxc storage](/docs/cmd/dlx_lxc_storage)	 - Manage storage pools and volumes

