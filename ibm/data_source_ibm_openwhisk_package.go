package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMOpenWhiskPackage() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMOpenWhiskPackageRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of package",
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to publish the item or not.",
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
			"bind_package_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of binded package",
			},
		},
	}
}

func dataSourceIBMOpenWhiskPackageRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages
	name := d.Get("name").(string)

	pkg, _, err := packageService.Get(name)
	if err != nil {
		return fmt.Errorf("Error retrieving OpenWhisk package %s : %s", name, err)
	}

	d.SetId(pkg.Name)
	d.Set("name", pkg.Name)
	d.Set("publish", pkg.Publish)
	d.Set("version", pkg.Version)
	annotations, err := flattenAnnotations(pkg.Annotations)
	if err != nil {
		return err
	}
	d.Set("annotations", annotations)
	parameters, err := flattenParameters(pkg.Parameters)
	if err != nil {
		return err
	}
	d.Set("parameters", parameters)

	if !isEmpty(*pkg.Binding) {
		d.Set("bind_package_name", fmt.Sprintf("/%s/%s", pkg.Binding.Namespace, pkg.Binding.Name))
	}
	return nil
}
