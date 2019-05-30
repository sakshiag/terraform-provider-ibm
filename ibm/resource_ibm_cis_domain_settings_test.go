package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccIBMCisSettings_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	//t.Parallel()

	name := "ibm_cis_domain_settings." + "test"

	resource.Test(t, resource.TestCase{
<<<<<<< HEAD
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisSettingsConfigBasic("test", cis_domain),
=======
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisSettingsConfigBasic1("test", cisDomainStatic),
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "on"),
					resource.TestCheckResourceAttr(name, "ssl", "full"),
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.2"),
				),
			},
<<<<<<< HEAD
=======
			{
				Config: testAccCheckCisSettingsConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "waf", "off"),
					resource.TestCheckResourceAttr(name, "ssl", "flexible"),
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.3"),
				),
			},
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		},
	})
}

<<<<<<< HEAD
func testAccCheckCisSettingsConfigBasic(id string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic("test", cis_domain) + fmt.Sprintf(`
resource "ibm_cis_domain_settings" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  domain_id = "${ibm_cis_domain.%[1]s.id}"
  "waf" = "on"
  "ssl" = "full"	
  "min_tls_version" = "1.2"
}`, id)
=======
func testAccCheckCisSettingsConfigBasic1(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1("test", cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_domain_settings" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
  domain_id = "${data.ibm_cis_domain.%[1]s.id}"
  waf = "on"
  ssl = "full"	
  min_tls_version = "1.2"
}`, id, cisInstance)
}

func testAccCheckCisSettingsConfigBasic2(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1("test", cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_domain_settings" "%[1]s" {
  cis_id = "${data.ibm_cis.%[2]s.id}"
  domain_id = "${data.ibm_cis_domain.%[1]s.id}"
  waf = "off"
  ssl = "flexible"	
  min_tls_version = "1.3"
}`, id, cisInstance)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
}
