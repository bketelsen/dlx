---
title: dlx exec
description: dlx exec
lead: dlx exec
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
## dlx exec

Execute a command in a container

### Synopsis

Executes a command in the named container.  The command should be enclosed in 
single quotes.  e.g. exec mycontainer 'ls -la'

```
dlx exec [container] '[commands here]' [flags]
```

### Options

```
  -h, --help   help for exec
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx](/docs/cmd/dlx)	 - Provision lxd containers for development

