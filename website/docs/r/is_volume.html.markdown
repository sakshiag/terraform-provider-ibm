---
layout: "ibm"
page_title: "IBM : volume"
sidebar_current: "docs-ibm-resource-is-volume"
description: |-
  Manages IBM Volume.
---

# ibm\_is_volume

Provides a volume resource. This allows volume to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a volume:

```hcl
resource "ibm_is_volume" "testacc_volume" {
    name 		= "test_volume"
    type 		= "boot"
    zone 		= "test"
    iops 		= 10000
    capacity    = 100
    auto_delete = true
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The user-defined name for this volume.
* `type` - (Required, string) The type of volume attachment.
* `zone` - (Required, string)  The zone this volume resides in.
* `iops` - (Required, int) The bandwidth for the volume.
* `capacity` - (Optional, int) The capacity of the volume in gigabytes. This defaults to `100`.
* `profile` - (Optional, string) The profile this volume uses.
* `source_snapshot` - (Optional, string) The snapshot from which this volume was cloned.
* `encryption_key` - (Optional, string) The encryption key for the volume.
* `auto_delete` - (Optional, bool) If set to true, this volume will be automatically deleted if the only instance it is attached to is deleted. The default is `false`.
* `billing_term` - (Optional, string) The billing term.
* `resource_group` - (Optional, string) The resource group for this volume.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the volume.
* `status` - The status of volume.