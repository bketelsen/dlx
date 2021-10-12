---
title: dlx lxc config set
description: dlx lxc config set
lead: dlx lxc config set
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
## dlx lxc config set

Set instance or server configuration keys

### Synopsis

Description:
  Set instance or server configuration keys

  For backward compatibility, a single configuration key may still be set with:
      lxc config set [<remote>:][<instance>] <key> <value>



```
dlx lxc config set [<remote>:][<instance>] <key>=<value>... [flags]
```

### Examples

```
  lxc config set [<remote>:]<instance> limits.cpu=2
      Will set a CPU limit of "2" for the instance.

  lxc config set core.https_address=[::]:8443
      Will have LXD listen on IPv4 and IPv6 port 8443.

  lxc config set core.trust_password=blah
      Will set the server's trust password to blah.
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

