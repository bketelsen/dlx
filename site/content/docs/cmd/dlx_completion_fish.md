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
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx completion](/docs/cmd/dlx_completion)	 - generate the autocompletion script for the specified shell

