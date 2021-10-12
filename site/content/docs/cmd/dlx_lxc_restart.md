---
title: dlx lxc restart
description: dlx lxc restart
lead: dlx lxc restart
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
## dlx lxc restart

Restart instances

### Synopsis

Description:
  Restart instances

  The opposite of "lxc pause" is "lxc start".



```
dlx lxc restart [<remote>:]<instance> [[<remote>:]<instance>...] [flags]
```

### Options

```
      --all                   Run against all instances
      --console[="console"]   Immediately attach to the console
  -f, --force                 Force the instance to shutdown
      --timeout               Time to wait for the instance before killing it (default -1)
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

