# Building a Dynamic Security Response Environment (Strangeloop 2015 Workshop)

This repository contains the slides and requisite material for the
*Building a Dynamic Security Response Environment* workshop. The
workshop will focus on setting up and building out a fully functional
web application defense system.

In order to make the most of our time together you should come
prepared. This means installing the appropriate things on your system
ahead of time. The following list of dependencies are required to
complete the exercises. If you have trouble installing any of them
there will be some time at the beginning of the workshop to sort
things out, but not long.

## An important note

There is no support for Windows in this workshop. The examples will
only work on Linux or Unix systems. If you have a windows laptop,
please install Linux inside a virtual machine or come prepared to work
with another person during the workshop.

## Pre-Workshop Setup

You will need to install the following dependencies.

* redis
* gcc
* autotools
* make
* pkgconfig
* zlib
* pcre
* check
* ruby and rubygems
* go

With the multitude of package managers out there it may not be obvious
how to best install these dependencies. This is an attempt to cover
the most common cases. If you don't see your specific setup and have
questions, please file an issue and I will try to accommodate.

There are many options for installing ruby, but the most common are
[rvm](https://rvm.io/) and
[rbenv](https://github.com/sstephenson/rbenv). You can use Ruby 1.9 or
higher for this workshop.


#### OS X

You will need to have [homebrew](http://brew.sh/) installed. Please
see their website for install instructions. Make sure you have run
`brew doctor` and resolved any issues and that you have run `brew
update` to get the latest version of any dependencies.

Installing the requisite xcode develper tools will get you most of the
dependencies listed above. Next run the following:

```sh
$ brew install redis check go
```

#### Apt based systems (Debian, Ubuntu, etc)

This particular example was done on Ubuntu 14.04 server. Your
experience may vary but it should be close.

```sh
$ sudo apt-get install build-essential autoconf automake libtool pkgconfig redis-server check libpcre3-dev zlib1g-dev libcurl4-openssl-dev
```

#### Rpm based systems (Redhat, Fedora, Centos, etc)

This particular example was done on Centos 7 minimal. Your experience
may vary but it should be close.

```sh
$ sudo yum groupinstall "Developer Tools"
$ curl -O http://dl.fedoraproject.org/pub/epel/7/x86_64/e/epel-release-7-5.noarch.rpm
$ sudo rpm -ivh epel-release-7-5.noarch.rpm
$ sudo yum update
$ sudo yum install check check-devel redis pcre-devel zlib-devel curl-devel go
$ sudo systemctl start redis
```
