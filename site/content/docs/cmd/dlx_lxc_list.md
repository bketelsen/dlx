---
title: dlx lxc list
description: dlx lxc list
lead: dlx lxc list
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
## dlx lxc list

List instances

### Synopsis

Description:
  List instances

  Default column layout: ns46tS
  Fast column layout: nsacPt

  A single keyword like "web" which will list any instance with a name starting by "web".
  A regular expression on the instance name. (e.g. .*web.*01$).
  A key/value pair referring to a configuration item. For those, the
  namespace can be abbreviated to the smallest unambiguous identifier.
  A key/value pair where the key is a shorthand. Multiple values must be delimited by ','. Available shorthands:
    - type={instance type}
    - status={instance current lifecycle status}
    - architecture={instance architecture}
    - location={location name}
    - ipv4={ip or CIDR}
    - ipv6={ip or CIDR}

  Examples:
    - "user.blah=abc" will list all instances with the "blah" user property set to "abc".
    - "u.blah=abc" will do the same
    - "security.privileged=true" will list all privileged instances
    - "s.privileged=true" will do the same
    - "type=container" will list all container instances
    - "type=container status=running" will list all running container instances

  A regular expression matching a configuration item or its value. (e.g. volatile.eth0.hwaddr=00:16:3e:.*).

  When multiple filters are passed, they are added one on top of the other,
  selecting instances which satisfy them all.

  == Columns ==
  The -c option takes a comma separated list of arguments that control
  which instance attributes to output when displaying in table or csv
  format.

  Column arguments are either pre-defined shorthand chars (see below),
  or (extended) config keys.

  Commas between consecutive shorthand chars are optional.

  Pre-defined column shorthand chars:
    4 - IPv4 address
    6 - IPv6 address
    a - Architecture
    b - Storage pool
    c - Creation date
    d - Description
    D - disk usage
    l - Last used date
    m - Memory usage
    M - Memory usage (%)
    n - Name
    N - Number of Processes
    p - PID of the instance's init process
    P - Profiles
    s - State
    S - Number of snapshots
    t - Type (persistent or ephemeral)
    u - CPU usage (in seconds)
    L - Location of the instance (e.g. its cluster member)
    f - Base Image Fingerprint (short)
    F - Base Image Fingerprint (long)

  Custom columns are defined with "[config:|devices:]key[:name][:maxWidth]":
    KEY: The (extended) config or devices key to display. If [config:|devices:] is omitted then it defaults to config key.
    NAME: Name to display in the column header.
    Defaults to the key if not specified or empty.

    MAXWIDTH: Max width of the column (longer results are truncated).
    Defaults to -1 (unlimited). Use 0 to limit to the column header size.



```
dlx lxc list [<remote>:] [<filter>...] [flags]
```

### Examples

```
  lxc list -c nFs46,volatile.eth0.hwaddr:MAC,config:image.os,devices:eth0.parent:ETHP
    Show instances using the "NAME", "BASE IMAGE", "STATE", "IPV4", "IPV6" and "MAC" columns.
    "BASE IMAGE", "MAC" and "IMAGE OS" are custom columns generated from instance configuration keys.
    "ETHP" is a custom column generated from a device key.

  lxc list -c ns,user.comment:comment
    List instances with their running state and user comment.
```

### Options

```
  -c, --columns   Columns (default "ns46tSL")
      --fast      Fast mode (same as --columns=nsacPt)
  -f, --format    Format (csv|json|table|yaml) (default "table")
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

* [dlx lxc](/docs/cmd/dlx_lxc)	 - Command line client for LXD

