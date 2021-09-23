# Profile examples

The lxd profiles provided in this directory are examples only. If you use them blindly, you will most definitely break things.

## Default Profile

This profile shows how to map users from the host to the containers. I've added it to my default profile because I'm always going to have user 1000 in the container mapped to my user (which is UID and GID 1000) on the host. Not every flavor of linux uses UID 1000, so this won't work in every scenario.

```yaml
config:
  raw.idmap: |
    both 1000 1000
```

## Bridge Profile

I added an ethernet bridge named `br0` to my host which bridges host and container ethernet interfaces. This provides containers IP addresses assigned by my network's DHCP server and makes them accessible from any computers on my network. 

This change to LXD's default isn't required, but I've found it to be the single most useful part of this process. Not only can I connect to the containers from `dlx connect` but I can also just use ssh to connect to them from anywhere on my network.

If you want to do something similar, use this profile as a starting point, being sure to change the name of the bridge to match what you have on your host.

```yaml
devices:
  eth0:
    name: eth0
    nictype: bridged
    parent: br0
    type: nic
```