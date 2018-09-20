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

func TestAccIBMISVolumeSnapshotSnapshot_basic(t *testing.T) {
	var volumeSnapshot *models.VolumeSnapshot
	volname := fmt.Sprintf("terraformvpcvolumesnpvuat_create_step_name_%d", acctest.RandInt())
	snapname := fmt.Sprintf("terraformvpcvolumesnpsuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeSnapshotConfig(volname, snapname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeSnapshotExists("ibm_is_volume_snapshot.testacc_volume_snapshot", &volumeSnapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_volume_snapshot.testacc_volume_snapshot", "name", snapname),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeSnapshotDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	volumeC := storage.NewVolumeClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_volume_snapshot" {
			continue
		}

		volID, snapID, err := parseISTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = volumeC.GetSnapshot(volID, snapID)

		if err == nil {
			return fmt.Errorf("volume snapshot still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVolumeSnapshotExists(n string, volumeSnap **models.VolumeSnapshot) resource.TestCheckFunc {
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

		volID, snapID, err := parseISTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}

		foundVolumeSnap, err := volumeC.GetSnapshot(volID, snapID)

		if err != nil {
			return err
		}

		*volumeSnap = foundVolumeSnap
		return nil
	}
}

func testAccCheckIBMISVolumeSnapshotConfig(volname, snapname string) string {
	return fmt.Sprintf(`
resource "ibm_is_volume" "testacc_volume" {
    name = "%s"
    type = "boot"
    zone = "%s"
    iops = 10000
    capacity = 100
    auto_delete = true
}
resource "ibm_is_volume_snapshot" "testacc_volume_snapshot" {
    volume_id = "${ibm_is_volume.testacc_volume.id}"
    name = "%s"
}
`, volname, ISZoneName, snapname)

}
