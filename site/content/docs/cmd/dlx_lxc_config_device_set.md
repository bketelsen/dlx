---
title: dlx lxc config device set
description: dlx lxc config device set
lead: dlx lxc config device set
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
## dlx lxc config device set

Set device configuration keys

### Synopsis

Description:
  Set device configuration keys

  For backward compatibility, a single configuration key may still be set with:
      lxc config device set [<remote>:]<instance> <device> <key> <value>



```
dlx lxc config device set [<remote>:]<instance> <device> <key>=<value>... [flags]
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

* [dlx lxc config device](/docs/cmd/dlx_lxc_config_device)	 - Manage devices

