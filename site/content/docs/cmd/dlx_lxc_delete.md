---
title: dlx lxc delete
description: dlx lxc delete
lead: dlx lxc delete
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
## dlx lxc delete

Delete instances and snapshots

### Synopsis

Description:
  Delete instances and snapshots



```
dlx lxc delete [<remote>:]<instance>[/<snapshot>] [[<remote>:]<instance>[/<snapshot>]...] [flags]
```

### Options

```
  -f, --force         Force the removal of running instances
  -i, --interactive   Require user confirmation
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

