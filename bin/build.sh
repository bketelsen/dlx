#!/bin/bash

# arguments
# $1: yaml file
# $2: image alias

sudo distrobuilder build-lxd $1
lxc image import lxd.tar.xz rootfs.squashfs --alias $2

