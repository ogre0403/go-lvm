#!/bin/sh

# Install required package
apk update
apk add libblockdev-dev lvm2 gcc make libc-dev

# setup lo device
dd if=/dev/zero of=/tmp/disk.img bs=1G count=1
LOOP=/dev/loop3
losetup $LOOP /tmp/disk.img

# run go test
cd /tmp/sample
go test -v .

# remove lo device
losetup -d $LOOP
