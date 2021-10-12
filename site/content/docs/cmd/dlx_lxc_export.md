---
title: dlx lxc export
description: dlx lxc export
lead: dlx lxc export
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
## dlx lxc export

Export instance backups

### Synopsis

Description:
  Export instances as backup tarballs.



```
dlx lxc export [<remote>:]<instance> [target] [--instance-only] [--optimized-storage] [flags]
```

### Examples

```
  lxc export u1 backup0.tar.gz
      Download a backup tarball of the u1 instance.
```

### Options

```
      --compression none    Compression algorithm to use (none for uncompressed)
      --instance-only       Whether or not to only backup the instance (without snapshots)
      --optimized-storage   Use storage driver optimized format (can only be restored on a similar pool)
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

