---
layout: bootstrap_post
title: EC2 Placement Groups
date: 2015-09-04 13:13:13
author: Oz Akan
abstract: Wonders of Placement Groups and Enchanced Networking
categories:
    - AWS
tags:
    - aws, networking, performance, hpc
    
---

If you have an application that requires best possible network performance between instances, AWS provides two features you can benefit from. *Placement groups* and *enchanced networking*.

> A placement group is a logical grouping of instances within a single Availability Zone. Using placement groups enables applications to participate in a low-latency, 10 Gbps network.

These instances that are part of same placement group will be in the same availibility zone, they will be on the 10 Gbps network which will provide very low latency and high network throughput. To get the most from the underlying network infrastructure, you will have to choose instances from C3, C4, D2, I2, M4, R3 families as only these instance families support enchanced networking.    

> Amazon EC2 provides enhanced networking capabilities using single root I/O virtualization (SR-IOV) on supported instance types. Enabling enhanced networking on your instance results in higher performance (packets per second), lower latency, and lower jitter.

I wanted to see latency and throughput between instances created within a placement group, without a placement group and in different zones.

I dediced to use two different instance types in order to compare enchanced networking with non-enchanced networking.

- c4.large: has enchanced networking, supports placement groups.
- m3.large: doesn't have enchanced networking 

I ran a simple ping command for latency and iperf for throughput. 

- Test A: c4.large, enchanced networking, without placement group, same zone.
- Test B: c4.large, enchanced networking, without placement group, different zones
- Test C:  c4.large, enchanced networking, with placement group, same zone
- Test D: m3.large, no enchanced networkgin, no placement group, same zone

-

|                        | Test A  | Test B  | Test C | Test D |
|------------------------|--------:|--------:|-------:|-------:|
| latency (ms)           | 0.132   | 0.778   | 0.129  | 0.464  |
| throughput (Gbits/sec) | 617     | 617     | 617    | 716    |

-



As seen c4.large instances within the same placement group has the best latency value which is just little better than same instance type without a placement group. 

Throughput is same even when two instances are in different zones and seems to be limited with the instance size. Larger instances would achieve even better throughput (c4.xlarge does a little more than 1 Gbits/sec)


<sub>
sources:
[Placement Groups](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) | 
[Enabling enhanced networking](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/enhanced-networking.html) | 
[Pricing](https://aws.amazon.com/ec2/pricing/)

</sub>