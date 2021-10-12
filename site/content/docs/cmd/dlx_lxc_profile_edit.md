---
title: dlx lxc profile edit
description: dlx lxc profile edit
lead: dlx lxc profile edit
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
## dlx lxc profile edit

Edit profile configurations as YAML

### Synopsis

Description:
  Edit profile configurations as YAML



```
dlx lxc profile edit [<remote>:]<profile> [flags]
```

### Examples

```
  lxc profile edit <profile> < profile.yaml
      Update a profile using the content of profile.yaml
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

* [dlx lxc profile](/docs/cmd/dlx_lxc_profile)	 - Manage profiles

