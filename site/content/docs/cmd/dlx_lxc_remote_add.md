---
title: dlx lxc remote add
description: dlx lxc remote add
lead: dlx lxc remote add
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
## dlx lxc remote add

Add new remote servers

### Synopsis

Description:
  Add new remote servers

  URL for remote resources must be HTTPS (https://).

  Basic authentication can be used when combined with the "simplestreams" protocol:
    lxc remote add some-name https://LOGIN:PASSWORD@example.com/some/path --protocol=simplestreams




```
dlx lxc remote add [<remote>] <IP|FQDN|URL> [flags]
```

### Options

```
      --accept-certificate   Accept certificate
      --auth-type            Server authentication type (tls or candid)
      --domain               Candid domain to use
      --password             Remote admin password
      --protocol             Server protocol (lxd or simplestreams)
      --public               Public image server
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

* [dlx lxc remote](/docs/cmd/dlx_lxc_remote)	 - Manage the list of remote servers

