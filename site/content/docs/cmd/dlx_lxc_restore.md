---
title: dlx lxc restore
description: dlx lxc restore
lead: dlx lxc restore
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
## dlx lxc restore

Restore instances from snapshots

### Synopsis

Description:
  Restore instances from snapshots

  If --stateful is passed, then the running state will be restored too.



```
dlx lxc restore [<remote>:]<instance> <snapshot> [flags]
```

### Examples

```
  lxc snapshot u1 snap0
      Create the snapshot.

  lxc restore u1 snap0
      Restore the snapshot.
```

### Options

```
      --stateful   Whether or not to restore the instance's running state from snapshot (if available)
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

