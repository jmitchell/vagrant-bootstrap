# Vagrant Bootstrap

Vagrant Bootstrap is very much in progress, and doesn't offer any
value yet.


## The Vision

Lots of people have computers at their disposal, but the average user
isn't familiar with every platform's conventions. Even writing generic
instructions on how to determine a computer's local IP is
challenging. It's reminiscent of the pain when programming OS-specific
code.

```
#ifdef _WIN32
// blah, blah, blah
#endif

#ifdef linux
// this hurts; can I stop?
#endif
```

Vagrant Bootstrap will ease this pain by making it easy for any user
to set up Vagrant and its dependencies on their machine. Whenever the
user wants, she can install a new standardized VM (known as a
*world*). It will probably be a modern version of Linux with good
support for [Linux Containers](https://linuxcontainers.org/) and
[Docker](http://www.docker.com/), such as
[Ubuntu 14.04](http://www.ubuntu.com/download/desktop).

Each *world* serves as a consistent foundation for further
work. By themselves, worlds will be a useful construct for users who
want to learn more about computing via this standardized
platform. But there's room for more.

A *world* can have several *terroritories*. A *terrority* is a
semi-isolated environment such as a Linux Container, and it's hosted
on the *world* VM. There will be standard ways to specify and
implement territory specifications, possibly building on a
configuration management and remote execution framework such as
[Salt Stack](http://www.saltstack.com/). Knowledgeable computer users
will be able to implement specifications for the territories that
should exist in a world and how they may or may not communicate with
each other. Since these territory specifications all target the same
world platform, there's little concern about cross-platform
programming or strange variation in end-user
documentation. Additionally, user-friendly tooling will make it easy
for users to populate their new *worlds* with territories based on a
specification code repository.


## Getting there

The initial hurdle is identifying the latest stable Vagrant code to
download. For months the Vagrant project has diligently updated
[version.txt](https://github.com/mitchellh/vagrant/blob/master/version.txt)
and tagged each release. Additionally, OS-specific code will be
required to install Vagrant and its dependencies, ideally in some
sandbox to prevent collisions with existing or future installations.

One of the most critical dependencies is a
[provider](http://docs.vagrantup.com/v2/providers/index.html). For
this project [VirtualBox](virtualbox.org) will serve as the default
provider as it runs on a wide variety of operating systems and can be
installed free of charge.
