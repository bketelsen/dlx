# lxdev

`lxdev` is a development tool that provisions temporary development environments.  It uses [lxd](https://linuxcontainers.org) and `zfs` to make efficient, copy-on-write, workspaces from a user-provided template.

[Watch this slightly outdated DEMO VIDEO](https://youtu.be/W6A00CHiDQ8)

## Getting Started

* Install LXD
* Install lxdev

### Create configuration and template files

```
lxdev config -c
lxdev config -t
```

These commands write `$HOME/.lxdev.yaml` and `$HOME/.lxdev/profiles/*.yaml`, which are configuration files and templates for your containers.

### Create Templates

```
lxdev template create guitemplate --profile gui --provisioners vscode
lxdev template create clitemplate --profile cli --provisioners go,yadm
```
Let's unwrap that:

The name of the template {guitemplate,clitemplate} is totally up to you.  These are base images that will be used later to create your containers.  You "provision" them by passing in a comma separated list of `provisioners`, which are bash scripts that install things or otherwise modify the base image.  Provisioners live in the `~/.lxdev/provision` directory in your $HOME.  They're created once and never again modified by `lxdev` unless you remove the directory and run `lxdev config -t` again.

The `guibase` and `clibase` provisioning templates are automatically applied to `gui` and `cli` profiles, you do not need to specify them separately.  Use caution in editing these provisioners, as it is possible features installed in these provisioners are expected by `lxdev`.

You can, and should, modify the existing provisioners or create new ones based on your needs.

The `profile` {gui,cli} is an lxc profile that's stored in `~/.lxdev/profiles`.  They're standard `lxc` profiles that are applied when you create a template, then inherited by every container that's instantiated from those templates.


### Create your first container

```
lxdev create myproject --template guitemplate
```
This creates a container called `myproject` from the template `guitemplate`, which has X11 and audio support by default.

### Connect to your container

```
lxdev shell myproject
```

When using the shell (or its alias connect) command, you get dropped into a login shell in the container.  You can run commands just like it was an SSH session, and you can open X11 apps which will be displayed on your host's X session.  (I KNOW RIGHT??)

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
Special thanks to [Simos Xenitellis](https://blog.simos.info/how-to-easily-run-graphics-accelerated-gui-apps-in-lxd-containers-on-your-ubuntu-desktop/) for the tireless blogging. I learned nearly everything about this process from those posts.
