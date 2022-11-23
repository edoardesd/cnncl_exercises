#! /bin/sh
# how to run it: sh run_qemu.sh

KERNEL=linux-6.0.1
BZIMAGE=arch/x86/boot/bzImage
if [ ! -f "$KERNEL.tar.gz" ]; then
  echo -n "Downloading kernel source..."
  wget https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-6.0.1.tar.gz
  tar xvf linux-6.0.1.tar.gz
fi

if [ ! -f "$KERNEL/$BZIMAGE" ]; then
  echo -n "Making the kernel..."
  make -C linux-6.0.1 ARCH=x86_64 defconfig
  make -C linux-6.0.1 -j 8
fi