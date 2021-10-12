---
title: dlx lxc storage volume list
description: dlx lxc storage volume list
lead: dlx lxc storage volume list
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
## dlx lxc storage volume list

List storage volumes

### Synopsis

Description:
  List storage volumes

  The -c option takes a (optionally comma-separated) list of arguments
  that control which image attributes to output when displaying in table
  or csv format.

  Default column layout is: lfpdasu

  Column shorthand chars:

      t - Type of volume (custom, image, container or virtual-machine)
      n - Name
      d - Description
      c - Content type (filesystem or block)
      u - Number of references (used by)
      L - Location of the instance (e.g. its cluster member)
      U - Current disk usage



```
dlx lxc storage volume list [<remote>:]<pool> [flags]
```

### Options

```
  -c, --columns   Columns (default "tndcuL")
  -f, --format    Format (csv|json|table|yaml) (default "table")
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

* [dlx lxc storage volume](/docs/cmd/dlx_lxc_storage_volume)	 - Manage storage volumes

