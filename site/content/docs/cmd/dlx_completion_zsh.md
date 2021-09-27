---
title: dlx completion zsh
description: dlx completion zsh
lead: dlx completion zsh
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
## dlx completion zsh

generate the autocompletion script for zsh

### Synopsis


Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:
# Linux:
$ dlx completion zsh > "${fpath[1]}/_dlx"
# macOS:
$ dlx completion zsh > /usr/local/share/zsh/site-functions/_dlx

You will need to start a new shell for this setup to take effect.


```
dlx completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx completion](/docs/cmd/dlx_completion)	 - generate the autocompletion script for the specified shell

