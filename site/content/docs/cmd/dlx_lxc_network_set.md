---
title: dlx lxc network set
description: dlx lxc network set
lead: dlx lxc network set
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
## dlx lxc network set

Set network configuration keys

### Synopsis

Description:
  Set network configuration keys

  For backward compatibility, a single configuration key may still be set with:
      lxc network set [<remote>:]<network> <key> <value>



```
dlx lxc network set [<remote>:]<network> <key>=<value>... [flags]
```

### Options

```
      --target   Cluster member name
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

* [dlx lxc network](/docs/cmd/dlx_lxc_network)	 - Manage and attach instances to networks

