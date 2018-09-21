package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISInstanceNic_basic(t *testing.T) {
	var instance *models.Instance
	vpcname := fmt.Sprintf("terraforminstanceuat_vpc_%d", acctest.RandInt())
	name := fmt.Sprintf("terraforminstanceuat-%d", acctest.RandInt())
	nicname := fmt.Sprintf("terraforminstancenicuat-%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraforminstanceuat_subnet_%d", acctest.RandInt())
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("terraformsecurityuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceNicDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceNicConfig(vpcname, subnetname, sshname, publicKey, name, nicname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceNicExists("ibm_is_instance.testacc_instance", &instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceNicDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	instanceC := compute.NewInstanceClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance" {
			continue
		}

		_, err := instanceC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("instance still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISInstanceNicExists(n string, instance **models.Instance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		instanceC := compute.NewInstanceClient(sess)
		foundinstance, err := instanceC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*instance = foundinstance
		return nil
	}
}

func testAccCheckIBMISInstanceNicConfig(vpcname, subnetname, sshname, publicKey, name, nicname string) string {
	return fmt.Sprintf(`
	
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "10.0.3.0/24"
	}
	
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name = "%s"
		public_key = "%s"
	}

	resource "ibm_is_instance" "testacc_instance" {
		name        = "%s"
		image       = "%s"
		profile     = "b-2x4"
		primary_network_interface = {
			port_speed  = "100"
			subnet      = "${ibm_is_subnet.testacc_subnet.id}"
		}
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		keys = ["${ibm_is_ssh_key.testacc_sshkey.id}"]
	}

	resource "ibm_is_instance_nic" "testacc_instance_nic" {
		name        = "%s"
		instance_id = "${ibm_is_instance.testacc_instance.id}"
		port_speed = "100"
		subnet = "${ibm_is_subnet.testacc_subnet.id}"
	}
	
`, vpcname, subnetname, ISZoneName, sshname, publicKey, name, isImage, ISZoneName, nicname)
}
