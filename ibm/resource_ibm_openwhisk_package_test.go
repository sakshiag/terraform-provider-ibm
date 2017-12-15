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

func TestAccOpenWhiskPackage_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageNameUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageWithAnnotations(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageWithAnnotationsUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageWithParameters(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageWithParametersUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageUpdatePublish(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "true"),
				),
			},
		},
	})
}

func TestAccOpenWhiskPackage_Bind_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())
	bindName := "/whisk.system/alarms"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageBindCreate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "bind_package_name", bindName),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageNameBindUpdate(updatedName, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageBindWithAnnotations(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageBindWithAnnotationsUpdate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageBindWithParameters(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageBindWithParametersUpdate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageBindUpdatePublish(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.bindpackage", "publish", "true"),
				),
			},
		},
	})
}

func TestAccOpenWhiskPackage_Import(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_openwhisk_package.package",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckOpenWhiskPackageExists(n string, obj *whisk.Package) resource.TestCheckFunc {

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

		pkg, _, err := client.Packages.Get(name)
		if err != nil {
			return err
		}

		*obj = *pkg
		return nil
	}
}

func testAccCheckOpenWhiskPackageDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).OpenWhiskClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "openwhisk_package" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Packages.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for OpenWhisk Package (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckOpenWhiskPackageCreate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
	   name = "%s"
}`, name)

}

func testAccCheckOpenWhiskPackageNameUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
	   name = "%s"
}`, updatedName)
}

func testAccCheckOpenWhiskPackageWithAnnotations(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
   	name = "%s"
	user_defined_annotations = <<EOF
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

func testAccCheckOpenWhiskPackageWithAnnotationsUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
   	name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskPackageWithParameters(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
   	name = "%s"
	user_defined_parameters = <<EOF
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
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskPackageWithParametersUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
   	name = "%s"
	user_defined_parameters = <<EOF
	[
    {
        "key":"name",
        "value":"utils"
    }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckOpenWhiskPackageImport(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
   	name = "%s"
}`, name)

}

func testAccCheckOpenWhiskPackageUpdatePublish(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
	   name = "%s"
	   publish = true
	   user_defined_parameters = <<EOF
	[
    {
        "key":"name",
        "value":"utils"
    }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF
}`, name)
}

func testAccCheckOpenWhiskPackageBindCreate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
}`, name, bind)

}

func testAccCheckOpenWhiskPackageNameBindUpdate(updatedName, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
}`, updatedName, bind)
}

func testAccCheckOpenWhiskPackageBindWithAnnotations(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"binded alaram package"
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

}`, name, bind)

}
func testAccCheckOpenWhiskPackageBindWithAnnotationsUpdate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	name = "%s"
	bind_package_name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"binded alaram package"
    }
]
EOF

}`, name, bind)

}

func testAccCheckOpenWhiskPackageBindWithParameters(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	name = "%s"
	bind_package_name = "%s"
	user_defined_parameters = <<EOF
        [
    {
        "key":"cron",
        "value":"0 0 1 0 *"
    },
    {
        "key":"trigger_payload ",
        "value":"{'message':'bye old Year!'}"
    },
    {
        "key":"maxTriggers",
        "value":1
    },
    {
        "key":"userdefined",
        "value":"test"
    }
]
EOF

	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name, bind)

}

func testAccCheckOpenWhiskPackageBindWithParametersUpdate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	name = "%s"
	bind_package_name = "%s"
	user_defined_parameters = <<EOF
	   [
   {
	   "key":"cron",
	   "value":"0 0 1 0 *"
   }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name, bind)

}

func testAccCheckOpenWhiskPackageBindUpdatePublish(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
	   publish = true
	   user_defined_parameters = <<EOF
	   [
   {
	   "key":"cron",
	   "value":"0 0 1 0 *"
   }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF
}`, name, bind)
}
