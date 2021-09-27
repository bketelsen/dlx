## dlx completion bash

generate the autocompletion script for bash

### Synopsis


Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:
$ source <(dlx completion bash)

To load completions for every new session, execute once:
Linux:
  $ dlx completion bash > /etc/bash_completion.d/dlx
MacOS:
  $ dlx completion bash > /usr/local/etc/bash_completion.d/dlx

You will need to start a new shell for this setup to take effect.
  

```
dlx completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx completion](/docs/cmd/dlx_completion)	 - generate the autocompletion script for the specified shell

