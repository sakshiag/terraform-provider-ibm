---
layout: "ibm"
page_title: "IBM : vpc"
sidebar_current: "docs-ibm-resource-is-vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm\_is_vpc

Provides a vpc resource. This allows VPC to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a VPC:

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
    name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `default_network_acl` - (Optional, string) ID of the default network ACL.
* `is_default` - (Optional, string)  Make this VPC default for the account. The default is false.
* `name` - (Required, string) The name of the VPC.
* `resource_group` - (Optional, string) The resource group where the VPC to be created
* `tags` - (Optional, array of strings) Tags associated with the VPC. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPC.
* `default_security_group` - The unique identifier of the VPC default security group.
* `status` - The status of VPC.