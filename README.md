go-lvm
=======================================================================

## Overview

go-lvm is a go library to call lib2app API to manipulate [LVM2](https://sourceware.org/lvm2/).

## Required package
```bash
$ sudo yum install device-mapper-devel
$ sudo yum install lvm2-devel
```


## CLI Example

```bash
# Step-1. create a loop device partition
$ sudo dd if=/dev/zero of=disk.img bs=1G count=1
$ export LOOP=`sudo losetup -f`
$ sudo losetup $LOOP disk.img

# Step-2. create PV for loop device
$ pvcreate $LOOP

# Step-3. create VG for new created PV
$ vgcreate vg-0 $LOOP

# Step-4. create 10MB LV 
$ lvcreate -n lv01 -L10M vg-0
$ ll /dev/vg-0/lv01

# Step-5. remove LV
$ lvremove vg-0/lv01

# Step-6. remove VG
$ vgremove vg-0

# Step-7. remove PV
$ pvremove $LOOP
```
### Thin provision example
Reference: http://manpages.ubuntu.com/manpages/xenial/man7/lvmthin.7.html

```bash
# Create a LV for data of thin pool
$ lvcreate -n pool0 -L 1G vg-0

# Create a LV for metadata
$ lvcreate -n pool0meta -L 100M vg-0

# Combine the data and metadata LVs into a thin pool LV.
$ lvconvert --type thin-pool --poolmetadata vg-0/pool0meta vg-0/pool0

# Create thin LV within pool LV
$ lvcreate -n thin1 -V 2G --thinpool vg-0/pool0
```


## Test run

Let's create a available volume group and create and delete a LV.

### step-1. set up a free VG
```bash
sudo dd if=/dev/zero of=disk.img bs=1G count=1
export LOOP=`sudo losetup -f`
sudo losetup $LOOP disk.img
```

### step-2. Run an example script
```bash
go run cmd/example.go
```