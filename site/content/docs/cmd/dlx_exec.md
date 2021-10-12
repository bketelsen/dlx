---
title: dlx exec
description: dlx exec
lead: dlx exec
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
## dlx exec

Execute commands in instances

### Synopsis

Description:
  Execute commands in instances

  The command is executed directly using exec, so there is no shell and
  shell patterns (variables, file redirects, ...) won't be understood.
  If you need a shell environment you need to execute the shell
  executable, passing the shell commands as arguments, for example:

    dlx exec <instance> -- sh -c "cd /tmp && pwd"

  Mode defaults to non-interactive, interactive mode is selected if both stdin AND stdout are terminals (stderr is ignored).



```
dlx exec <instance> [flags] [--] <command line>
```

### Options

```
      --cwd                    Directory to run the command in (default /root)
  -n, --disable-stdin          Disable stdin (reads from /dev/null)
      --env                    Environment variable to set (e.g. HOME=/home/foo)
  -t, --force-interactive      Force pseudo-terminal allocation
  -T, --force-noninteractive   Disable pseudo-terminal allocation
      --group                  Group ID to run the command as (default 0)
      --mode                   Override the terminal mode (auto, interactive or non-interactive) (default "auto")
      --user                   User ID to run the command as (default 0)
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

