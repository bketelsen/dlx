---
title: dlx lxc operation show
description: dlx lxc operation show
lead: dlx lxc operation show
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
## dlx lxc operation show

Show details on a background operation

### Synopsis

Description:
  Show details on a background operation



```
dlx lxc operation show [<remote>:]<operation> [flags]
```

### Examples

```
  lxc operation show 344a79e4-d88a-45bf-9c39-c72c26f6ab8a
      Show details on that operation UUID
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

* [dlx lxc operation](/docs/cmd/dlx_lxc_operation)	 - List, show and delete background operations

