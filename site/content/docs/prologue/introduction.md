---
title: "Introduction"
description: "DLX is a helper tool for LXD to quickly provision opinionated development environments on your own hardware."
lead: "DLX is a helper tool for LXD to quickly provision opinionated development environments on your own hardware."
date: 2020-10-06T08:48:57+00:00
lastmod: 2020-10-06T08:48:57+00:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 100
toc: true
---

## What is DLX?

DLX is a cli tool for LXD to quickly provision opinionated development environments on your own hardware. You provide a LXD server and DLX does the rest, rapidly provisioning new development environments for you.

DLX is especially delightful if you install it on a spare computer or laptop that you aren't currently using.

### Why DLX?

DLX is optimized for rapidly spinning up development environments. While there's nothing inherently development-centric about DLX, it was built with by a developer who switches projects frequently, and also switches between multiple computers frequently. DLX lets your create and destroy development environments quickly, and lets you access them from anywhere using standard development tools like SSH or VS Code Remote Development.

DLX gives you a stable container for development that feels like a full VM. LXD runs a full `init` in each container, so you can install required services like databases but keep them separate from other projects.

If you set up LXD with ZFS storage, you get the additional benefit of copy-on-write storage, so you can have dozens of containers on the host but the base operating system files are only stored once.

### Components

DLX has two components:

* **DLX** - The cli tool that interacts with the server.
* An **LXD** server. This is the server that DLX will use to provision new development environments.

## Get started

There are two main ways to get started with Doks:

### Tutorial

{{< alert icon="👉" text="The Tutorial is intended for novice to intermediate users." />}}

Step-by-step instructions on how to start a new Doks project. [Tutorial →](https://getdoks.org/tutorial/introduction/)

### Quick Start

{{< alert icon="👉" text="The Quick Start is intended for intermediate to advanced users." />}}

One page summary of how to start a new Doks project. [Quick Start →]({{< relref "quick-start" >}})

## Go further

Recipes, Reference Guides, Extensions, and Showcase.

### Recipes

Get instructions on how to accomplish common tasks with Doks. [Recipes →](https://getdoks.org/docs/recipes/project-configuration/)

### Reference Guides

Learn how to customize Doks to fully make it your own. [Reference Guides →](https://getdoks.org/docs/reference-guides/security/)

### Extensions

Get instructions on how to add even more to Doks. [Extensions →](https://getdoks.org/docs/extensions/breadcrumb-navigation/)

### Showcase

See what others have build with Doks. [Showcase →](https://getdoks.org/showcase/electric-blocks/)

## Contributing

Find out how to contribute to Doks. [Contributing →](https://getdoks.org/docs/contributing/how-to-contribute/)

## Help

Get help on Doks. [Help →]({{< relref "how-to-update" >}})
