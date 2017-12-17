package ibm

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

const feedLifeCycleEvent = "lifecycleEvent"
const feedTriggerName = "triggerName"
const feedAuthKey = "authKey"
const feedCreate = "CREATE"
const feedDelete = "DELETE"

func resourceIBMOpenWhiskTrigger() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMOpenWhiskTriggerCreate,
		Read:     resourceIBMOpenWhiskTriggerRead,
		Update:   resourceIBMOpenWhiskTriggerUpdate,
		Delete:   resourceIBMOpenWhiskTriggerDelete,
		Exists:   resourceIBMOpenWhiskTriggerExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of Trigger",
				ValidateFunc: validateOpenwhiskName,
			},
			"feed": {
				Type:        schema.TypeSet,
				ForceNew:    true,
				Optional:    true,
				MaxItems:    1,
				Description: "Trigger feed",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Trigger feed ACTION_NAME",
						},
						"parameters": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "[]",
							Description:  "Parameters for feed action invoke",
							ValidateFunc: validateJSONString,
							DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
								if o == "" && n == "" {
									return false
								}
								if o == "[]" {
									return true
								}
								return false
							},
							StateFunc: func(v interface{}) string {
								json, _ := normalizeJSONString(v)
								return json
							},
						},
					},
				},
				Set: resourceIBMOpenWhiskTriggerFeedHash,
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
			"user_defined_annotations": {
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
			"user_defined_parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "[]",
				Description:      "Parameter bindings included in the context passed to the Trigger.",
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
		},
	}
}

func resourceIBMOpenWhiskTriggerCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	TriggerService := wskClient.Triggers
	feed := false
	feedPayload := map[string]interface{}{}
	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	payload := whisk.Trigger{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}

	userDefinedAnnotations := d.Get("user_defined_annotations").(string)
	payload.Annotations, err = expandAnnotations(userDefinedAnnotations)
	if err != nil {
		return err
	}

	userDefinedParameters := d.Get("user_defined_parameters").(string)
	payload.Parameters, err = expandParameters(userDefinedParameters)
	if err != nil {
		return err
	}

	if v, ok := d.GetOk("feed"); ok {
		feed = true
		value := v.(*schema.Set).List()[0].(map[string]interface{})
		feedPaylod := whisk.KeyValue{
			Key:   "feed",
			Value: value["name"],
		}
		feedArray := make([]whisk.KeyValue, 0, 1)
		feedArray = append(feedArray, feedPaylod)
		payload.Annotations = payload.Annotations.AppendKeyValueArr(feedArray)
	}

	log.Println("[INFO] Creating OpenWhisk trigger")
	result, _, err := TriggerService.Insert(&payload, false)
	if err != nil {
		return fmt.Errorf("Error creating OpenWhisk trigger: %s", err)
	}

	if feed {
		feed := d.Get("feed").(*schema.Set).List()[0].(map[string]interface{})
		actionName := feed["name"].(string)
		parameters := feed["parameters"].(string)
		var err error
		feedParameters, err := expandParameters(parameters)
		if err != nil {
			return err
		}
		for _, value := range feedParameters {
			feedPayload[value.Key] = value.Value
		}
		var feedQualifiedName = new(QualifiedName)

		if feedQualifiedName, err = NewQualifiedName(actionName); err != nil {
			_, _, delerr := TriggerService.Delete(name)
			if delerr != nil {
				return fmt.Errorf("Error creating OpenWhisk trigger with feed: %s", err)
			}
			return NewQualifiedNameError(actionName, err)
		}

		feedPayload[feedLifeCycleEvent] = feedCreate
		feedPayload[feedAuthKey] = wskClient.Config.AuthToken
		feedPayload[feedTriggerName] = fmt.Sprintf("/%s/%s", qualifiedName.GetNamespace(), name)

		c, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
			Namespace: feedQualifiedName.GetNamespace(),
			AuthToken: wskClient.AuthToken,
			Host:      wskClient.Host,
		})
		ActionService := c.Actions
		_, _, err = ActionService.Invoke(feedQualifiedName.GetEntityName(), feedPayload, true, false)
		if err != nil {
			_, _, delerr := TriggerService.Delete(name)
			if delerr != nil {
				return fmt.Errorf("Error creating OpenWhisk trigger with feed: %s", err)
			}
			return fmt.Errorf("Error creating OpenWhisk trigger with feed: %s", err)
		}
	}

	d.SetId(result.Name)

	return resourceIBMOpenWhiskTriggerRead(d, meta)
}

func resourceIBMOpenWhiskTriggerRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	TriggerService := wskClient.Triggers
	id := d.Id()

	trigger, _, err := TriggerService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving OpenWhisk Trigger %s : %s", id, err)
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
	d.Set("user_defined_parameters", parameters)

	userDefinedAnnotations, err := filterTriggerAnnotations(trigger.Annotations)
	if err != nil {
		return err
	}
	d.Set("user_defined_annotations", userDefinedAnnotations)

	found := trigger.Annotations.FindKeyValue("feed")

	if found >= 0 {
		d.Set("feed", flattenFeed(trigger.Annotations.GetValue("feed").(string)))
	}

	return nil
}

func resourceIBMOpenWhiskTriggerUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	TriggerService := wskClient.Triggers

	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	payload := whisk.Trigger{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	ischanged := false

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

	if ischanged {
		log.Println("[INFO] Update OpenWhisk Trigger")

		_, _, err = TriggerService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating OpenWhisk Trigger: %s", err)
		}
	}

	return resourceIBMOpenWhiskTriggerRead(d, meta)
}

func resourceIBMOpenWhiskTriggerDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	TriggerService := wskClient.Triggers
	id := d.Id()

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}
	trigger, _, err := TriggerService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving OpenWhisk Trigger %s : %s", id, err)
	}
	found := trigger.Annotations.FindKeyValue("feed")
	if found >= 0 {
		actionName := trigger.Annotations.GetValue("feed").(string)
		var feedQualifiedName = new(QualifiedName)

		if feedQualifiedName, err = NewQualifiedName(actionName); err != nil {
			return NewQualifiedNameError(actionName, err)
		}

		feedPayload := map[string]interface{}{
			feedLifeCycleEvent: feedDelete,
			feedAuthKey:        wskClient.Config.AuthToken,
			feedTriggerName:    fmt.Sprintf("/%s/%s", qualifiedName.GetNamespace(), id),
		}
		c, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
			Namespace: feedQualifiedName.GetNamespace(),
			AuthToken: wskClient.AuthToken,
			Host:      wskClient.Host,
		})
		ActionService := c.Actions
		_, _, err = ActionService.Invoke(feedQualifiedName.GetEntityName(), feedPayload, true, false)
		if err != nil {
			return fmt.Errorf("Error deleting OpenWhisk trigger with feed: %s", err)
		}
		wskClient.Namespace = qualifiedName.GetNamespace()
	}

	_, _, err = TriggerService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting OpenWhisk Trigger: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMOpenWhiskTriggerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return false, err
	}
	TriggerService := wskClient.Triggers
	id := d.Id()
	trigger, resp, err := TriggerService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with OpenWhisk Client : %s", err)
	}
	return trigger.Name == id, nil
}

func resourceIBMOpenWhiskTriggerFeedHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-",
		m["name"].(string)))

	return hashcode.String(buf.String())
}
