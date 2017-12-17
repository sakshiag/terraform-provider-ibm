---
layout: "ibm"
page_title: "IBM : OpenWhisk_trigger"
sidebar_current: "docs-ibm-resource-openwhisk-trigger"
description: |-
  Manages IBM Cloud Open Whisk trigger.
---

# ibm\_openwhisk_trigger

Create, update, or delete for IBM Cloud Open Whisk trigger.

## Example Usage

```hcl
resource "ibm_openwhisk_trigger" "trigger" {
			name = "trigger-name"		  
			user_defined_parameters = <<EOF
			[
				{
				   "key":"place",
					"value":"India"
			   }
		   ]
	   EOF
	   user_defined_annotations = <<EOF
	   [
		   {
			  "key":"Description",
			   "value":"Sample code to display hello"
		  }
	  ]
  EOF
			}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of trigger.
* `feed` - (Optional, set) A nested block describing the feed assigned to . Nested `feed` blocks have   the following structure:
    * `name` - (Required, string) Trigger feed ACTION_NAME
    * `parameters` - (Optinal, string) Parameters for feed action invoke
* `user_defined_annotations` - (Optional, string) User defined annotations on the item.
* `user_defined_parameters` - (Optional, string) User defined parameter bindings included in the context passed to the action.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new trigger.
* `publish` - Whether to publish the item or not.Default is false
* `version` - Semantic version of the item.
* `annotations` -  Annotations on the item.
* `parameters` - Parameter on the item
