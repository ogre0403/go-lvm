#!/bin/sh

apk update
apk add libblockdev-dev lvm2 gcc make libc-dev

dd if=/dev/zero of=/tmp/disk.img bs=1G count=1
LOOP=`losetup -f`
losetup $LOOP /tmp/disk.img
pvcreate $LOOP
vgcreate vg-0 $LOOP

lvcreate -n pool0 -L 500M vg-0
lvcreate -n pool0meta -L 50M vg-0
lvconvert --yes --type thin-pool --poolmetadata vg-0/pool0meta vg-0/pool0



cd /tmp/sample
go test -v .

lvremove --yes vg-0/pool0
vgremove --yes vg-0
pvremove --yes $LOOP
losetup -d $LOOP
