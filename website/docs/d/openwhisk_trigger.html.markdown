---
layout: "ibm"
page_title: "IBM : OpenWhisk_trigger"
sidebar_current: "docs-ibm-data-source-openwhisk-trigger"
description: |-
  Get information on an IBM Cloud Open Whisk Trigger.
---

# ibm\_openwhisk_trigger

Import the details of an existing triiger for IBM Cloud Open Whisk as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_openwhisk_trigger" "trigger" {
			name = "trigger-name"		  
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of trigger.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new trigger.
* `publish` - Whether to publish the item or not.Default is false
* `version` - Semantic version of the item.
* `annotations` -  Annotations on the item.
* `parameters` - Parameter on the item
