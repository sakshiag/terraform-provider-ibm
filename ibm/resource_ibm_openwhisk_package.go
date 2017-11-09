package ibm

import (
	"fmt"
	"log"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMOpenWhiskPackage() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMOpenWhiskPackageCreate,
		Read:     resourceIBMOpenWhiskPackageRead,
		Update:   resourceIBMOpenWhiskPackageUpdate,
		Delete:   resourceIBMOpenWhiskPackageDelete,
		Exists:   resourceIBMOpenWhiskPackageExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of package",
			},
			"publish": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to publish the item or not.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item",
			},
			"annotations": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Annotations on the item.",
				Default:          "[]",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
			},
			"parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "[]",
				Description:      "Parameter bindings included in the context passed to the package.",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
			},
		},
	}
}

func resourceIBMOpenWhiskPackageCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages

	payload := whisk.Package{
		Name:      d.Get("name").(string),
		Namespace: wskClient.Namespace,
	}
	if v, ok := d.GetOk("annotations"); ok {
		var err error
		payload.Annotations, err = expandAnnotations(v.(string))
		if err != nil {
			return err
		}
	}

	if v, ok := d.GetOk("parameters"); ok {
		var err error
		payload.Parameters, err = expandParameters(v.(string))
		if err != nil {
			return err
		}
	}

	if publish, ok := d.GetOk("publish"); ok {
		p := publish.(bool)
		payload.Publish = &p
	}

	if version, ok := d.GetOk("version"); ok {
		payload.Version = version.(string)
	}

	log.Println("[INFO] Creating OpenWhisk pkg")
	result, _, err := packageService.Insert(&payload, true)
	if err != nil {
		return fmt.Errorf("Error creating OpenWhisk pkg: %s", err)
	}

	d.SetId(result.Name)

	return resourceIBMOpenWhiskPackageRead(d, meta)
}

func resourceIBMOpenWhiskPackageRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages
	id := d.Id()

	pkg, _, err := packageService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving OpenWhisk package %s : %s", id, err)
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
	return nil
}

func resourceIBMOpenWhiskPackageUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages
	payload := whisk.Package{
		Name:      d.Get("name").(string),
		Namespace: wskClient.Namespace,
	}
	ischanged := false
	if d.HasChange("publish") {
		p := d.Get("publish").(bool)
		payload.Publish = &p
		ischanged = true
	}

	if d.HasChange("parameters") {
		var err error
		payload.Parameters, err = expandParameters(d.Get("parameters").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if d.HasChange("annotations") {
		var err error
		payload.Annotations, err = expandAnnotations(d.Get("annotations").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if ischanged {
		log.Println("[INFO] Update OpenWhisk Package")

		_, _, err = packageService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating OpenWhisk Package: %s", err)
		}
	}

	return resourceIBMOpenWhiskPackageRead(d, meta)
}

func resourceIBMOpenWhiskPackageDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages
	id := d.Id()
	_, err = packageService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting OpenWhisk Package: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMOpenWhiskPackageExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return false, err
	}
	packageService := wskClient.Packages
	id := d.Id()
	pkg, resp, err := packageService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with OpenWhisk Client : %s", err)
	}
	return pkg.Name == id, nil
}
