# dlx

`dlx` is a development tool that provisions temporary development environments.  It uses [lxd](https://linuxcontainers.org) and [zfs](https://wiki.ubuntu.com/ZFS) to make efficient, copy-on-write, workspaces from a user-provided template.

Turn a spare computer into a development environment server in just a few minutes.

## Why dlx?

I built `dlx` because I bounce between laptops frequently and I wanted a consistent development environment without having to use something like Syncthing to sync my code. I had an older Intel NUC in the closet doing not very much work, and I wanted to find a way to use it for development.

## Components

`dlx` is composed of the following components:

	* [lxd](https://linuxcontainers.org) - the container management system that runs on a remote host, either locally on your network or in the cloud.
	* [dlx](https://github.com/bketelsen/dlx) - client program that runs on your local computer.
	* distrobuilder - a tool that builds a Linux container image from a template
	* Optionally [zfs](https://wiki.ubuntu.com/ZFS)
	* Optionally [tailscale](https://www.tailscale.com)

LXD is a containerization technology that is used to create a container that can be used to run a development environment or longer running services.  It is a lightweight, easy to use, and highly available containerization technology. LXD has several different options for storage, the best of which for these purposes is [zfs](https://wiki.ubuntu.com/ZFS). [btrfs](https://btrfs.wiki.kernel.org/index.php/Main_Page) is also supported. Either will be significantly faster than other LXD storage options, and either supports Copy on Write, which is a feature that is used to make the container's filesystems more efficient.

`dlx` is a CLI client that uses the LXD API to provision a development environment.  

As a convenience, all of the `lxc` client commands are available as subcommands in `dlx`.  For example, `dlx lxc list` is equivalent to `lxc list`. This is possible because the `lxc` client code is written in Go and also uses the [cobra](https://cobra.dev) library. Whether this is a good idea or not remains to be seen. For now, it's a nice convenience because it means you're not required to install the `lxc` client on your local development machine.

## How does it work?

`dlx` uses [lxd](https://linuxcontainers.org) to provision a new container from a custom Linux image. A typical workflow might look like this:

```bash
dlx create projectname
```

`dlx` instructs the LXD daemon to create a container called `projectname` from the image `dlxbase`. `dlxbase` is a convention that is used to refer to a base LXD image that exists locally on the LXD server. The documentation includes information on how to edit the template image and create it on the LXD server. This `dlxbase` image is used to create all containers created by `dlx`.

`dlx` also creates a `projectname` directory on the LXD host filesystem and mounts it into the container in the `/home/user` directory. This is where your source code should be stored. This allows you to backup your source code from the host using conventional Linux backup tools.

Depending on your needs, you may configure LXD to bridge network traffic between the containers and your local network, giving each container an IP address on your LAN. This is how I use `dlx` at home. You can also use something like [tailscale](https://www.tailscale.com) to provide connectivity to the containers from your VPN network, which would allow you to use `dlx` to provision a development environment on a remote host that is not on your LAN and can't provide routeable IP addresses.

After the container is provisioned, you can access it over SSH using the IP address of the container. This works perfectly for [VS Code's](https://code.visualstudio.com/) remote SSH feature. If you've configured bridged traffic, the containers appear as peers on your local network, so you can use `ssh` to connect to them directly. If you're using Tailscale, you can setup Magic DNS to achieve the same effect. Note: Tailscale isn't required for this to work, you can use Wireguard or other VPN solutions, or you can use LXD's built-in port forwarding features to access your containers.

When you're done with the container, you can stop it with `dlx stop projectname`, and optionally delete it with `dlx rm projectname`. The host's `projectname` directory is not deleted, which means that you can later create a new container with the same project name and get the same source code.

### Conventions and Assumptions

`dlx` makes several assumptions about your setup. Not all of these are required, but they are recommended.

	* The LXD server is running Ubuntu 18.04 or higher.
	* LXD is installed via `snap` on the server.
	* Your account on the server has sudo privileges.
	* Your account on the server has public and private keys in `~/.ssh`. These are mounted in the containers to allow you to SSH into the containers.
	* You have passwordless SSH access to the LXD server.
	* Your base LXD image `dlxbase` is configured to create a user with the same username you use on your laptop/desktop.

Using these conventions allows `dlx` to create a container that requires no extra work to access.


## Installation

See the [documentation](https://dlx.rocks)

## Usage Scenarios

### Locally

Locally I have an LXD server running headless in my closet. I have LXD using a bridge to connect the containers to my local network. When I create a new container, they're immediately available on my local network by using the container name as the hostname.

### In the Cloud

I also have an instance of LXD running in the cloud. The LXD server is using the default `lxdbr0` bridge to route traffic to and from the containers. The `dlxbase` image is configured to install Tailscale in each container. After the container is provisioned I use `dlx connect` to access the container and run `tailscale up` to authorize the container (this could be automated with a Tailscale Personal Access Token, but I haven't done this yet). Because I have Magic DNS configured on my Tailscale account, the containers are available by hostname in DNS lookups. SSH and VS Code Remote SSH work the same as they do locally.


## Security

## Contributing

## License

## Documentation

Documentation (Work in progress!) is available at [dlx.rocks](https://dlx.rocks/).