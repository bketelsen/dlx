---
title: dlx lxc move
description: dlx lxc move
lead: dlx lxc move
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
## dlx lxc move

Move instances within or in between LXD servers

### Synopsis

Description:
  Move instances within or in between LXD servers



```
dlx lxc move [<remote>:]<instance>[/<snapshot>] [<remote>:][<instance>[/<snapshot>]] [flags]
```

### Examples

```
  lxc move [<remote>:]<source instance> [<remote>:][<destination instance>] [--instance-only]
      Move an instance between two hosts, renaming it if destination name differs.

  lxc move <old name> <new name> [--instance-only]
      Rename a local instance.

  lxc move <instance>/<old snapshot name> <instance>/<new snapshot name>
      Rename a snapshot.
```

### Options

```
  -c, --config           Config key/value to apply to the target instance
  -d, --device           New key/value to apply to a specific device
      --instance-only    Move the instance without its snapshots
      --mode             Transfer mode. One of pull (default), push or relay. (default "pull")
      --no-profiles      Unset all profiles on the target instance
  -p, --profile          Profile to apply to the target instance
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

