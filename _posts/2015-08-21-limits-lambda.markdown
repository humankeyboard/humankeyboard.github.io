---
layout: bootstrap_post
title: Lambda Limitations
date: 2015-08-21 13:13:13
author: Oz Akan
abstract: Limitations around Lambda service.
categories:
    - AWS
tags:
    - aws, restrictions, lambda, limits
    
---

## 512 MB
Ephemeral disk capacity (“/tmp” space)

## 1024
Number of file descriptors

## 1024
Number of processes + threads

## 100
Concurrent requests per account (requests per second * average duration per invocation). This number can be increased upon request.

## 60 Seconds
Maximum execution duration per request

## 50 MB
Lambda function deployment package size (.zip/.jar file)

## 250 MB
Lambda function deployment package size (.zip/.jar file)

## 6 MB
Invoke request body payload size

## 6 MB
Invoke response body payload size

## 1.5 GB
Total size of all the deployment packages that can be uploaded per account


<sub>
Reference: [AWS/Documentation/Lambda/Limits](http://docs.aws.amazon.com/lambda/latest/dg/limits.html)
</sub>
