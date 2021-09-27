---
title: dlx completion powershell
description: dlx completion powershell
lead: dlx completion powershell
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
## dlx completion powershell

generate the autocompletion script for powershell

### Synopsis


Generate the autocompletion script for powershell.

To load completions in your current shell session:
PS C:\> dlx completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
dlx completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx completion](/docs/cmd/dlx_completion)	 - generate the autocompletion script for the specified shell

