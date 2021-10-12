---
title: dlx lxc info
description: dlx lxc info
lead: dlx lxc info
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
## dlx lxc info

Show instance or server information

### Synopsis

Description:
  Show instance or server information



```
dlx lxc info [<remote>:][<instance>] [flags]
```

### Examples

```
  lxc info [<remote>:]<instance> [--show-log]
      For instance information.

  lxc info [<remote>:] [--resources]
      For LXD server information.
```

### Options

```
      --resources   Show the resources available to the server
      --show-log    Show the instance's last 100 log lines?
      --target      Cluster member name
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

* [dlx lxc](/docs/cmd/dlx_lxc)	 - Command line client for LXD

