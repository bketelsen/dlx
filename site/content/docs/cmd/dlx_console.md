---
title: dlx console
description: dlx console
lead: dlx console
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
## dlx console

Attach to instance consoles

### Synopsis

Description:
  Attach to instance consoles

  This command allows you to interact with the boot console of an instance
  as well as retrieve past log entries from it.



```
dlx console [<remote>:]<instance> [flags]
```

### Options

```
      --show-log   Retrieve the instance's console log
  -t, --type       Type of connection to establish: 'console' for serial console, 'vga' for SPICE graphical output (default "console")
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

