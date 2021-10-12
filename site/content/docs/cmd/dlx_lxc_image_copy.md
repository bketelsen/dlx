---
title: dlx lxc image copy
description: dlx lxc image copy
lead: dlx lxc image copy
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
## dlx lxc image copy

Copy images between servers

### Synopsis

Description:
  Copy images between servers

  The auto-update flag instructs the server to keep this image up to date.
  It requires the source to be an alias and for it to be public.



```
dlx lxc image copy [<remote>:]<image> <remote>: [flags]
```

### Options

```
      --alias          New aliases to add to the image
      --auto-update    Keep the image up to date after initial copy
      --copy-aliases   Copy aliases from source
      --mode           Transfer mode. One of pull (default), push or relay (default "pull")
      --public         Make image public
      --vm             Copy virtual machine images
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

