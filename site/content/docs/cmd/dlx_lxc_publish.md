---
title: dlx lxc publish
description: dlx lxc publish
lead: dlx lxc publish
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
## dlx lxc publish

Publish instances as images

### Synopsis

Description:
  Publish instances as images



```
dlx lxc publish [<remote>:]<instance>[/<snapshot>] [<remote>:] [flags] [key=value...]
```

### Options

```
      --alias              New alias to define at target
      --compression none   Compression algorithm to use (none for uncompressed)
      --expire             Image expiration date (format: rfc3339)
  -f, --force              Stop the instance if currently running
      --public             Make the image public
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

