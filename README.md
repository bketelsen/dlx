# lxdev

`lxdev` is a development tool that provisions temporary development environments.  It uses [lxd](https://linuxcontainers.org) and `zfs` to make efficient, copy-on-write, workspaces from a user-provided template.


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

```
lxdev provisions lxd containers for local development.

Usage:
  lxdev [command]

Available Commands:
  connect     connect to a running container
  create      Create a container
  help        Help about any command
  list        list containers
  remove      remove a container
  start       start a paused container
  stop        stop a running container
  version     version of lxdev

Flags:
      --config string   config file (default is $HOME/.lxdev.yaml)
  -h, --help            help for lxdev
  -t, --toggle          Help message for toggle

Use "lxdev [command] --help" for more information about a command.
```

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
