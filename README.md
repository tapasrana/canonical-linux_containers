# canonical-linux_containers
This repo contains the 2 tasks for Canonical Embedded Linux Containers Software Engineer position.
Task 1:
Overview of the qemu_script.sh ➖
1] Download the latest kernel from source and build and make the kernel.
2] Create, build and install a minimal root file system using Busybox
3] Combine the kernel and root file system into a bootable image
4] Run the bootable image using qemu emulator

Task 2:
Here, the test cases that I have implemented for the Shred function are a random data which
generates integer values randomly. This makes sure that the randomfile.txt is overwritten with
new random data every time the file is overwritten, thus making sure that the overwritten text is
not the same with each iteration. In addition to that, I have also added a text snippet for every
overwrite sequence just so that alphabets are also included along with integers. This makes
sure the file is overwritten with a variety of text strings.
Code Description:
generate_random_data() : This function generates random data of a particular size to be
overwritten into the file..
Iteration : array of string texts to write into the file. Each string snippet is written for each
iteration.
shred(): This function takes a path argument, which is the path to the file you want to shred. It
overwrites the file three times with random data and then deletes the file using os.Remove(path)
