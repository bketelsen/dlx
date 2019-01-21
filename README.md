# lxdev

`lxdev` is a development tool that provisions temporary development environments.  It uses [lxd](https://linuxcontainers.org) and `zfs` to make efficient, copy-on-write, workspaces from a user-provided template.

[Watch this DEMO VIDEO](https://youtu.be/W6A00CHiDQ8)

## Getting Started

* Install LXD
* Install lxdev

### Create configuration and template files

```
lxdev config -c
lxdev config -t
```

These commands write `$HOME/.lxdev.yaml` and `$HOME/.lxdev/profiles/*.yaml`, which are configuration files and templates for your containers.

### Apply the profiles

```
lxdev profile -w gui
lxdev profile -w util
lxdev profile -w cli
lxdev profile -w go
```

These commands write the container profiles that you'll apply to new containers you create.  Each container created must have one of `gui,util,cli` as it's base profile.  Extra profiles can be specified (the `go` profile is an extra one)

### Create your first container

```
lxdev create myproject
```

With no additional flags, this project will get the `default` (lxc provided) profile, plus the `gui` profile for X11 support.  You can specify other profiles:

```
lxdev create myproject --profiles 'go'
```

This will get `default`, `gui` and `go`.

```
lxdev create slack --util
```

This will get the `util` profile, which I intend for use with things like Slack, Firefox, etc.



### Prerequisites

You need a Linux development system, with LXD at a recent version.  Developed and tested on Ubuntu 18.10, but LXD targets any Linux installation with a modern kernel.


### Installing

#### Download a Release

When there is one, you can download a release from GitHub.

#### Build From Source

Requires Go, tested with 1.12beta2.
```
git clone https://github.com/bketelsen/lxdev

make deps // install dependencies
make test  // run tests
make install // build and install the lxdev tool into your path
```

## Running the tests

```
make test
```

## Built With

* [cobra](http://github.com/spf13/cobra/) - The easiest way to make command line tools in Go

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/bketelsen/lxdev/tags).

## Authors

* **Brian Ketelsen** - *Initial work* - [BrianKetelsen.com](https://brianketelsen.com)

See also the list of [contributors](https://github.com/bketelsen/lxdev/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

See [SHOULDERS](SHOULDERS.md) for acknowledgments and thanks to the other projects that `lxdev` was built with.
