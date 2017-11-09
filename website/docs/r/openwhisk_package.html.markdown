---
layout: "ibm"
page_title: "IBM : OpenWhisk_package"
sidebar_current: "docs-ibm-resource-openwhisk-package"
description: |-
  Manages IBM Cloud Open Whisk Package.
---

# ibm\_openwhisk_package

Create, update, or delete package for IBM Cloud Open Whisk.

## Example Usage

```hcl
resource "ibm_openwhisk_package" "package" {
  name    = "utils"
  publish = "true"
  parameters = <<EOF
        [
    {
        "key":"name",
        "value":"utils"
    }
]
EOF
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of package.
* `publish` - (Optional, boolean) Whether to publish the item or not. Default false.
* `annotations` - (Optional, string) Annotations on the item.
* `parameters` - (Optional, string) Parameter bindings included in the context passed to the package.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new package.
* `version` - Semantic version of the item.
