package ibm

import (
	//"errors"
	"fmt"
<<<<<<< HEAD
=======
	"log"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	"testing"

	//"regexp"

<<<<<<< HEAD
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
=======
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisGlb_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	//t.Parallel()
<<<<<<< HEAD
	var glb v1.Glb
=======
	var glb string
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
<<<<<<< HEAD
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		// If the DNS record failed to delete, the destroy of resource_ibm_cis used in this test suite will have been failed by the Resource Manager
		// and test execution aborted prior to this test.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigBasic("test", cis_domain),
=======
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
				),
			},
		},
	})
}

<<<<<<< HEAD
=======
func TestAccIBMCisGlb_CreateAfterManualDestroy(t *testing.T) {
	//t.Parallel()
	var glbOne, glbTwo string
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbOne),
					testAccCisGlbManuallyDelete(&glbOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbTwo),
					func(state *terraform.State) error {
						if glbOne == glbTwo {
							return fmt.Errorf("load balancer id is unchanged even after we thought we deleted it ( %s )",
								glbTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccIBMCisGlb_CreateAfterManualCisRIDestroy(t *testing.T) {
	//t.Parallel()
	var glbOne, glbTwo string
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisGlbConfigCisRI_Basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbOne),
					testAccCisGlbManuallyDelete(&glbOne),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_pool" {
								log.Printf("[WARN]  Manually removing pool")
								poolId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Pools().DeletePool(cisId, poolId)

							}

						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_domain" {
								log.Printf("[WARN] Manually removing domain")
								zoneId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisId, zoneId)
								cisPtr := &cisId
								log.Printf("[WARN]  Manually removing Cis Instance")
								_ = testAccCisInstanceManuallyDeleteUnwrapped(state, cisPtr)
							}

						}
						return nil
					},
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisGlbConfigCisRI_Basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glbTwo),
					func(state *terraform.State) error {
						if glbOne == glbTwo {
							return fmt.Errorf("load balancer id is unchanged even after we thought we deleted it ( %s )",
								glbTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
func TestAccIBMCisGlb_import(t *testing.T) {
	name := "ibm_cis_global_load_balancer.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
<<<<<<< HEAD
				Config: testAccCheckCisGlbConfigBasic("test", cis_domain),
=======
				Config: testAccCheckCisGlbConfigCisDS_Basic("test", cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
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

func TestAccIBMCisGlb_SessionAffinity(t *testing.T) {
<<<<<<< HEAD
	t.Parallel()
	var glb v1.Glb
=======
	//t.Parallel()
	var glb string
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	name := "ibm_cis_global_load_balancer." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
<<<<<<< HEAD
				Config: testAccCheckCisGlbConfigSessionAffinity("test", cis_domain),
=======
				Config: testAccCheckCisGlbConfigSessionAffinity("test", cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisGlbExists(name, &glb),
					// explicitly verify that our session_affinity has been set
					resource.TestCheckResourceAttr(name, "session_affinity", "cookie"),
					//resource.TestCheckResourceAttr(name, "pop_pools.#", "0"),
					//resource.TestCheckResourceAttr(name, "region_pools.#", "0"),
				),
			},
		},
	})
}

<<<<<<< HEAD
func testAccCheckCisGlbExists(n string, glb *v1.Glb) resource.TestCheckFunc {
=======
func testAccCheckCisGlbDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_global_load_balancer" {
			continue
		}
		glbId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		_, err = cisClient.Glbs().GetGlb(cisId, zoneId, glbId)
		if err == nil {
			return fmt.Errorf("Global Load balancer still exists")
		}
	}

	return nil
}

func testAccCisGlbManuallyDelete(tfGlbId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfGlb := *tfGlbId
		log.Printf("[WARN] Manually removing glb")
		glbId, zoneId, cisId, _ := convertTfToCisThreeVar(tfGlb)
		err = cisClient.Glbs().DeleteGlb(cisId, zoneId, glbId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISGlb Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisGlbExists(n string, tfGlbId *string) resource.TestCheckFunc {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Load Balancer ID is set")
		}
<<<<<<< HEAD
		glbId, zoneId, _, _ := convertTfToCisThreeVar(rs.Primary.ID)
=======
		glbId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		foundGlb, err := cisClient.Glbs().GetGlb(rs.Primary.Attributes["cis_id"], zoneId, glbId)
		if err != nil {
			return err
		}
<<<<<<< HEAD

		glb = foundGlb

=======
		*tfGlbId = convertCisToTfThreeVar(foundGlb.Id, zoneId, cisId)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		return nil
	}
}

<<<<<<< HEAD
func testAccCheckCisGlbConfigBasic(id string, cis_domain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  domain_id = "${ibm_cis_domain.%[1]s.id}"
  name = "%[2]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[1]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[1]s.id}"]
}`, id, cis_domain)
}

func testAccCheckCisGlbConfigSessionAffinity(id string, cis_domain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
=======
func testAccCheckCisGlbConfigCisDS_Basic(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${data.ibm_cis.%[3]s.id}"
  domain_id = "${data.ibm_cis_domain.%[1]s.id}"
  name = "%[2]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[1]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[1]s.id}"]
}`, id, cisDomainStatic, cisInstance)
}

func testAccCheckCisGlbConfigCisRI_Basic(id string, cisDomain string) string {
	return testAccCheckCisPoolConfigCisRI_Basic(id, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${ibm_cis.%[3]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
  domain_id = "${ibm_cis_domain.%[1]s.id}"
  name = "%[2]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[1]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[1]s.id}"]
<<<<<<< HEAD
  session_affinity = "cookie"
}`, id, cis_domain)
=======
}`, id, cisDomain, "testacc_ds_cis")
}

func testAccCheckCisGlbConfigSessionAffinity(id string, cisDomainStatic string) string {
	return testAccCheckCisPoolConfigFullySpecified(id, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_global_load_balancer" "%[1]s" {
  cis_id = "${data.ibm_cis.%[3]s.id}"
  domain_id = "${data.ibm_cis_domain.%[1]s.id}"
  name = "%[2]s"
  fallback_pool_id = "${ibm_cis_origin_pool.%[1]s.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.%[1]s.id}"]
  session_affinity = "cookie"
}`, id, cisDomainStatic, cisInstance)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}
