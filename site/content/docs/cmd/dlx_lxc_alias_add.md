---
title: dlx lxc alias add
description: dlx lxc alias add
lead: dlx lxc alias add
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
## dlx lxc alias add

Add new aliases

### Synopsis

Description:
  Add new aliases



```
dlx lxc alias add <alias> <target> [flags]
```

### Examples

```
  lxc alias add list "list -c ns46S"
      Overwrite the "list" command to pass -c ns46S.
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

* [dlx lxc alias](/docs/cmd/dlx_lxc_alias)	 - Manage command aliases

