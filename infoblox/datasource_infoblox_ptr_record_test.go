package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePTRRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourcePTRRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "results.0.dns_view", "nondefault_dnsview1"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "results.0.ptrdname", "rec1.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "results.0.ip_addr", "2002:1f93::10"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "results.0.ttl", "300"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "results.0.comment", "workstation #1"),
					resource.TestCheckResourceAttrPair("data.infoblox_ptr_record.ds1", "results.0.ext_attrs.Location", "infoblox_ptr_record.rec1", "ext_attrs.Location"),

					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "results.0.ptrdname", "rec2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "results.0.ip_addr", "10.0.0.101"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "results.0.ttl", "301"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "results.0.comment", "workstation #2"),
					resource.TestCheckResourceAttrPair("data.infoblox_ptr_record.ds2", "results.0.ext_attrs.Location", "infoblox_ptr_record.rec2", "ext_attrs.Location"),
				),
			},
		},
	})
}

var testAccDataSourcePTRRecordsRead = `
resource "infoblox_dns_view" "view" {
	name = "nondefault_dnsview1"
}

resource "infoblox_zone_auth" "test" {
	fqdn = "test.com"
	view = "nondefault_dnsview1"
	depends_on = [infoblox_dns_view.view]
}

resource "infoblox_zone_auth" "itest" {
	fqdn = "2002:1f93::/64"
	view = "nondefault_dnsview1"
	zone_format = "IPV6"
	depends_on = [infoblox_dns_view.view]
}

resource "infoblox_ptr_record" "rec1" {
  ptrdname = "rec1.test.com"
  dns_view = "nondefault_dnsview1"
  ip_addr = "2002:1f93::10"
  comment = "workstation #1"
  ttl = 300
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
	  depends_on = [infoblox_zone_auth.test, infoblox_zone_auth.itest]
}

data "infoblox_ptr_record" "ds1" {
	filters = {
		ptrdname = "rec1.test.com"
		view = "nondefault_dnsview1"
		ipv6addr = "2002:1f93::10"
	}
	
	depends_on = [infoblox_ptr_record.rec1]
}

resource "infoblox_zone_auth" "test1" {
	fqdn = "test.com"
}

resource "infoblox_zone_auth" "itest1" {
	fqdn = "10.0.0.0/24"
	zone_format = "IPV4"
	depends_on = [infoblox_zone_auth.test1]
}
resource "infoblox_ptr_record" "rec2" {
	ptrdname = "rec2.test.com"
	// the default 'dns_view'
	ip_addr = "10.0.0.101"
	comment = "workstation #2"
	ttl = 301
	ext_attrs = jsonencode({
		"Location" = "the external office"
	})
	  depends_on = [infoblox_zone_auth.test1, infoblox_zone_auth.itest1]
}

data "infoblox_ptr_record" "ds2" {
  filters = {
	ptrdname = "rec2.test.com"
    // the default 'dns_view'
    ipv4addr = "10.0.0.101"
  }

  depends_on = [infoblox_ptr_record.rec2]
}
`

func TestAccDataSourcePTRRecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test" {
						fqdn = "test.com"
					}
					resource "infoblox_zone_auth" "itest" {
						fqdn = "10.0.0.0/24"
						zone_format = "IPV4"
						depends_on = [infoblox_zone_auth.test]
					}
					resource "infoblox_ptr_record" "ptr1"{
						ptrdname = "decoyptr.test.com"
						ip_addr = "10.0.0.89"
						dns_view = "default"
						comment = "new sample ptr-record"
						ext_attrs = jsonencode({
							"Location" = "PTR test location"
						})
						depends_on = [infoblox_zone_auth.itest, infoblox_zone_auth.test]
					}

					data "infoblox_ptr_record" "dptr1" {
						filters = {
							"*Location" = "PTR test location"
						}
						depends_on = [infoblox_ptr_record.ptr1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.dptr1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.dptr1", "results.0.ip_addr", "10.0.0.89"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.dptr1", "results.0.ptrdname", "decoyptr.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.dptr1", "results.0.comment", "new sample ptr-record"),
					resource.TestCheckResourceAttrPair("data.infoblox_ptr_record.dptr1", "results.0.ext_attrs.Location", "infoblox_ptr_record.ptr1", "ext_attrs.Location"),
				),
			},
		},
	})
}
