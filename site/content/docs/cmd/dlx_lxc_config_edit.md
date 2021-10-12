---
title: dlx lxc config edit
description: dlx lxc config edit
lead: dlx lxc config edit
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
## dlx lxc config edit

Edit instance or server configurations as YAML

### Synopsis

Description:
  Edit instance or server configurations as YAML



```
dlx lxc config edit [<remote>:][<instance>[/<snapshot>]] [flags]
```

### Examples

```
  lxc config edit <instance> < instance.yaml
      Update the instance configuration from config.yaml.
```

### Options

```
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

* [dlx lxc config](/docs/cmd/dlx_lxc_config)	 - Manage instance and server configuration options

