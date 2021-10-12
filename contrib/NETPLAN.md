# Netplan

This is a template for a netplan configuration file that creates a bridge from the primary ethernet interface.

I use this on my HOME network to give each container an IP address on the network, rather than the default lxdbr0 NAT. In cloud environments, I just add Tailscale to the containers when I need external access.

It removes DHCP from the ethernet interface and adds DHCP to the bridge.  If you want to use this recipe, you will need to ensure that you modify the ethernet interface to match your device name. My primary ethernet interface is called eno1.

** Special Note **: The `renderer` section of this configuration is very important. Different versions of Ubuntu use different methods of applying network configuration. Ubuntu Server uses `networkd` and Ubuntu Desktop flavors frequently use `NetworkManager`. Don't blindly copy this or you will end up with a broken network configuration. 

```yaml
network:
  version: 2
  renderer: networkd

  ethernets:
    eno1:
      dhcp4: no

  bridges:
    br0:
      dhcp4: yes
      interfaces:
        - eno1
```