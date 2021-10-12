---
title: dlx lxc file pull
description: dlx lxc file pull
lead: dlx lxc file pull
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
## dlx lxc file pull

Pull files from instances

### Synopsis

Description:
  Pull files from instances



```
dlx lxc file pull [<remote>:]<instance>/<path> [[<remote>:]<instance>/<path>...] <target path> [flags]
```

### Examples

```
  lxc file pull foo/etc/hosts .
     To pull /etc/hosts from the instance and write it to the current directory.
```

### Options

```
  -p, --create-dirs   Create any directories necessary
  -r, --recursive     Recursively transfer files
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

