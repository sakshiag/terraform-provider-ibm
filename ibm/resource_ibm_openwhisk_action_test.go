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

func TestAccOpenWhiskAction_NodeJS(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionNodeJS(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.nodehello", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.nodehello", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.nodehello", "exec.1613817926.kind", "nodejs:6"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_NodeJSWithParams(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionNodeJSWithParams(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.nodehellowithparameter", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.nodehellowithparameter", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.nodehellowithparameter", "exec.2223737754.kind", "nodejs:6"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_NodeJSZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionNodeJSZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.nodezip", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.nodezip", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.nodezip", "exec.915916944.kind", "nodejs:6"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_Python(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionPython(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.pythonhello", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.pythonhello", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.pythonhello", "exec.2177277460.kind", "python"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_PythonZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionPythonZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.pythonzip", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.pythonzip", "exec.799663035.kind", "python"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_PHP(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionPHP(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.phphello", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.phphello", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.phphello", "exec.2196642062.kind", "php:7.1"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_PHPZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionPHPZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.phpzip", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.phpzip", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.phpzip", "exec.36392495.kind", "php:7.1"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_Swift(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionSwift(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.swifthello", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.swifthello", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.swifthello", "exec.455016215.kind", "swift:3.1.1"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_Sequence(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionSequence(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.sequence", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.sequence", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.sequence", "exec.3443818261.kind", "sequence"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_Basic(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.action", "publish", "true"),
				),
			},
		},
	})
}

func TestAccOpenWhiskAction_Import(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenWhiskActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckOpenWhiskActionImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenWhiskActionExists("ibm_openwhisk_action.import", &conf),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.import", "name", name),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_openwhisk_action.import", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_openwhisk_action.import",
				ImportState:       true,
				ImportStateVerify: true,
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

func testAccCheckOpenWhiskActionNodeJS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "nodehello" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionNodeJSWithParams(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "nodehellowithparameter" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonodewithparameter.js")}"
			}
			user_defined_parameters = <<EOF
			 [
				 {
					"key":"place",
					 "value":"India"
				}
			]
		EOF
			}
`, name)

}

func testAccCheckOpenWhiskActionNodeJSZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "nodezip" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${base64encode("${file("test-fixtures/nodeaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionPython(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "pythonhello" {
			name = "%s"		  
			exec = {
			  kind = "python"
			  code = "${file("test-fixtures/hellopython.py")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionPythonZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "pythonzip" {
			name = "%s"		  
			exec = {
			  kind = "python"
			  code = "${base64encode("${file("test-fixtures/pythonaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionPHP(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "phphello" {
			name = "%s"		  
			exec = {
			  kind = "php:7.1"
			  code = "${file("test-fixtures/hellophp.php")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionPHPZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "phpzip" {
			name = "%s"		  
			exec = {
			  kind = "php:7.1"
			  code = "${base64encode("${file("test-fixtures/phpaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionSwift(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "swifthello" {
			name = "%s"		  
			exec = {
			  kind = "swift:3.1.1"
			  code = "${file("test-fixtures/helloSwift.swift")}"
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionSequence(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "sequence" {
			name = "%s"		  
			exec = {
			  kind = "sequence"
			  components = ["/whisk.system/utils/split","/whisk.system/utils/sort"]
			}
		  }
	
`, name)

}

func testAccCheckOpenWhiskActionCreate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "action" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
			}
`, name)

}

func testAccCheckOpenWhiskActionUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "action" {
			name = "%s"	
			publish = "true"	  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonodewithparameter.js")}"
			}
			user_defined_parameters = <<EOF
			[
				{
				   "key":"place",
					"value":"India"
			   }
		   ]
	   EOF
	   user_defined_annotations = <<EOF
	   [
		   {
			  "key":"Description",
			   "value":"Sample code to display hello"
		  }
	  ]
  EOF
			}
`, name)

}

func testAccCheckOpenWhiskActionImport(name string) string {
	return fmt.Sprintf(`
		resource "ibm_openwhisk_action" "import" {
			name = "%s"	
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonodewithparameter.js")}"
			}
			user_defined_parameters = <<EOF
			[
				{
				   "key":"place",
					"value":"India"
			   }
		   ]
	   EOF
	   user_defined_annotations = <<EOF
	   [
		   {
			  "key":"Description",
			   "value":"Sample code to display hello"
		  }
	  ]
  EOF
			}
`, name)

}
