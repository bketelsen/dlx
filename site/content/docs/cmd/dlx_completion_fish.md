---
title: dlx completion fish
description: dlx completion fish
lead: dlx completion fish
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
## dlx completion fish

generate the autocompletion script for fish

### Synopsis


Generate the autocompletion script for the fish shell.

To load completions in your current shell session:
$ dlx completion fish | source

To load completions for every new session, execute once:
$ dlx completion fish > ~/.config/fish/completions/dlx.fish

You will need to start a new shell for this setup to take effect.


```
dlx completion fish [flags]
```

### Options

```
      --no-descriptions   disable completion descriptions
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

* [dlx completion](/docs/cmd/dlx_completion)	 - generate the autocompletion script for the specified shell

