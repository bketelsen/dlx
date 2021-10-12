---
title: dlx lxc file push
description: dlx lxc file push
lead: dlx lxc file push
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
## dlx lxc file push

Push files into instances

### Synopsis

Description:
  Push files into instances



```
dlx lxc file push <source path> [<remote>:]<instance>/<path> [[<remote>:]<instance>/<path>...] [flags]
```

### Examples

```
  lxc file push /etc/hosts foo/etc/hosts
     To push /etc/hosts into the instance "foo".
```

### Options

```
  -p, --create-dirs   Create any directories necessary
      --gid           Set the file's gid on push (default -1)
      --mode          Set the file's perms on push
  -r, --recursive     Recursively transfer files
      --uid           Set the file's uid on push (default -1)
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

* [dlx lxc file](/docs/cmd/dlx_lxc_file)	 - Manage files in instances

