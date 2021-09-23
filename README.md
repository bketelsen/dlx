# dlx

`dlx` is a development tool that provisions temporary development environments.  It uses [lxd](https://linuxcontainers.org) and [zfs](https://wiki.ubuntu.com/ZFS) to make efficient, copy-on-write, workspaces from a user-provided template.

## Assumptions

* provision directory
* devices script
