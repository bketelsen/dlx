---
title: dlx lxc snapshot
description: dlx lxc snapshot
lead: dlx lxc snapshot
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
## dlx lxc snapshot

Create instance snapshots

### Synopsis

Description:
  Create instance snapshots

  When --stateful is used, LXD attempts to checkpoint the instance's
  running state, including process memory state, TCP connections, ...



```
dlx lxc snapshot [<remote>:]<instance> [<snapshot name>] [flags]
```

### Examples

```
  lxc snapshot u1 snap0
      Create a snapshot of "u1" called "snap0".
```

### Options

```
      --no-expiry   Ignore any configured auto-expiry for the instance
      --reuse       If the snapshot name already exists, delete and create a new one
      --stateful    Whether or not to snapshot the instance's running state
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

