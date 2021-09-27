---
title: dlx config
description: dlx config
lead: dlx config
date: 2021-09-27T07:23:53-04:00
lastmod: 2021-09-27T07:23:53-04:00
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
  -h, --help                   help for config
  -p, --profiles stringArray   Profiles to use
  -s, --sshkey string          Path to ssh private key authorized for HOST
  -u, --user string            Container username (default "ubuntu")
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx](/docs/cmd/dlx)	 - Provision lxd containers for development

