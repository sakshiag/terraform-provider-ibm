package ibm

import (
	"fmt"
<<<<<<< HEAD
	"testing"

	"github.com/IBM-Cloud/bluemix-go/models"

	"strings"

=======
	"strings"
	"testing"
	"time"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisInstance_Basic(t *testing.T) {
<<<<<<< HEAD
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
=======
	var cisInstanceOne string
	testName := "test_acc"
	name := "ibm_cis." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
<<<<<<< HEAD
				Config: testAccCheckIBMCisInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists("ibm_cis.instance", conf),
					resource.TestCheckResourceAttr("ibm_cis.instance", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_cis.instance", "service", "internet-svcs"),
					resource.TestCheckResourceAttr("ibm_cis.instance", "plan", "standard"),
					resource.TestCheckResourceAttr("ibm_cis.instance", "location", "global"),
=======
				Config: testAccCheckIBMCisInstance_basic(cisResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(name, &cisInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "global"),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				),
			},
		},
	})
}

<<<<<<< HEAD
func TestAccIBMCisInstance_import(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_cis.instance"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisInstance_basic(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
=======
func TestAccIBMCisInstance_CreateAfterManualDestroy(t *testing.T) {
	//t.Parallel()
	var cisInstanceOne, cisInstanceTwo string
	testName := "test_acc"
	name := "ibm_cis." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisInstance_basic(cisResourceGroup, testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisInstanceExists(name, &cisInstanceOne),
					testAccCisInstanceManuallyDelete(&cisInstanceOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisInstance_basic(cisResourceGroup, testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisInstanceExists(name, &cisInstanceTwo),
					func(state *terraform.State) error {
						if cisInstanceOne == cisInstanceTwo {
							return fmt.Errorf("Cis instance id is unchanged even after we thought we deleted it ( %s )", cisInstanceTwo)
						}
						return nil
					},
				),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
			},
		},
	})
}

<<<<<<< HEAD
func TestAccIBMCisInstance_with_resource_group(t *testing.T) {
	var conf models.ServiceInstance
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_cis.instance"
=======
func TestAccIBMCisInstance_import(t *testing.T) {
	var cisInstanceOne string
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_cis." + serviceName
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
<<<<<<< HEAD
				Config: testAccCheckIBMCisInstance_with_resource_group(serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(resourceName, conf),
=======
				Config: testAccCheckIBMCisInstance_basic(cisResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCisInstanceExists(resourceName, &cisInstanceOne),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "internet-svcs"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", "global"),
				),
			},
<<<<<<< HEAD
=======
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		},
	})
}

func testAccCheckIBMCisInstanceDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis" {
			continue
		}

		instanceID := rs.Primary.ID

		_, err := rsContClient.ResourceServiceInstance().GetInstance(instanceID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

<<<<<<< HEAD
func testAccCheckIBMCisInstanceExists(n string, obj models.ServiceInstance) resource.TestCheckFunc {
=======
func testAccCisInstanceManuallyDelete(tfCisId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_ = testAccCisInstanceManuallyDeleteUnwrapped(s, tfCisId)
		return nil
	}
}

func testAccCisInstanceManuallyDeleteUnwrapped(s *terraform.State, tfCisId *string) error {
	rsConClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	instance := *tfCisId
	var instanceId string
	// if Id does not start with CRN, then zoneId/Pool/HealthCheckId passed. Extract InstanceId
	if strings.HasPrefix(instance, "crn") {
		instanceId = instance
	} else {
		_, instanceId, _ = convertTftoCisTwoVar(instance)
	}
	err = rsConClient.ResourceServiceInstance().DeleteInstance(instanceId, true)
	if err != nil {
		return fmt.Errorf("Error deleting resource instance: %s", err)
	}

	_ = &resource.StateChangeConf{
		Pending: []string{cisInstanceProgressStatus, cisInstanceInactiveStatus, cisInstanceSuccessStatus},
		Target:  []string{cisInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceId)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, cisInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if instance.State == cisInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed to delete: %v", instanceId, err)
			}
			return instance, instance.State, nil
		},
		Timeout:    90 * time.Second,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	if err != nil {
		return fmt.Errorf(
			"Error waiting for resource instance (%s) to be deleted: %s", instanceId, err)
	}
	return nil
}

func testAccCheckIBMCisInstanceExists(n string, tfCisId *string) resource.TestCheckFunc {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceControllerAPI()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID

		instance, err := rsContClient.ResourceServiceInstance().GetInstance(instanceID)
<<<<<<< HEAD

		if err != nil {
			return err
		}

		obj = instance
=======
		if err != nil {
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*tfCisId = ""
				return nil
			}
			return fmt.Errorf("Error retrieving resource instance: %s", err)
		}
		if strings.Contains(instance.State, "removed") {
			*tfCisId = ""
			return nil
		}

		*tfCisId = instanceID
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		return nil
	}
}

<<<<<<< HEAD
func testAccCheckIBMCisInstance_basic(serviceName string) string {
	return fmt.Sprintf(`
		
		resource "ibm_cis" "instance" {
			name              = "%s"		
			plan              = "standard"
			location          = "global"
			
			timeouts {
				create = "15m"
				update = "15m"
				delete = "15m"
			  }
		}
	`, serviceName)
}

func testAccCheckIBMCisInstance_with_resource_group(serviceName string) string {
	return fmt.Sprintf(`

		data "ibm_resource_group" "group" {
			name = "default"
		}
		
		resource "ibm_cis" "instance" {
			name              = "%s"		
			plan              = "standard"
			location          = "global"
			resource_group_id = "${data.ibm_resource_group.group.id}"
			
		}
	`, serviceName)
=======
func testAccCheckIBMCisInstance_basic(cisResourceGroup string, name string) string {
	return fmt.Sprintf(`
				data "ibm_resource_group" "test_acc" {
				  name = "%[1]s"
				}

				resource "ibm_cis" "%[2]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[2]s"
				  plan              = "standard"
				  location          = "global"
				}`, cisResourceGroup, name)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}
