---
layout: bootstrap_post
title: KVM, virt-install and virsh
date: 2015-06-16 09:25:00
author: Oz Akan
abstract: It has been a while since the last time there wasn't an API between me and the infrastructure. You might be on the same boat and want to have a reference to do a few tricks with KVM.
categories:
    - Cloud
tags:
    - kvm
    - cloud
    - virtualization
---

## Create An Instance of CentOS 7

I decided to use `/home/qemu` as my work folder. `qemu` user is used by kvm so it is convient.

Download CentOS 7 image under `/home/qemu/ISO`

### Create QCOW2

    # qemu-img create -f qcow2 -o preallocation=metadata /home/qemu/images/centos7-base.qcow2 10G
    
Convert to sparse so it will only allocate about 1GB instead of 11GB.

    # mv centos7-base.qcow2 centos7-base.qcow2_backup
    # qemu-img convert -O qcow2 centos7-base.qcow2_backup centos7-base.qcow2

Run the command below to start boot process.

    # virt-install \
     --network bridge=br1700,virtualport_type=openvswitch \
     --name=centos7-base \
     --ram=1024 \
     --vcpus=1 \
     --os-type=linux \
     --os-variant=rhel7 \
     --disk path=/home/qemu/images/centos7-base.qcow2,format=qcow2 \
     --cdrom /home/qemu/ISO/cent7.iso \
     --graphics vnc,password=mypassword,listen=0.0.0.0 \
     --noautoconsole

Connect to the instance to finish boot process. You can do a nogui installation as well.

On mac you can use `screen sharing` application or any other VNC client. (I don't know the real vnc client was crashing on mac and logs didn't say much) Simply type the command below which opens up `screen sharing` application.

    open vnc://192.168.1.2:5900

Enter the password `mypassword` and you will be in.

Once the installation complete, we want to create a base image in order to use for other virtual machines.

So poweroff the instance

    # poweroff
    
On the host 

    # virsh list
     Id    Name                           State
    ----------------------------------------------------

    # virsh list --all
     Id    Name                           State
    ----------------------------------------------------
     -     centos7-base                   shut off

    

## Creating New Instance By Cloning Base Image

`/home/qemu/images/centos7-base.qcow2` is the base image we have. It is powered off.

    # virsh list --all
     Id    Name                           State
    ----------------------------------------------------
     -     centos7-base                   shut off

### Clone centos7-base instance

    virt-clone \
     --connect qemu:///system \
     --original centos7-base \
     --name core02 \
     --file /home/qemu/images/core02.qcow2

It will take a bit of time to copy the image. KVM will update it's database and create an entry for `core02`. Let's check it.

    # virsh list --all
     Id    Name                           State
    ----------------------------------------------------
     -     centos7-base                   shut off
     -     core02                         shut off

If you want to see configuration for `core02` this is the command;

    # virsh dumpxml core02
    <domain type='kvm'>
      <name>core02</name>
      <uuid>4fcf71ce-3a09-427a-b386-dcf739d221e3</uuid>
      <memory unit='KiB'>1048576</memory>
      ....
      <devices>
        <emulator>/usr/libexec/qemu-kvm</emulator>
          <disk type='file' device='disk'>
            <driver name='qemu' type='qcow2'/>
            <source file='/home/qemu/images/core02.qcow2'/>
        ...
        <memballoon model='virtio'>
          <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
        </memballoon>
      </devices>
    </domain>

This will give a long xml file with every detail about the vm.

### Boot VM

    # virsh start core02
    Domain core02 started

### Find out VNC port

    # virsh vncdisplay core02
    :0

{% for count in (0..3) %}
 `:{{ count }}` means port `590{{ count }}`
{% endfor %}

...and it goes like that. If you have several virtual machines running you would see more of these ports being allocated.

### Connect With a VNC Client

Connect to the IP address of the host server and port `5900`.

    # open vnc://ip.off.the.server:5900

### Change the hostname

In two locations

    # vi /etc/hostname
    hostname.domain.com

    # vi /etc/hosts    
    127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
    ::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
    127.0.0.1 hostname.domain.com hostname
    
Restart networking

    # /etc/init.d/network restart

### Assign IP Address

    # ip addr add 10.100.8.100/24 dev eth0

## Conclusion

I don't think it is fun to set IP address manually, unless you are a `human keyboard`! We are not input devices, we are the makers. Right? So we are not done, expect for a follow up post soon...
