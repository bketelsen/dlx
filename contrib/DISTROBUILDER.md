# Distrobuilder Template

This directory contains the template I use to build the base image for all my containers. It's nearly unmodified from base Ubuntu image provided in the official [lxc-ci](https://github.com/lxc/lxc-ci) project. I modified it to create a user with my preferred username (bjk) instead of the default (ubuntu). That ensures that my user is UID/GID 1000. I also added the Tailscale apt repository to the list of sources and added the `tailscale` package to the list of packages to install. You can use those two sections as templates for adding other repositories and packages to your base install.

In the future I will likely add a tailscale access token to the base image and configure it to automatically join my tailscale network, for convenience. For now, I simply run `sudo tailscale up` to join the network the first time I start the container.

You can also modify things like the default profile by adding files in the `/etc/profile.d` directory. 

## Tradeoffs

You could install all your favorite things in this base image. That would increase the size of the base, but decrease your time provisioning a new development container with the tools required for each project.

I think I'm going to settle with a happy medium and add shell scripts to the base image to install the tools I need for each project. Then the base isn't bloated, and I don't have to hunt to find installation scripts for every project I work on.