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

func TestAccOpenWhiskAction_Basic(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "action.#", "1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "exec.1865481321.kind", "nodejs:6"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "action.455559736.name", updatedName),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_With_Annotations(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionCreateWithAnnotations(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "action.0.name", name),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_With_Sequence(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionCreateWithSequence(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "action.0.name", name),
				),
			},
		},
	})
}

func testAccCheckOpenWhiskActionExists(n string, obj *whisk.Action) resource.TestCheckFunc {

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

		action, _, err := client.Actions.Get(name)
		if err != nil {
			return err
		}

		*obj = *action
		return nil
	}
}

func testAccCheckOpenWhiskActionDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "openwhisk_action" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Actions.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for OpenWhisk Action (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckOpenWhiskActionCreate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_action" "action" {
	action = [
		{
		  "name"      = "%s"
		},
	  ]
	exec = {
	 kind = "nodejs:6"
     code = "${file("test-fixtures/wsk-create-cat.js")}"
  	}
	limits = {
    	timeout = 600
        memory = 128
 	}
}`, name)

}

func testAccCheckOpenWhiskActionUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_action" "action" {
	action = [
		{
		  "name"      = "%s"
		},
	  ]
	exec = {
	 kind = "nodejs:6"
     code = "${file("test-fixtures/wsk-create-cat.js")}"
  	}
	limits = {
    	timeout = 600
        memory = 128
 	}
}`, updatedName)
}

func testAccCheckOpenWhiskActionCreateWithAnnotations(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_action" "action" {
	action = [
		{
		  "name"      = "%s"
		},
	  ]
	exec = {
	 kind = "nodejs:6"
     code = "${file("test-fixtures/wsk-create-cat.js")}"
  	}
	limits = {
    	timeout = 600
        memory = 128
 	}
	annotations = <<EOF
	[
    {
        "key":"web-export",
        "value":true
    },
    {
        "key":"raw-http",
        "value":false
    },
    {
        "key":"final",
        "value":true
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskActionCreateWithSequence(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_action" "action" {
	action = [
		{
		  "name"      = "%s"
		},
	  ]
	exec = {
	 components = ["/whisk.system/utils/split", "/whisk.system/utils/sort"]
	}

}`, name)

}
