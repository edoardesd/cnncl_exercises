#! /bin/sh
# how to run it: sh run_qemu.sh

KERNEL=linux-6.0.1
BZIMAGE=arch/x86/boot/bzImage
HELLO_SOURCE=hello.c
HELLO_EXE=hello

KERNEL_TAR="$KERNEL.tar.gz"

echo "===== Exercise 1 - QEMU FILESYSTEM ====="
sleep 1

if [ ! -f "$KERNEL_TAR" ]; then
  echo "Downloading kernel source..."
  wget https://cdn.kernel.org/pub/linux/kernel/v6.x/$KERNEL_TAR
  tar xvf "$KERNEL_TAR"
fi

if [ ! -f "$KERNEL/$BZIMAGE" ]; then
  echo "Making the kernel... This can take a while, grab a cuppa."
  sleep 1
  make -C "$KERNEL" ARCH=x86_64 defconfig
  make -C "$KERNEL" -j 8
fi

if [ ! -f "$HELLO_SOURCE" ]; then
  echo "Creating the hello world c file..."
  touch "$HELLO_SOURCE"
  echo "#include <stdio.h>" >> "$HELLO_SOURCE"
  echo "void main() { printf(\"Hello World!\"); while(1); }" >> "$HELLO_SOURCE"
fi

if [ ! -f "$HELLO_EXE" ]; then
  echo "Compiling hello world..."
  gcc -static "$HELLO_SOURCE" -o "$HELLO_EXE"
  echo "... and creating the rootfs"
  echo "$HELLO_EXE" | cpio -o --format=newc > rootfs

fi

echo "Running the file system!"
sleep 1
qemu-system-x86_64 -kernel $KERNEL/$BZIMAGE \
                    -initrd rootfs \
                    -append "root=/dev/ram rdinit=/$HELLO_EXE console=ttyS0" \
                    -serial stdio \
                    -display none