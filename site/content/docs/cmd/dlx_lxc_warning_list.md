---
title: dlx lxc warning list
description: dlx lxc warning list
lead: dlx lxc warning list
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
## dlx lxc warning list

List warnings

### Synopsis

Description:
  List warnings

  The -c option takes a (optionally comma-separated) list of arguments
  that control which warning attributes to output when displaying in table
  or csv format.

  Default column layout is: utSscpLl

  Column shorthand chars:

      c - Count
      l - Last seen
      L - Location
      f - First seen
      p - Project
      s - Severity
      S - Status
      u - UUID
      t - Type



```
dlx lxc warning list [<remote>:] [flags]
```

### Options

```
  -a, --all       List all warnings
  -c, --columns   Columns (default "utSscpLl")
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

* [dlx lxc warning](/docs/cmd/dlx_lxc_warning)	 - Manage warnings

