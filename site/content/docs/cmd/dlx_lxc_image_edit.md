---
title: dlx lxc image edit
description: dlx lxc image edit
lead: dlx lxc image edit
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
## dlx lxc image edit

Edit image properties

### Synopsis

Description:
  Edit image properties



```
dlx lxc image edit [<remote>:]<image> [flags]
```

### Examples

```
  lxc image edit <image>
      Launch a text editor to edit the properties

  lxc image edit <image> < image.yaml
      Load the image properties from a YAML file
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

