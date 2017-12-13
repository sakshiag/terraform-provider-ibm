package ibm

import (
	"bytes"
	"fmt"
	"log"
	"reflect"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/hashcode"
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
			"user_defined_annotations": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Annotations on the item.",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
			},
			"user_defined_parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Parameter bindings included in the context passed to the package.",
				ValidateFunc:     validateJSONString,
				DiffSuppressFunc: suppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := normalizeJSONString(v)
					return json
				},
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
			"binding": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Set: resourceIBMOpenWhiskPackageBindingHash,
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

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	payload := whisk.Package{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	if v, ok := d.GetOk("user_defined_annotations"); ok {
		var err error
		payload.Annotations, err = expandAnnotations(v.(string))
		if err != nil {
			return err
		}
	}

	if v, ok := d.GetOk("user_defined_parameters"); ok {
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

	if v, ok := d.GetOk("binding"); ok {
		var err error
		payload.Binding, err = expandBinding(v.(*schema.Set).List())
		if err != nil {
			return err
		}
	}

	log.Println("[INFO] Creating OpenWhisk pkg")
	result, _, err := packageService.Insert(&payload, false)
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

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

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

	if !isEmpty(*pkg.Binding) {
		binding := flattenBinding(pkg.Binding)
		d.Set("binding", binding)
	}
	return nil
}

func resourceIBMOpenWhiskPackageUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	packageService := wskClient.Packages

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	payload := whisk.Package{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	ischanged := false
	if d.HasChange("publish") {
		p := d.Get("publish").(bool)
		payload.Publish = &p
		ischanged = true
	}

	if d.HasChange("user_defined_parameters") {
		var err error
		payload.Parameters, err = expandParameters(d.Get("user_defined_parameters").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if d.HasChange("user_defined_annotations") {
		var err error
		payload.Annotations, err = expandAnnotations(d.Get("user_defined_annotations").(string))
		if err != nil {
			return err
		}
		ischanged = true
	}

	if d.HasChange("binding") {
		var err error
		payload.Binding, err = expandBinding(d.Get("binding").(*schema.Set).List())
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
	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()
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

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return false, NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()
	pkg, resp, err := packageService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with OpenWhisk Client : %s", err)
	}
	return pkg.Name == id, nil
}

func resourceIBMOpenWhiskPackageBindingHash(v interface{}) int {
	var buf bytes.Buffer
	binding := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", binding["namespace"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", binding["name"].(string)))
	return hashcode.String(buf.String())
}

func isEmpty(object interface{}) bool {
	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}
