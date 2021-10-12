---
title: dlx lxc launch
description: dlx lxc launch
lead: dlx lxc launch
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
## dlx lxc launch

Create and start instances from images

### Synopsis

Description:
  Create and start instances from images



```
dlx lxc launch [<remote>:]<image> [<remote>:][<name>] [flags]
```

### Examples

```
  lxc launch ubuntu:18.04 u1

  lxc launch ubuntu:18.04 u1 < config.yaml
      Create and start the instance with configuration from config.yaml
```

### Options

```
  -c, --config                Config key/value to apply to the new instance
      --console[="console"]   Immediately attach to the console
      --empty                 Create an empty instance
  -e, --ephemeral             Ephemeral instance
  -n, --network               Network name
      --no-profiles           Create the instance with no profiles applied
  -p, --profile               Profile to apply to the new instance
  -s, --storage               Storage pool name
      --target                Cluster member name
  -t, --type                  Instance type
      --vm                    Create a virtual machine
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

