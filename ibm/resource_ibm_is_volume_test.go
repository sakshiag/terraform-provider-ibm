package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/storage"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISVolume_basic(t *testing.T) {
	var volume *models.VolumeExtended
	name1 := fmt.Sprintf("terraformvpcvolumeuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.testacc_volume", &volume),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.testacc_volume", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	volumeC := storage.NewVolumeClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_volume" {
			continue
		}

		_, err := volumeC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("volume still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVolumeExists(n string, volume **models.VolumeExtended) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		volumeC := storage.NewVolumeClient(sess)
		foundVolume, err := volumeC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*volume = foundVolume
		return nil
	}
}

func testAccCheckIBMISVolumeConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_volume" "testacc_volume" {
    name = "%s"
    type = "boot"
    zone = "test"
    iops = 10000
    capacity = 100
    auto_delete = true
}`, name)

}
