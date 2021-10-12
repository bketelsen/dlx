---
title: dlx lxc import
description: dlx lxc import
lead: dlx lxc import
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
## dlx lxc import

Import instance backups

### Synopsis

Description:
  Import backups of instances including their snapshots.



```
dlx lxc import [<remote>:] <backup file> [<instance name>] [flags]
```

### Examples

```
  lxc import backup0.tar.gz
      Create a new instance using backup0.tar.gz as the source.
```

### Options

```
  -s, --storage   Storage pool name
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

* [dlx lxc](/docs/cmd/dlx_lxc)	 - Command line client for LXD

