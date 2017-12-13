package ibm

import (
	"bytes"
	"fmt"
	"log"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMOpenWhiskRule() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMOpenWhiskRuleCreate,
		Read:     resourceIBMOpenWhiskRuleRead,
		Update:   resourceIBMOpenWhiskRuleUpdate,
		Delete:   resourceIBMOpenWhiskRuleDelete,
		Exists:   resourceIBMOpenWhiskRuleExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of rule",
			},
			"trigger": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The namespace of the trigger",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the trigger",
						},
					},
				},
				Set: resourceIBMOpenWhiskRuleHash,
			},
			"action": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The namespace and the package of the action",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the action",
						},
					},
				},
				Set: resourceIBMOpenWhiskRuleHash,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the rule",
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
		},
	}
}

func resourceIBMOpenWhiskRuleCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	trigger := d.Get("trigger").(*schema.Set).List()[0]
	action := d.Get("action").(*schema.Set).List()[0]

	payload := whisk.Rule{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
		Trigger:   trigger,
		Action:    action,
	}

	if publish, ok := d.GetOk("publish"); ok {
		p := publish.(bool)
		payload.Publish = &p
	}

	if version, ok := d.GetOk("version"); ok {
		payload.Version = version.(string)
	}

	log.Println("[INFO] Creating OpenWhisk rule")
	result, _, err := ruleService.Insert(&payload, true)
	if err != nil {
		return fmt.Errorf("Error creating OpenWhisk rule: %s", err)
	}

	d.SetId(result.Name)

	return resourceIBMOpenWhiskRuleRead(d, meta)
}

func resourceIBMOpenWhiskRuleRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules
	id := d.Id()

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	rule, _, err := ruleService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving OpenWhisk package %s : %s", id, err)
	}

	d.Set("name", rule.Name)
	d.Set("publish", rule.Publish)
	d.Set("version", rule.Version)
	d.Set("status", rule.Status)
	d.Set("trigger", flattenActionOrTrigger(rule.Trigger))
	d.Set("action", flattenActionOrTrigger(rule.Action))
	return nil
}

func resourceIBMOpenWhiskRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()

	payload := whisk.Rule{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	ischanged := false
	if d.HasChange("publish") {
		p := d.Get("publish").(bool)
		payload.Publish = &p
		ischanged = true
	}

	if d.HasChange("trigger") {
		trigger := d.Get("trigger").(*schema.Set).List()[0]
		payload.Trigger = trigger
		ischanged = true
	}

	if d.HasChange("action") {
		action := d.Get("action").(*schema.Set).List()[0]
		payload.Action = action
		ischanged = true
	}

	if ischanged {
		log.Println("[INFO] Update OpenWhisk Rule")
		result, _, err := ruleService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating OpenWhisk Rule: %s", err)
		}
		_, _, err = ruleService.SetState(result.Name, "active")
		if err != nil {
			return fmt.Errorf("Error updating OpenWhisk Rule: %s", err)
		}
	}

	return resourceIBMOpenWhiskRuleRead(d, meta)
}

func resourceIBMOpenWhiskRuleDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules
	id := d.Id()
	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()
	_, err = ruleService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting OpenWhisk Rule: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMOpenWhiskRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).OpenWhiskClient()
	if err != nil {
		return false, err
	}
	ruleService := wskClient.Rules
	id := d.Id()

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(id); err != nil {
		return false, NewQualifiedNameError(id, err)
	}

	wskClient.Namespace = qualifiedName.GetNamespace()
	rule, resp, err := ruleService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with OpenWhisk Client : %s", err)
	}
	return rule.Name == id, nil
}

func resourceIBMOpenWhiskRuleHash(v interface{}) int {
	var buf bytes.Buffer
	qualifiedName := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", qualifiedName["path"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", qualifiedName["name"].(string)))
	return hashcode.String(buf.String())
}
