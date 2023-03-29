package infoblox

import (
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
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "dns_view", "nondefault_dnsview1"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "ptrdname", "rec1.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "ip_addr", "2002:1f93::10"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "ttl", "300"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "comment", "workstation #1"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds1", "ext_attrs", "{\"Location\":\"the main office\"}"),

					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "ptrdname", "rec2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "ip_addr", "10.0.0.101"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "ttl", "301"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "comment", "workstation #2"),
					resource.TestCheckResourceAttr("data.infoblox_ptr_record.ds2", "ext_attrs", "{\"Location\":\"the external office\"}"),
				),
			},
		},
	})
}

var testAccDataSourcePTRRecordsRead = `
resource "infoblox_ptr_record" "rec1" {
  ptrdname = "rec1.test.com"
  dns_view = "nondefault_dnsview1"
  ip_addr = "2002:1f93::10"
  comment = "workstation #1"
  ttl = 300
  ext_attrs = jsonencode({
    "Location" = "the main office"
  })
}

data "infoblox_ptr_record" "ds1" {
  ptrdname = "rec1.test.com"
  dns_view = "nondefault_dnsview1"
  ip_addr = "2002:1f93::10"

  depends_on = [infoblox_ptr_record.rec1]
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
}

data "infoblox_ptr_record" "ds2" {
  ptrdname = "rec2.test.com"
  // the default 'dns_view'
  ip_addr = "10.0.0.101"

  depends_on = [infoblox_ptr_record.rec2]
}
`
