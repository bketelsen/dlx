---
title: dlx lxc copy
description: dlx lxc copy
lead: dlx lxc copy
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
## dlx lxc copy

Copy instances within or in between LXD servers

### Synopsis

Description:
  Copy instances within or in between LXD servers



```
dlx lxc copy [<remote>:]<source>[/<snapshot>] [[<remote>:]<destination>] [flags]
```

### Options

```
  -c, --config           Config key/value to apply to the new instance
  -d, --device           New key/value to apply to a specific device
  -e, --ephemeral        Ephemeral instance
      --instance-only    Copy the instance without its snapshots
      --mode             Transfer mode. One of pull (default), push or relay (default "pull")
      --no-profiles      Create the instance with no profiles applied
  -p, --profile          Profile to apply to the new instance
      --refresh          Perform an incremental copy
      --stateless        Copy a stateful instance stateless
  -s, --storage          Storage pool name
      --target           Cluster member name
      --target-project   Copy to a project different from the source
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

