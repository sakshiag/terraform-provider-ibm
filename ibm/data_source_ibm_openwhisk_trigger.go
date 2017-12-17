package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMOpenWhiskTrigger() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMOpenWhiskTriggerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Trigger",
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The publish status of the item",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item",
			},

			"annotations": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Annotations on the item.",
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Parameters on the item.",
			},
		},
	}
}

func dataSourceIBMOpenWhiskTriggerRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	triggerService := wskClient.Triggers
	name := d.Get("name").(string)

	trigger, _, err := triggerService.Get(name)
	if err != nil {
		return fmt.Errorf("Error retrieving OpenWhisk Trigger %s : %s", name, err)
	}

	d.SetId(trigger.Name)
	d.Set("name", trigger.Name)
	d.Set("publish", trigger.Publish)
	d.Set("version", trigger.Version)
	annotations, err := flattenAnnotations(trigger.Annotations)
	if err != nil {
		return err
	}
	d.Set("annotations", annotations)
	parameters, err := flattenParameters(trigger.Parameters)
	if err != nil {
		return err
	}
	d.Set("parameters", parameters)

	return nil
}
