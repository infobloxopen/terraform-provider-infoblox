package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceARecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "results.0.zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "results.0.fqdn", "test-name.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "results.0.ip_addr", "10.0.0.20"),
				),
			},
		},
	})
}

var testAccDataSourceARecordsRead = fmt.Sprintf(`
resource "infoblox_zone_auth" "test" {
	fqdn = "test.com"
}

resource "infoblox_a_record" "foo"{
	dns_view="default"
	fqdn="test-name.test.com"
	ip_addr="10.0.0.20"
	depends_on = [infoblox_zone_auth.test]
}

data "infoblox_a_record" "acctest" {
	filters = {
		view="default"
		name=infoblox_a_record.foo.fqdn
		ipv4addr=infoblox_a_record.foo.ip_addr
	}
}
`)

func TestAccDataSourceARecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test" {
						fqdn = "test.com"
					}
					resource "infoblox_a_record" "arec1"{
						ip_addr = "10.4.0.18"
						fqdn = "sample.test.com"
						dns_view = "default"
						comment = "test sample A-record"
						ext_attrs = jsonencode({
							"Location": "Las Vegas"
						})
						depends_on = [infoblox_zone_auth.test]
					}

					data "infoblox_a_record" "ds1" {
						filters = {
							"*Location" = "Las Vegas"
						}
						depends_on = [infoblox_a_record.arec1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_a_record.ds1", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.ds1", "results.0.ip_addr", "10.4.0.18"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.ds1", "results.0.fqdn", "sample.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.ds1", "results.0.comment", "test sample A-record"),
					resource.TestCheckResourceAttrPair("data.infoblox_a_record.ds1", "results.0.ext_attrs.Location", "infoblox_a_record.arec1", "ext_attrs.Location"),
				),
			},
		},
	})
}
