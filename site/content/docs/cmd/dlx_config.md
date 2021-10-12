---
title: dlx config
description: dlx config
lead: dlx config
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
## dlx config

manage global configurations

```
dlx config [flags]
```

### Options

```
  -b, --baseimage string       Default base image for new containers (default "dlxbase")
  -c, --create                 Create global config file
  -p, --profiles stringArray   Profiles to use
  -s, --sshkey string          Path to ssh private key authorized for HOST
  -u, --user string            Container username (default "ubuntu")
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

* [dlx](/docs/cmd/dlx)	 - 

