---
title: dlx lxc image list
description: dlx lxc image list
lead: dlx lxc image list
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
## dlx lxc image list

List images

### Synopsis

Description:
  List images

  Filters may be of the <key>=<value> form for property based filtering,
  or part of the image hash or part of the image alias name.

  The -c option takes a (optionally comma-separated) list of arguments
  that control which image attributes to output when displaying in table
  or csv format.

  Default column layout is: lfpdasu

  Column shorthand chars:

      l - Shortest image alias (and optionally number of other aliases)
      L - Newline-separated list of all image aliases
      f - Fingerprint (short)
      F - Fingerprint (long)
      p - Whether image is public
      d - Description
      a - Architecture
      s - Size
      u - Upload date
      t - Type



```
dlx lxc image list [<remote>:] [<filter>...] [flags]
```

### Options

```
  -c, --columns   Columns (default "lfpdatsu")
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

* [dlx lxc image](/docs/cmd/dlx_lxc_image)	 - Manage images

