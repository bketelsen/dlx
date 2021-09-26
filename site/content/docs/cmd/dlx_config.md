## dlx config

manage global configurations

### Synopsis

manage global configurations

```
dlx config [flags]
```

### Options

```
  -b, --baseimage string       Default base image for new containers (default "dlxbase")
  -t, --clientcert string      Path to client certificate
  -k, --clientkey string       Path to client key
  -c, --create                 Create global config file
  -h, --help                   help for config
  -p, --profiles stringArray   Profiles to use
  -r, --remote string          LXD host network name or IP
  -s, --sshkey string          Path to ssh private key authorized for HOST
  -u, --user string            Container username (default "ubuntu")
```

### Options inherited from parent commands

```
  -v, --verbose   verbose logging
```

### SEE ALSO

* [dlx](/docs/cmd/dlx)	 - Provision lxd containers for development

