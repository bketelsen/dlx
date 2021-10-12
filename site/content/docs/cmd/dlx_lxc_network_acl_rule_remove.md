---
title: dlx lxc network acl rule remove
description: dlx lxc network acl rule remove
lead: dlx lxc network acl rule remove
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
## dlx lxc network acl rule remove

Remove rules from an ACL

### Synopsis

Description:
  Remove rules from an ACL



```
dlx lxc network acl rule remove [<remote>:]<ACL> <direction> <key>=<value>... [flags]
```

### Options

```
      --force   Remove all rules that match
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

* [dlx lxc network acl rule](/docs/cmd/dlx_lxc_network_acl_rule)	 - Manage network ACL rules

