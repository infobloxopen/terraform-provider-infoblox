package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAAAARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceAAAARecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "results.0.zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "results.0.fqdn", "aaaa-test.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "results.0.ipv6_addr", "2002:1111::7"),
				),
			},
		},
	})
}

var testAccDataSourceAAAARecordsRead = fmt.Sprintf(`
resource "infoblox_zone_auth" "test" {
	fqdn = "test.com"
}
resource "infoblox_aaaa_record" "foo"{
	dns_view="default"
	fqdn="aaaa-test.test.com"
	ipv6_addr="2002:1111::7"
	depends_on = [infoblox_zone_auth.test]
}

data "infoblox_aaaa_record" "acctest" {
	filters = {
		view="default"
		name=infoblox_aaaa_record.foo.fqdn
		ipv6addr=infoblox_aaaa_record.foo.ipv6_addr
	}
}
`)

func TestAccDataSourceAAAARecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test" {
						fqdn = "test.com"
					}
					resource "infoblox_aaaa_record" "qarec1"{
						ipv6_addr = "2002:1f93:0:4::11"
						fqdn = "sampleqa.test.com"
						dns_view = "default"
						comment = "test sample AAAA-record"
						ext_attrs = jsonencode({
							"Location": "Norway"
						})
						depends_on = [infoblox_zone_auth.test]
					}

					data "infoblox_aaaa_record" "dqa1" {
						filters = {
							"*Location" = "Norway"
						}
						depends_on = [infoblox_aaaa_record.qarec1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.dqa1", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.dqa1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.dqa1", "results.0.ipv6_addr", "2002:1f93:0:4::11"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.dqa1", "results.0.fqdn", "sampleqa.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.dqa1", "results.0.comment", "test sample AAAA-record"),
					resource.TestCheckResourceAttrPair("data.infoblox_aaaa_record.dqa1", "results.0.ext_attrs.Location", "infoblox_aaaa_record.qarec1", "ext_attrs.Location"),
				),
			},
		},
	})
}
