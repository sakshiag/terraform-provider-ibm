package ibm

import (
	"fmt"
	"testing"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
)

func TestAccOpenWhiskRule_Basic(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	triggerName := fmt.Sprintf("terraformtrigger_%d", acctest.RandInt())
	actionName := fmt.Sprintf("terraformaction_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())
	updatedTriggerName := fmt.Sprintf("terraformtrigger_%d", acctest.RandInt())
	updatedActionName := fmt.Sprintf("terraformaction_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleCreate(actionName, triggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskRuleExists("ibm_openwhisk_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleNameUpdate(actionName, triggerName, updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleWithTriggerUpdate(actionName, updatedTriggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskRuleExists("ibm_openwhisk_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleWithActionUpdate(updatedActionName, updatedTriggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskRuleExists("ibm_openwhisk_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "false"),
				),
			},
		},
	})
}

func TestAccOpenWhiskRule_Publish(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	triggerName := fmt.Sprintf("terraformtrigger_%d", acctest.RandInt())
	actionName := fmt.Sprintf("terraformaction_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleDefaultPublish(actionName, triggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskRuleExists("ibm_openwhisk_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleUpdatePublish(actionName, triggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "true"),
				),
			},
		},
	})
}

func TestAccOpenWhiskRule_Import(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	triggerName := fmt.Sprintf("terraformtrigger_%d", acctest.RandInt())
	actionName := fmt.Sprintf("terraformaction_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskRuleImport(actionName, triggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskRuleExists("ibm_openwhisk_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_rule.rule", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_openwhisk_rule.rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckOpenWhiskRuleExists(n string, obj *whisk.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client, err := testAccProvider.Meta().(ClientSession).OpenWhiskClient()
		if err != nil {
			return err
		}
		name := rs.Primary.ID

		rule, _, err := client.Rules.Get(name)
		if err != nil {
			return err
		}

		*obj = *rule
		return nil
	}
}

func testAccCheckOpenWhiskRuleDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "openwhisk_rule" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Rules.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for OpenWhisk Rule (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckOpenWhiskRuleCreate(actionName, triggerName, name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "action" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "trigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.trigger.name}"
			action_name = "${ibm_openwhisk_action.action.name}"
		  }`, actionName, triggerName, name)

}

func testAccCheckOpenWhiskRuleNameUpdate(actionName, triggerName, updatedName string) string {
	return fmt.Sprintf(`
	
		resource "ibm_openwhisk_action" "action" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "trigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.trigger.name}"
			action_name = "${ibm_openwhisk_action.action.name}"
		  }`, actionName, triggerName, updatedName)
}

func testAccCheckOpenWhiskRuleWithTriggerUpdate(actionName, updatedTriggerName, name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "action" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "UpdatedTrigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.UpdatedTrigger.name}"
			action_name = "${ibm_openwhisk_action.action.name}"
		  }`, actionName, updatedTriggerName, name)

}

func testAccCheckOpenWhiskRuleWithActionUpdate(updatedActionName, updatedTriggerName, name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "UpdatedAction" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "UpdatedTrigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.trigger.name}"
			action_name = "${ibm_openwhisk_action.UpdatedAction.name}"
		  }`, updatedActionName, updatedTriggerName, name)

}

func testAccCheckOpenWhiskRuleImport(actionName, triggerName, name string) string {
	return fmt.Sprintf(`
	
		resource "ibm_openwhisk_action" "action" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "trigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.trigger.name}"
			action_name = "${ibm_openwhisk_action.action.name}"
		  }`, actionName, triggerName, name)

}

func testAccCheckOpenWhiskRuleDefaultPublish(actionName, triggerName, name string) string {
	return fmt.Sprintf(`
	
		resource "ibm_openwhisk_action" "action" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "trigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.trigger.name}"
			action_name = "${ibm_openwhisk_action.action.name}"
		  }`, actionName, triggerName, name)

}

func testAccCheckOpenWhiskRuleUpdatePublish(actionName, triggerName, name string) string {
	return fmt.Sprintf(`
	
		resource "ibm_openwhisk_action" "action" {
			name = "%s"
		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/wsk-create-cat.js")}"
			}
		  
			limits = {
			  timeout = 600
			  memory  = 128
			}
		  }
		  
		  resource "ibm_openwhisk_trigger" "trigger" {
			name = "%s"
		  }
		  
		  resource "ibm_openwhisk_rule" "rule" {
			name = "%s"
			trigger_name = "${ibm_openwhisk_trigger.trigger.name}"
			action_name = "${ibm_openwhisk_action.action.name}"
			publish = true
		  }`, actionName, triggerName, name)
}
