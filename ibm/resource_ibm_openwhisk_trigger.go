package ibm

import (
	"fmt"
	"log"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
)

const FEED_LIFECYCLE_EVENT = "lifecycleEvent"
const FEED_TRIGGER_NAME = "triggerName"
const FEED_AUTH_KEY = "authKey"
const FEED_CREATE = "CREATE"
const FEED_DELETE = "DELETE"

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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of Trigger",
			},
			"feed": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Trigger feed ACTION_NAME",
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
				Description:      "Parameter bindings included in the context passed to the Trigger.",
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

func resourceIBMOpenWhiskTriggerCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	TriggerService := wskClient.Triggers
	feedval := false
	feedPayload := map[string]interface{}{}
	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	payload := whisk.Trigger{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	if v, ok := d.GetOk("annotations"); ok {
		var err error
		payload.Annotations, err = expandAnnotations(v.(string))
		if err != nil {
			return err
		}
	}
	if v, ok := d.GetOk("feed"); ok {
		feedval = true
		feedPaylod := whisk.KeyValue{
			Key:   "feed",
			Value: v.(string),
		}
		feedArray := make([]whisk.KeyValue, 0, 1)
		feedArray = append(feedArray, feedPaylod)
		payload.Annotations = payload.Annotations.AppendKeyValueArr(feedArray)
	}

	if v, ok := d.GetOk("parameters"); ok {
		if feedval {
			var err error
			feedParameters, err := expandParameters(v.(string))
			if err != nil {
				return err
			}
			for _, value := range feedParameters {
				feedPayload[value.Key] = value.Value
			}
		} else {
			var err error
			payload.Parameters, err = expandParameters(v.(string))
			if err != nil {
				return err
			}
		}
	}

	if publish, ok := d.GetOk("publish"); ok {
		p := publish.(bool)
		payload.Publish = &p
	}

	if version, ok := d.GetOk("version"); ok {
		payload.Version = version.(string)
	}

	log.Println("[INFO] Creating OpenWhisk trigger")
	result, _, err := TriggerService.Insert(&payload, true)
	if err != nil {
		return fmt.Errorf("Error creating OpenWhisk trigger: %s", err)
	}

	if _, ok := d.GetOk("feed"); ok {
		actionName := d.Get("feed").(string)
		var feedQualifiedName = new(QualifiedName)

		if feedQualifiedName, err = NewQualifiedName(actionName); err != nil {
			_, _, delerr := TriggerService.Delete(name)
			if delerr != nil {
				return fmt.Errorf("Error creating OpenWhisk trigger with feed: %s", err)
			}
			return NewQualifiedNameError(actionName, err)
		}

		feedPayload[FEED_LIFECYCLE_EVENT] = FEED_CREATE
		feedPayload[FEED_AUTH_KEY] = wskClient.Config.AuthToken
		feedPayload[FEED_TRIGGER_NAME] = fmt.Sprintf("/%s/%s", qualifiedName.GetNamespace(), name)

		wskClient.Namespace = feedQualifiedName.GetNamespace()
		ActionService := wskClient.Actions
		log.Println("******", feedPayload)
		_, _, err := ActionService.Invoke(feedQualifiedName.GetEntityName(), feedPayload, true, false)
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

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

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
	return nil
}

func resourceIBMOpenWhiskTriggerUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	TriggerService := wskClient.Triggers

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()
	payload := whisk.Trigger{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
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

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

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
			FEED_LIFECYCLE_EVENT: FEED_DELETE,
			FEED_AUTH_KEY:        wskClient.Config.AuthToken,
			FEED_TRIGGER_NAME:    fmt.Sprintf("/%s/%s", qualifiedName.GetNamespace(), id),
		}
		wskClient.Namespace = feedQualifiedName.GetNamespace()
		ActionService := wskClient.Actions
		_, _, err := ActionService.Invoke(feedQualifiedName.GetEntityName(), feedPayload, true, false)
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
