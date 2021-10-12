---
title: dlx lxc network forward port remove
description: dlx lxc network forward port remove
lead: dlx lxc network forward port remove
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
## dlx lxc network forward port remove

Remove ports from a forward

### Synopsis

Description:
  Remove ports from a forward



```
dlx lxc network forward port remove [<remote>:]<network> <listen_address> [<protocol>] [<listen_port(s)>] [flags]
```

### Options

```
      --force    Remove all ports that match
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

* [dlx lxc network forward port](/docs/cmd/dlx_lxc_network_forward_port)	 - Manage network forward ports

