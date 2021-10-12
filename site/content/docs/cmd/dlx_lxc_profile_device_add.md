---
title: dlx lxc profile device add
description: dlx lxc profile device add
lead: dlx lxc profile device add
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
## dlx lxc profile device add

Add instance devices

### Synopsis

Description:
  Add instance devices



```
dlx lxc profile device add [<remote>:]<profile> <device> <type> [key=value...] [flags]
```

### Examples

```
  lxc profile device add [<remote>:]profile1 <device-name> disk source=/share/c1 path=opt
      Will mount the host's /share/c1 onto /opt in the instance.
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

* [dlx lxc profile device](/docs/cmd/dlx_lxc_profile_device)	 - Manage devices

