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
				Config: testAccCheckOpenWhiskPackageUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "parameters", "[]"),
				),
			},
		},
	})
}

func TestAccOpenWhiskPackage_With_Annotations(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskPackageDestroy,
		Steps: []resource.TestStep{

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
		},
	})
}

func TestAccOpenWhiskPackage_With_Parameters(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageWithParameters(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageWithParametersUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskPackageExists("ibm_openwhisk_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "false"),
				),
			},
		},
	})
}

func TestAccOpenWhiskPackage_Publish(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskPackageDefaultPublish(name),
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
				Config: testAccCheckOpenWhiskPackageUpdatePublish(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "publish", "true"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_openwhisk_package.package", "parameters", "[]"),
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

func testAccCheckOpenWhiskPackageUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
	   name = "%s"
}`, updatedName)
}

func testAccCheckOpenWhiskPackageWithAnnotations(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
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

func testAccCheckOpenWhiskPackageWithAnnotationsUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
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

func testAccCheckOpenWhiskPackageWithParameters(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
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

func testAccCheckOpenWhiskPackageWithParametersUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
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

func testAccCheckOpenWhiskPackageImport(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_openwhisk_package" "package" {
   	name = "%s"
}`, name)

}

func testAccCheckOpenWhiskPackageDefaultPublish(name string) string {
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
}`, name)
}
