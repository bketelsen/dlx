---
title: dlx lxc config trust add
description: dlx lxc config trust add
lead: dlx lxc config trust add
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
## dlx lxc config trust add

Add new trusted clients

### Synopsis

Description:
  Add new trusted clients

  The following certificate types are supported:
  - client (default)
  - metrics




```
dlx lxc config trust add [<remote>:] <cert> [flags]
```

### Options

```
      --name         Alternative certificate name
      --projects     List of projects to restrict the certificate to
      --restricted   Restrict the certificate to one or more projects
      --type         Type of certificate (default "client")
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

* [dlx lxc config trust](/docs/cmd/dlx_lxc_config_trust)	 - Manage trusted clients

