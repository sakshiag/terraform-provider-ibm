package ibm

import (
	"fmt"
	"testing"

<<<<<<< HEAD
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
=======
	//v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisDNSRecord_Basic(t *testing.T) {
	//t.Parallel()
<<<<<<< HEAD
	var record v1.DnsRecord
=======
	var record string
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	testName := "tf-acctest-basic"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
<<<<<<< HEAD
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		// If the DNS record failed to delete, the destroy of resource_ibm_cis used in this test suite will have been failed by the Resource Manager
		// and test execution aborted prior to this test.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigBasic(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cis_domain),
=======
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_import(t *testing.T) {
	name := "ibm_cis_dns_record.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
<<<<<<< HEAD
				Config: testAccCheckIBMCisDNSRecordConfigBasic("test", cis_domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "name", "test."+cis_domain),
=======
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "name", "test."+cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr(name, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(name, "data.%", "0"),
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

func TestAccIBMCisDNSRecord_CaseInsensitive(t *testing.T) {
	//t.Parallel()
<<<<<<< HEAD
	var record v1.DnsRecord
	testName := "tf-acctest-case-insensitive"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", "test")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cis_domain),
				),
			},
			{
				Config:   testAccCheckIBMCisDNSRecordConfigCaseSensitive("tf-acctest-CASE-INSENSITIVE", cis_domain),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-case-insensitive"+"."+cis_domain),
=======
	var record string
	testName := "tf-acctest-case-insensitive"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cisDomainStatic),
				),
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive("tf-acctest-CASE-INSENSITIVE", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists("ibm_cis_dns_record.tf-acctest-CASE-INSENSITIVE", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_dns_record.tf-acctest-CASE-INSENSITIVE", "name", "tf-acctest-case-insensitive."+cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_Apex(t *testing.T) {
	//t.Parallel()
<<<<<<< HEAD
	var record v1.DnsRecord
=======
	var record string
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	testName := "test"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
<<<<<<< HEAD
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigApex(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						// @ is replaced by domain name by CIS
						resourceName, "name", cis_domain),
=======
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigApex(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						// @ is replaced by domain name by CIS
						resourceName, "name", cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
				),
			},
		},
	})
}

<<<<<<< HEAD
func testAccCheckIBMCisDNSRecordDestroy(s *terraform.State) error {
	cisClient, _ := testAccProvider.Meta().(ClientSession).CisAPI()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_dns_record" {
			continue
		}

		_, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], rs.Primary.Attributes["domain_id"], rs.Primary.ID)
=======
func TestAccIBMCisDNSRecord_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	testName := "test_acc"
	var afterCreate, afterRecreate string
	name := "ibm_cis_dns_record.test_acc"

	afterCreate = "hello"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterCreate),
					testAccIBMCisManuallyDeleteDnsRecord(&afterCreate),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterRecreate),
					testAccCheckIBMCisDnsRecordRecreated(&afterCreate, &afterRecreate),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_CreateAfterManualCisRIDestroy(t *testing.T) {
	t.Parallel()
	testName := "test_acc"
	var afterCreate, afterRecreate string
	name := "ibm_cis_dns_record.test_acc"

	afterCreate = "hello"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterCreate),
					testAccIBMCisManuallyDeleteDnsRecord(&afterCreate),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_domain" {
								zoneId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisId, zoneId)
								cisPtr := &cisId
								_ = testAccCisInstanceManuallyDelete(cisPtr)
							}

						}
						return nil
					},
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterRecreate),
					testAccCheckIBMCisDnsRecordRecreated(&afterCreate, &afterRecreate),
				),
			},
		},
	})
}

func testAccIBMCisManuallyDeleteDnsRecord(tfRecordId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfRecord := *tfRecordId
		recordId, zoneId, cisId, _ := convertTfToCisThreeVar(tfRecord)
		err = cisClient.Dns().DeleteDns(cisId, zoneId, recordId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISDNS Record: %s", err)
		}
		return nil
	}
}

func testAccCheckIBMCisDnsRecordRecreated(beforeId, afterId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *beforeId == *afterId {
			return fmt.Errorf("Expected change of Record Ids, but both were %v", beforeId)
		}
		return nil
	}
}

func testAccCheckIBMCisDnsRecordDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_record" {
			continue
		}

		recordId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		err = cisClient.Dns().DeleteDns(cisId, zoneId, recordId)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

<<<<<<< HEAD
func testAccCheckIBMCisDNSRecordExists(n string, record *v1.DnsRecord) resource.TestCheckFunc {
=======
func testAccCheckIBMCisDnsRecordExists(n string, tfRecordId *string) resource.TestCheckFunc {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

<<<<<<< HEAD
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		recordId, zoneId, _, _ := convertTfToCisThreeVar(rs.Primary.ID)
		foundRecord, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], zoneId, recordId)
=======
		tfRecord := *tfRecordId
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		recordId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		foundRecordPtr, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], zoneId, recordId)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		if err != nil {
			return err
		}

<<<<<<< HEAD
=======
		foundRecord := *foundRecordPtr
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		if foundRecord.Id != recordId {
			return fmt.Errorf("Record not found")
		}

<<<<<<< HEAD
		record = foundRecord

=======
		tfRecord = convertCisToTfThreeVar(foundRecord.Id, zoneId, cisId)
		*tfRecordId = tfRecord
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		return nil
	}
}

<<<<<<< HEAD
func testAccCheckIBMCisDNSRecordConfigBasic(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
=======
func testAccCheckIBMCisDNSRecordConfigCisDS_Basic(resourceId string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigCisRI_Basic(resourceId string, cisDomain string) string {
	return testAccCheckCisDomainConfigCisRI_basic(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.%[2]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
    domain_id = "${ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
<<<<<<< HEAD
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigCaseSensitive(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic("test", cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "test" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.test.id}"
=======
}`, resourceId, "testacc_ds_cis")
}

func testAccCheckIBMCisDNSRecordConfigCaseSensitive(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
<<<<<<< HEAD
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigApex(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
    name = "@"
    content = "192.168.0.10"
    type = "A"
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigLOC(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
=======
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigApex(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
    name = "@"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigLOC(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
    name = "%[1]s"
    data {
      "lat_degrees" =  "37"
      "lat_minutes" = "46"
      "lat_seconds" = "46"
      "lat_direction" = "N"
      "long_degrees" = "122"
      "long_minutes" = "23"
      "long_seconds" = "35"
      "long_direction" = "W"
      "altitude" = 0
      "size" = 100
      "precision_horz" = 0
      "precision_vert" = 0
    }
    type = "LOC"
<<<<<<< HEAD
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigSRV(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "foobar" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
=======
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigSRV(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
    name = "%[1]s"
    data {
      "priority" = 5
      "weight" = 0
      "port" = 5222
      "target" = "talk.l.google.com"
      "service" = "_xmpp-client"
      "proto" = "_tcp"
    }
    type = "SRV"
<<<<<<< HEAD
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigProxied(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "foobar" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
=======
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigProxied(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

    name = "%[1]s"
    content = "%[1]s"
    type = "CNAME"
    proxied = true
<<<<<<< HEAD
}`, resourceId)
=======
}`, resourceId, cisInstance)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}
