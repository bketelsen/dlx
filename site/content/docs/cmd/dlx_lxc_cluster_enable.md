---
title: dlx lxc cluster enable
description: dlx lxc cluster enable
lead: dlx lxc cluster enable
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
## dlx lxc cluster enable

Enable clustering on a single non-clustered LXD server

### Synopsis

Description:
  Enable clustering on a single non-clustered LXD server

    This command turns a non-clustered LXD server into the first member of a new
    LXD cluster, which will have the given name.

    It's required that the LXD is already available on the network. You can check
    that by running 'lxc config get core.https_address', and possibly set a value
    for the address if not yet set.



```
dlx lxc cluster enable [<remote>:] <name> [flags]
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

* [dlx lxc cluster](/docs/cmd/dlx_lxc_cluster)	 - Manage cluster members

