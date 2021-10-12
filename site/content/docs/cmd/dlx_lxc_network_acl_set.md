---
title: dlx lxc network acl set
description: dlx lxc network acl set
lead: dlx lxc network acl set
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
## dlx lxc network acl set

Set network ACL configuration keys

### Synopsis

Description:
  Set network ACL configuration keys

  For backward compatibility, a single configuration key may still be set with:
      lxc network set [<remote>:]<ACL> <key> <value>



```
dlx lxc network acl set [<remote>:]<ACL> <key>=<value>... [flags]
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

* [dlx lxc network acl](/docs/cmd/dlx_lxc_network_acl)	 - Manage network ACLs

