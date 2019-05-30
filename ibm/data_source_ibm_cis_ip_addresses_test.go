package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"strconv"
	"testing"
)

func TestAccIBMCisIPDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckIBMCisIPDataSourceConfig_basic),
				Check: resource.ComposeTestCheckFunc(
<<<<<<< HEAD
					testAccIBMCisIPAddrs("data.ibm_cis_ip_addresses.wcpclouduk"),
=======
					testAccIBMCisIPAddrs("data.ibm_cis_ip_addresses.test_acc"),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				),
			},
		},
	})
}

func testAccIBMCisIPAddrs(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes
		log.Printf("%#v\n", a["ipv4_cidrs"])
		log.Printf("%#v\n", len(a["ipv4_cidrs"]))

		cidrs, _ := strconv.Atoi(a["ipv4_cidrs.#"])
		if cidrs == 0 {
			return fmt.Errorf("No ipv4 cidrs returned")
		}
		cidrs, _ = strconv.Atoi(a["ipv6_cidrs.#"])
		if cidrs == 0 {
			return fmt.Errorf("No ipv6 cidrs returned")
		}
		return nil
	}
}

const testAccCheckIBMCisIPDataSourceConfig_basic = `
<<<<<<< HEAD
data "ibm_cis_ip_addresses" "wcpclouduk" {
=======
data "ibm_cis_ip_addresses" "test_acc" {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}
`
