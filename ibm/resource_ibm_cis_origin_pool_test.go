package ibm

import (
	"fmt"
<<<<<<< HEAD
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
=======
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	"testing"
)

func TestAccIBMCisPool_Basic(t *testing.T) {
	//t.Parallel()
<<<<<<< HEAD
	var pool v1.Pool
=======
	var pool string
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool." + rnd

	resource.Test(t, resource.TestCase{
<<<<<<< HEAD
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigBasic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigBasic(rnd, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool, cis_domain),
=======
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigCisDS_Basic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDS_Basic(rnd, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCisPool_import(t *testing.T) {
	name := "ibm_cis_origin_pool.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
<<<<<<< HEAD
				Config: testAccCheckCisPoolConfigBasic("test", cis_domain),
=======
				Config: testAccCheckCisPoolConfigCisDS_Basic("test", cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "origins.#", "1"),
				),
			},
			resource.TestStep{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func TestAccIBMCisPool_FullySpecified(t *testing.T) {
	//t.Parallel()
<<<<<<< HEAD
	var pool v1.Pool
=======
	var pool string
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	rnd := acctest.RandString(10)
	name := "ibm_cis_origin_pool." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
<<<<<<< HEAD
				Config: testAccCheckCisPoolConfigFullySpecified(rnd, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool, cis_domain),
=======
				Config: testAccCheckCisPoolConfigFullySpecified(rnd, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &pool),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "description", "tfacc-fully-specified"),
					resource.TestCheckResourceAttr(name, "check_regions.#", "1"),
					resource.TestCheckResourceAttr(name, "minimum_origins", "2"),
					resource.TestCheckResourceAttr(name, "notification_email", "admin@outlook.com"),
					resource.TestCheckResourceAttr(name, "origins.#", "2"),
				),
			},
		},
	})
}

<<<<<<< HEAD
func testAccCheckCisPoolDestroy(s *terraform.State, cis_domain string) error {
=======
func TestAccIBMCisPool_CreateAfterManualDestroy(t *testing.T) {
	//t.Parallel()
	var poolOne, poolTwo string
	testName := "test_acc"
	name := "ibm_cis_origin_pool." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolOne),
					testAccCisPoolManuallyDelete(&poolOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisPoolConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolTwo),
					func(state *terraform.State) error {
						if poolOne == poolTwo {
							return fmt.Errorf("id is unchanged even after we thought we deleted it ( %s )",
								poolTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccIBMCisPool_CreateAfterCisRIManualDestroy(t *testing.T) {
	//t.Parallel()
	var poolOne, poolTwo string
	testName := "test"
	name := "ibm_cis_origin_pool." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisPoolConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolOne),
					testAccCisPoolManuallyDelete(&poolOne),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_domain" {
								log.Printf("[WARN] Removing domain")
								zoneId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisId, zoneId)
								cisPtr := &cisId
								log.Printf("[WARN] Removing Cis Instance")
								_ = testAccCisInstanceManuallyDeleteUnwrapped(state, cisPtr)
							}

						}
						return nil
					},
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisPoolConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisPoolExists(name, &poolTwo),
					func(state *terraform.State) error {
						if poolOne == poolTwo {
							return fmt.Errorf("id is unchanged even after we thought we deleted it ( %s )",
								poolTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCheckCisPoolDestroy(s *terraform.State) error {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_origin_pool" {
			continue
		}
<<<<<<< HEAD

		_, err = cisClient.Pools().GetPool(rs.Primary.Attributes["cis_id"], rs.Primary.ID)
=======
		poolId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		_, err = cisClient.Pools().GetPool(cisId, poolId)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		if err == nil {
			return fmt.Errorf("Load balancer pool still exists")
		}
	}

	return nil
}

<<<<<<< HEAD
func testAccCheckCisPoolExists(n string, pool *v1.Pool, cis_domain string) resource.TestCheckFunc {
=======
func testAccCheckCisPoolExists(n string, tfPoolId *string) resource.TestCheckFunc {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}

<<<<<<< HEAD
		poolId, _, _ := convertTftoCisTwoVar(rs.Primary.ID)
		foundPool, err := cisClient.Pools().GetPool(rs.Primary.Attributes["cis_id"], poolId)
=======
		poolId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		foundPoolPtr, err := cisClient.Pools().GetPool(rs.Primary.Attributes["cis_id"], poolId)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		if err != nil {
			return err
		}

<<<<<<< HEAD
		pool = foundPool

=======
		foundPool := *foundPoolPtr
		if foundPool.Id != poolId {
			return fmt.Errorf("Record not found")
		}

		tfPool := convertCisToTfTwoVar(foundPool.Id, cisId)
		*tfPoolId = tfPool
		return nil
	}
}

func testAccCisPoolManuallyDelete(tfPoolId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Printf("[WARN] Manually removing pool")
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfPool := *tfPoolId
		poolId, cisId, _ := convertTftoCisTwoVar(tfPool)
		err = cisClient.Pools().DeletePool(cisId, poolId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISPool Record: %s", err)
		}
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		return nil
	}
}

<<<<<<< HEAD
func testAccCheckCisPoolConfigBasic(resourceId string, cis_domain string) string {
	return testAccCheckIBMCisInstance_basic(cis_domain) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
=======
func testAccCheckCisPoolConfigCisDS_Basic(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
  name = "my-tf-pool-basic-%[1]s"
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  origins {
    name = "example-1"
    address = "www.google.com"
    enabled = true
    weight = 1
  }
}
`, resourceId, cisInstance)
}

func testAccCheckCisPoolConfigCisRI_Basic(resourceId string, cisDomain string) string {
	return testAccCheckCisDomainConfigCisRI_basic(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${ibm_cis.%[2]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
  name = "my-tf-pool-basic-%[1]s"
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  origins {
    name = "example-1"
    address = "www.google.com"
    enabled = true
    weight = 1
  }
}
<<<<<<< HEAD
`, resourceId)
}

func testAccCheckCisPoolConfigFullySpecified(resourceId string, cis_domain string) string {
	return testAccCheckCisHealthcheckConfigBasic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
=======
`, resourceId, "testacc_ds_cis")
}

func testAccCheckCisPoolConfigFullySpecified(resourceId string, cisDomainStatic string) string {
	return testAccCheckCisHealthcheckConfigCisDS_Basic(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_origin_pool" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
  name = "my-tf-pool-basic-%[1]s"
  notification_email = "admin@outlook.com"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
  origins {
    name = "example-2"
    address = "192.0.2.2"
    enabled = true
  }
  check_regions = ["WEU"]
  description = "tfacc-fully-specified"
  enabled = false
  minimum_origins = 2
  monitor = "${ibm_cis_healthcheck.%[1]s.id}"
<<<<<<< HEAD
}`, resourceId)
=======
}`, resourceId, cisInstance)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}
