package ibm

import (
	"fmt"
	"testing"

<<<<<<< HEAD
	"github.com/hashicorp/terraform/helper/acctest"
=======
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMCisDataSource_basic(t *testing.T) {
<<<<<<< HEAD
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  setupCisConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis.instance", "service", "internet-svcs"),
				),
			},
			resource.TestStep{
				Config:  testAccCheckIBMCisDataSourceConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "service", "internet-svcs"),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "plan", "standard"),
=======
	instanceName := fmt.Sprintf(cisInstance)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  testAccCheckIBMCisDataSourceConfig(instanceName),
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "service", "internet-svcs"),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "plan", "enterprise-usage"),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "location", "global"),
				),
			},
		},
	})
}

<<<<<<< HEAD
func setupCisConfig(instanceName string) string {
	return fmt.Sprintf(`

resource "ibm_cis" "instance" {
  name       = "%s"
  plan       = "standard"
  location   = "global"
}`, instanceName)

}

func testAccCheckIBMCisDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_cis" "instance" {
	name       = "%s"
	plan       = "standard"
	location   = "global"
}

data "ibm_cis" "testacc_ds_cis" {
  name = "${ibm_cis.instance.name}"
}
`, instanceName)
=======
func testAccCheckIBMCisDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  name = "%[1]s"
}

data "ibm_cis" "testacc_ds_cis" {
  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
  name = "%[2]s"
}`, cisResourceGroup, instanceName)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

}
