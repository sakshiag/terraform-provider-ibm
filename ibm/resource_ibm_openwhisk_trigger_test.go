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

func TestAccOpenWhiskTrigger_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "parameters", "[]"),
				),
			},
		},
	})
}

func TestAccOpenWhiskTrigger_With_Annotations(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerWithAnnotations(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerWithAnnotationsUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
				),
			},
		},
	})
}

func TestAccOpenWhiskTrigger_With_Parameters(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerWithParameters(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerWithParametersUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
				),
			},
		},
	})
}

func TestAccOpenWhiskTrigger_Publish(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerDefaultPublish(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerUpdatePublish(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "true"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "parameters", "[]"),
				),
			},
		},
	})
}

func TestAccOpenWhiskTrigger_Import(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskTriggerImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskTriggerExists("ibm_openwhisk_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_trigger.trigger", "parameters", "[]"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_openwhisk_trigger.trigger",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckOpenWhiskTriggerExists(n string, obj *whisk.Trigger) resource.TestCheckFunc {

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

		trigger, _, err := client.Triggers.Get(name)
		if err != nil {
			return err
		}

		*obj = *trigger
		return nil
	}
}

func testAccCheckOpenWhiskTriggerDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "openwhisk_Trigger" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Triggers.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for OpenWhisk Trigger (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckOpenWhiskTriggerCreate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
	   name = "%s"
}`, name)

}

func testAccCheckOpenWhiskTriggerUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
	   name = "%s"
}`, updatedName)
}

func testAccCheckOpenWhiskTriggerWithAnnotations(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
   	name = "%s"
	annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    },
    {
        "key":"sampleOutput",
        "value": {
			"count": 3
		}
    },
    {
        "key":"final",
        "value": [
			{
				"description": "A string",
				"name": "payload",
				"required": true
			}
		]
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskTriggerWithAnnotationsUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
   	name = "%s"
	annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskTriggerWithParameters(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
   	name = "%s"
	parameters = <<EOF
	[
    {
        "key":"place",
        "value":"city"
    },
    {
        "key":"parameter",
        "value": {
			"count": 3
		}
    },
    {
        "key":"final",
        "value": [
			{
				"description": "Set of Values",
				"name": "payload",
				"required": true
			}
		]
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskTriggerWithParametersUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
   	name = "%s"
	parameters = <<EOF
	[
    {
        "key":"name",
        "value":"utils"
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskTriggerImport(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
   	name = "%s"
}`, name)

}

func testAccCheckOpenWhiskTriggerDefaultPublish(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
   	name = "%s"
}`, name)

}

func testAccCheckOpenWhiskTriggerUpdatePublish(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_trigger" "trigger" {
	   name = "%s"
	   publish = true
}`, name)
}
