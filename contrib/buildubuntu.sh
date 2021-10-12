#!/bin/bash
set -e

SED="sed"
if which gsed >/dev/null 2>&1; then
	SED="gsed"
fi
USERNAME=`whoami`

"$SED" \
	-i'' \
	-e "s/dlxuser/${USERNAME}/g" \
	./ubuntu.yaml


sudo distrobuilder build-lxd ubuntu.yaml ubuntu
pushd ubuntu
lxc image import lxd.tar.xz rootfs.squashfs --alias dlxbase
popd

