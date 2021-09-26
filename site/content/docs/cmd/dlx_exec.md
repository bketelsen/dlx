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

