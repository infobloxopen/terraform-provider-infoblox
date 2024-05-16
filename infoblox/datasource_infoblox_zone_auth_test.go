package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceZoneAuth(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceZoneAuthsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.fqdn", "test2.com"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.comment", "test forward mapping zone"),
				),
			},
			{
				Config: testAccDataSourceZoneAuthRead01,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.fqdn", "10.1.0.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.zone_format", "IPV4"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.acctest", "results.0.comment", "test ipv4 reverse mapping zone"),
				),
			},
		},
	})
}

var testAccDataSourceZoneAuthsRead = fmt.Sprintf(`
resource "infoblox_zone_auth" "foo"{
	fqdn = "test2.com"
	zone_format = "FORWARD"
	restart_if_needed = true
	soa_default_ttl = 36002
	soa_expire = 72002
	soa_negative_ttl = 602
	soa_refresh = 1802
	soa_retry = 902
	comment = "test forward mapping zone"
	ext_attrs = jsonencode({
		Location = "TestMapping"
	})
}

data "infoblox_zone_auth" "acctest" {
	filters = {
		view="default"
		fqdn=infoblox_zone_auth.foo.fqdn
	}
}
`)

var testAccDataSourceZoneAuthRead01 = fmt.Sprintf(`
resource "infoblox_zone_auth" "foo1"{
	fqdn = "10.1.0.0/24"
	zone_format = "IPV4"
	restart_if_needed = true
	soa_default_ttl = 15000
	soa_expire = 57000
	soa_negative_ttl = 500
	soa_refresh = 1705
	soa_retry = 850
	comment = "test ipv4 reverse mapping zone"
	ext_attrs = jsonencode({
		Site = "Test MapReverse"
	})
}

data "infoblox_zone_auth" "acctest" {
	filters = {
		view = "default"
		fqdn = infoblox_zone_auth.foo1.fqdn
		zone_format = "IPV4"
	}
}
`)

func TestAccDataSourceZoneAuthSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "zone1"{
						fqdn = "17.1.0.0/16"
						zone_format = "IPV4"
						view = "default"
						comment = "test sample reverse zone"
						ext_attrs = jsonencode({
							Location =  "Test Zone Location"
						})
					}

					data "infoblox_zone_auth" "dzone1" {
						filters = {
							"*Location" = "Test Zone Location"
						}
						depends_on = [infoblox_zone_auth.zone1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.dzone1", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.dzone1", "results.0.view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.dzone1", "results.0.fqdn", "17.1.0.0/16"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.dzone1", "results.0.comment", "test sample reverse zone"),
					resource.TestCheckResourceAttr("data.infoblox_zone_auth.dzone1", "results.0.zone_format", "IPV4"),
					resource.TestCheckResourceAttrPair("data.infoblox_zone_auth.dzone1", "results.0.ext_attrs.Location", "infoblox_zone_auth.zone1", "ext_attrs.Location"),
				),
			},
		},
	})
}
