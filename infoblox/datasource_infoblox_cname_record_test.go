package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCNameRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCNameRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "results.0.zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "results.0.alias", "test.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "results.0.canonical", "test-name.test.com"),
				),
			},
		},
	})
}

var testAccDataSourceCNameRecordsRead = fmt.Sprintf(`
resource "infoblox_zone_auth" "test" {
	fqdn = "test.com"
}
resource "infoblox_cname_record" "foo"{
	dns_view="default"
	
	alias="test.test.com"
	canonical="test-name.test.com"
	depends_on = [infoblox_zone_auth.test]
  }

data "infoblox_cname_record" "acctest" {
	filters = {
		view="default"
		name= infoblox_cname_record.foo.alias
		canonical=infoblox_cname_record.foo.canonical
	}
  }
`)

func TestAccDataSourceCNameRecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test" {
						fqdn = "test.com"
					}
					resource "infoblox_cname_record" "cname1"{
						canonical = "somewhere.in.the.net"
						alias = "samplecname.test.com"
						dns_view = "default"
						comment = "test sample CName-record"
						ext_attrs = jsonencode({
							"Site": "test site one"
						})
						depends_on = [infoblox_zone_auth.test]
					}

					data "infoblox_cname_record" "dcname1" {
						filters = {
							"*Site" = "test site one"
						}
						depends_on = [infoblox_cname_record.cname1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_cname_record.dcname1", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.dcname1", "results.0.canonical", "somewhere.in.the.net"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.dcname1", "results.0.alias", "samplecname.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.dcname1", "results.0.comment", "test sample CName-record"),
					resource.TestCheckResourceAttrPair("data.infoblox_cname_record.dcname1", "results.0.ext_attrs.Site", "infoblox_cname_record.cname1", "ext_attrs.Site"),
				),
			},
		},
	})
}
