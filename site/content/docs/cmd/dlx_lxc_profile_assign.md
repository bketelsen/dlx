---
title: dlx lxc profile assign
description: dlx lxc profile assign
lead: dlx lxc profile assign
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
## dlx lxc profile assign

Assign sets of profiles to instances

### Synopsis

Description:
  Assign sets of profiles to instances



```
dlx lxc profile assign [<remote>:]<instance> <profiles> [flags]
```

### Examples

```
  lxc profile assign foo default,bar
      Set the profiles for "foo" to "default" and "bar".

  lxc profile assign foo default
      Reset "foo" to only using the "default" profile.

  lxc profile assign foo ''
      Remove all profile from "foo"
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

* [dlx lxc profile](/docs/cmd/dlx_lxc_profile)	 - Manage profiles

