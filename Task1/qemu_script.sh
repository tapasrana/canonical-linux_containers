#!/bin/bash

# Step 1: Download or build a Linux kernel (In this example, we'll download a pre-built kernel)
KERNEL_VERSION="5.15"
wget https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-${KERNEL_VERSION}.tar.xz
tar -xvf linux-${KERNEL_VERSION}.tar.xz

# Step 2: Create a minimal root filesystem (using BusyBox)
git clone git://busybox.net/busybox.git
cd busybox
make defconfig
make
make install
cd ..
mkdir rootfs
cp -rf busybox/_install/* rootfs/

# Step 3: Create a script to print "hello world"
echo -e '#!/bin/sh\n\necho "hello world"' > rootfs/init

# Step 4: Combine kernel and root filesystem into a bootable image
cd linux-${KERNEL_VERSION}
make defconfig
make bzImage
cd ..

# Step 5: Run the image using QEMU
qemu-system-x86_64 -kernel linux-${KERNEL_VERSION}/arch/x86/boot/bzImage -append "root=/dev/ram rdinit=/init console=ttyS0" -initrd rootfs/init -nographic
