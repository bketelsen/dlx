---
title: dlx lxc image
description: dlx lxc image
lead: dlx lxc image
date: 2021-10-12T10:37:58Z
lastmod: 2021-10-12T10:37:58Z
draft: false
images: []
menu:
  docs:
    parent: "cli"
weight: 100
toc: true
---
## dlx lxc image

Manage images

### Synopsis

Description:
  Manage images

  In LXD instances are created from images. Those images were themselves
  either generated from an existing instance or downloaded from an image
  server.

  When using remote images, LXD will automatically cache images for you
  and remove them upon expiration.

  The image unique identifier is the hash (sha-256) of its representation
  as a compressed tarball (or for split images, the concatenation of the
  metadata and rootfs tarballs).

  Images can be referenced by their full hash, shortest unique partial
  hash or alias name (if one is set).



```
dlx lxc image [flags]
```

### Options inherited from parent commands

```
      --debug         Show all debug messages
      --force-local   Force using the local unix socket
  -h, --help          Print help
      --project       Override the source project
  -q, --quiet         Don't show progress information
  -v, --verbose       Show all information messages
      --version       Print version number
```

### SEE ALSO

* [dlx lxc](/docs/cmd/dlx_lxc)	 - Command line client for LXD
* [dlx lxc image alias](/docs/cmd/dlx_lxc_image_alias)	 - Manage image aliases
* [dlx lxc image copy](/docs/cmd/dlx_lxc_image_copy)	 - Copy images between servers
* [dlx lxc image delete](/docs/cmd/dlx_lxc_image_delete)	 - Delete images
* [dlx lxc image edit](/docs/cmd/dlx_lxc_image_edit)	 - Edit image properties
* [dlx lxc image export](/docs/cmd/dlx_lxc_image_export)	 - Export and download images
* [dlx lxc image get-property](/docs/cmd/dlx_lxc_image_get-property)	 - Get image properties
* [dlx lxc image import](/docs/cmd/dlx_lxc_image_import)	 - Import images into the image store
* [dlx lxc image info](/docs/cmd/dlx_lxc_image_info)	 - Show useful information about images
* [dlx lxc image list](/docs/cmd/dlx_lxc_image_list)	 - List images
* [dlx lxc image refresh](/docs/cmd/dlx_lxc_image_refresh)	 - Refresh images
* [dlx lxc image set-property](/docs/cmd/dlx_lxc_image_set-property)	 - Set image properties
* [dlx lxc image show](/docs/cmd/dlx_lxc_image_show)	 - Show image properties
* [dlx lxc image unset-property](/docs/cmd/dlx_lxc_image_unset-property)	 - Unset image properties

