---
title: dlx lxc network forward set
description: dlx lxc network forward set
lead: dlx lxc network forward set
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
## dlx lxc network forward set

Set network forward keys

### Synopsis

Description:
  Set network forward keys

  For backward compatibility, a single configuration key may still be set with:
      lxc network set [<remote>:]<network> <listen_address> <key> <value>



```
dlx lxc network forward set [<remote>:]<network> <listen_address> <key>=<value>... [flags]
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

* [dlx lxc network forward](/docs/cmd/dlx_lxc_network_forward)	 - Manage network forwards

