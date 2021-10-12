---
title: dlx lxc cluster remove
description: dlx lxc cluster remove
lead: dlx lxc cluster remove
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
## dlx lxc cluster remove

Remove a member from the cluster

### Synopsis

Description:
  Remove a member from the cluster



```
dlx lxc cluster remove [<remote>:]<member> [flags]
```

### Options

```
  -f, --force   Force removing a member, even if degraded
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

