package ibm

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMContainerCluster_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_basic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "kube_version", kubeVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers.0.version", kubeVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers.1.version", kubeVersion),
				),
			},
			{
				Config: testAccCheckIBMContainerCluster_update(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "kube_version", kubeUpdateVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers.0.version", kubeUpdateVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers.1.version", kubeVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "is_trusted", "false"),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_trusted(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_trusted(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "kube_version", kubeVersion),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "is_trusted", "true"),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_nosubnet_false(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_nosubnet_false(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "ingress_hostname"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "ingress_secret"),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_worker_count(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_worker_count(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMContainerCluster_worker_count_update(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "workers_info.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_without_workers_worker_num(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMContainerCluster_without_workers_worker_num(clusterName),
				ExpectError: regexp.MustCompile(" Please set either the wokers with valid array"),
			},
		},
	})
}

func TestAccIBMContainerCluster_with_worker_num_zero(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMContainerCluster_with_worker_num_zero(clusterName),
				ExpectError: regexp.MustCompile("must be greater than 0"),
			},
		},
	})
}

func TestAccIBMContainerCluster_diskEnc(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_diskEnc(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
				),
			},
		},
	})
}

//testAccCheckIBMContainerClusterOptionalOrgSpace_basic
func TestAccIBMContainerClusterOptionalOrgSpace_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterOptionalOrgSpace_basic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "2"),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_private_subnet(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_private_subnet(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "ingress_hostname", ""),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "ingress_secret", ""),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_private_and_public_subnet(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerCluster_private_and_public_subnet(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "ingress_hostname"),
					resource.TestCheckResourceAttrSet(
						"ibm_container_cluster.testacc_cluster", "ingress_secret"),
				),
			},
		},
	})
}

func TestAccIBMContainerCluster_Tag(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMContainerClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterTag(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMContainerClusterUpdateTag(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "name", clusterName),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "worker_num", "1"),
					resource.TestCheckResourceAttr(
						"ibm_container_cluster.testacc_cluster", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDestroy(s *terraform.State) error {
	csClient, err := testAccProvider.Meta().(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_cluster" {
			continue
		}

		targetEnv := getClusterTargetHeaderTestACC()
		// Try to find the key
		_, err := csClient.Clusters().Find(rs.Primary.ID, targetEnv)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for cluster (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func getClusterTargetHeaderTestACC() v1.ClusterTargetHeader {
	org := cfOrganization
	space := cfSpace
	c := new(bluemix.Config)
	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	client, err := mccpv2.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org, BluemixRegion)

	if err != nil {
		log.Fatal(err)
	}

	spaceAPI := client.Spaces()
	myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space, BluemixRegion)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, c.Region)
	if err != nil {
		log.Fatal(err)
	}

	target := v1.ClusterTargetHeader{
		OrgID:     myorg.GUID,
		SpaceID:   myspace.GUID,
		AccountID: myAccount.GUID,
	}

	return target
}

func testAccCheckIBMContainerCluster_basic(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  },{
    name = "worker2"
    }]

  kube_version    = "%s"
  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, kubeVersion, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_trusted(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  worker_num      = 2

  kube_version    = "%s"
  machine_type    = "%s"
  isolation       = "private"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
  is_trusted  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, kubeVersion, trustedMachineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_nosubnet_false(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  },{
    name = "worker2"
    }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = false
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_without_workers_worker_num(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  account_guid = "${data.ibm_account.acc.id}"

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_with_worker_num_zero(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  account_guid = "${data.ibm_account.acc.id}"
  worker_num = 0
  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterOptionalOrgSpace_basic(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  },{
    name = "worker2"
    }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
  disk_encryption = true
}	`, cfOrganization, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_diskEnc(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  },{
    name = "worker2"
    }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
   disk_encryption = false
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_update(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
    version = "%s"
    },{
    name = "worker2"
    }]

  kube_version    = "%s"
  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, kubeUpdateVersion, kubeUpdateVersion, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_private_and_public_subnet(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  subnet_id		  = ["%s","%s"]
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID, privateSubnetID, publicSubnetID)
}

func testAccCheckIBMContainerCluster_private_subnet(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
  subnet_id		  = ["%s"]
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID, privateSubnetID)
}

func testAccCheckIBMContainerClusterTag(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  tags = ["test"]
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerClusterUpdateTag(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

  org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"

  workers = [{
    name = "worker1"
  }]

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  tags = ["test","once"]
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_worker_count(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

	account_guid = "${data.ibm_account.acc.id}"

  worker_num = 1

  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}

func testAccCheckIBMContainerCluster_worker_count_update(clusterName string) string {
	return fmt.Sprintf(`

data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name       = "%s"
  datacenter = "%s"

	account_guid = "${data.ibm_account.acc.id}"
  worker_num = 2
  machine_type    = "%s"
  isolation       = "public"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  no_subnet		  = true
}	`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID)
}
