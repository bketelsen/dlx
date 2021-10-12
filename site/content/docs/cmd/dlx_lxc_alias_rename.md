---
title: dlx lxc alias rename
description: dlx lxc alias rename
lead: dlx lxc alias rename
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
## dlx lxc alias rename

Rename aliases

### Synopsis

Description:
  Rename aliases



```
dlx lxc alias rename <old alias> <new alias> [flags]
```

### Examples

```
  lxc alias rename list my-list
      Rename existing alias "list" to "my-list".
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

