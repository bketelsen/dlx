---
title: dlx lxc storage set
description: dlx lxc storage set
lead: dlx lxc storage set
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
## dlx lxc storage set

Set storage pool configuration keys

### Synopsis

Description:
  Set storage pool configuration keys

  For backward compatibility, a single configuration key may still be set with:
      lxc storage set [<remote>:]<pool> <key> <value>



```
dlx lxc storage set [<remote>:]<pool> <key> <value> [flags]
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

* [dlx lxc storage](/docs/cmd/dlx_lxc_storage)	 - Manage storage pools and volumes

