---
layout: bootstrap_post
title: VPC Limitations
date: 2015-08-25 13:13:13
author: Oz Akan
abstract: Limitations around VPC service.
categories:
    - AWS
tags:
    - aws, restrictions, vpc, limits 
    
---

## 1
Internet gateway per VPC

Virtual private gateways per VPC

Expiry time in weeks for an unaccepted VPC peering connection request.

## 2
Flow logs per single network interface, single subnet, or single VPC in a region

## 5
VPC per region, can be increased upon request

Elastic IP addresses per region for each AWS account, can be increased upon request

Security groups per network interface, can be increased upon request to max 16

## 20
Rules per network ACL

VPC endpoints per region, can be increased to max 255.

## 25
Outstanding VPC peering connection requests

## 50
Customer gateways per region, can be increased upon request

VPN connections per region, can be increased upon request

VPN connections per virtual private gateway, can be increased upon request

Entries per route table.

Active VPC peering connections per VPC, can be increased to 125 but network performance may be impacted.

## 100
BGP Advertised Routes per VPN Connection.

## 200
Subnets per VPC, can be increased upon request

Network ACLs per VPC

Route tables per VPC

# [variable]
Number of network interfaces per instance changes by instance type.

<sub>
Reference: [AWS/Documentation/VPC/Limits](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Appendix_Limits.html)
</sub>