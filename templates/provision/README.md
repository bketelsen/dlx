# Provision

I have added a collection of scripts to my home directory on the host hachine in $HOME/provision.

This directory is mounted into the guest machine at $HOME/provision in the `devices` script.

A template for the `devices` script is included in this repository under `/scripts`.

This is a completely optional step to give me fast access to shell scripts to install things I commonly need in the containers.

## cloud-init

Note that you could accomplish the same thing with cloud-init, but I prefer to use the scripts in this repository. `cloud-init` is a great way to provision containers, but it runs in the background and doesn't provide any feedback to the user while it is provisioning. So the container might be available for connections, but your provisioning might not be complete, which could cause problems if you've provisioned paths in your dotfiles, or special commands in your shell that aren't yet installed. I've found it easier to just install them explicitly.