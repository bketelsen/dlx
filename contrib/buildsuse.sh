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
	./suse.yaml

#sudo distrobuilder build-lxd -o image.architecture=x86_64 -o image.release=tumbleweed -o image.variant=default suse.yaml opensuse
#pushd opensuse
#lxc image import lxd.tar.xz rootfs.squashfs --alias dlxbase
#popd

