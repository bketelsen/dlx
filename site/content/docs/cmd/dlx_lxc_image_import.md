---
title: dlx lxc image import
description: dlx lxc image import
lead: dlx lxc image import
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
## dlx lxc image import

Import images into the image store

### Synopsis

Description:
  Import image into the image store

  Directory import is only available on Linux and must be performed as root.



```
dlx lxc image import <tarball>|<directory>|<URL> [<rootfs tarball>] [<remote>:] [key=value...] [flags]
```

### Options

```
      --alias    New aliases to add to the image
      --public   Make image public
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

* [dlx lxc image](/docs/cmd/dlx_lxc_image)	 - Manage images

