package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMXRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceMXRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "fqdn", "test-name.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "mail_exchanger", "mx-test.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "preference", "25"),
					//resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "ttl", ""),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "comment", ""),

					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "fqdn", "test-name2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "mail_exchanger", "mx-test2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "preference", "25"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "ttl", "10"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "comment", "non-empty comment"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "ext_attrs", "{\"Site\":\"None\"}"),
				),
			},
		},
	})
}

var testAccDataSourceMXRecordsRead = fmt.Sprintf(`
resource "infoblox_mx_record" "rec1" {
	dns_view = "default"
	fqdn = "test-name.test.com"
	mail_exchanger = "mx-test.test.com"
	preference = 25
}

data "infoblox_mx_record" "rec1" {
	dns_view = infoblox_mx_record.rec1.dns_view
	fqdn = infoblox_mx_record.rec1.fqdn
	mail_exchanger = infoblox_mx_record.rec1.mail_exchanger
}

resource "infoblox_mx_record" "rec2" {
	dns_view = "nondefault_view"
	fqdn = "test-name2.test.com"
	mail_exchanger = "mx-test2.test.com"
	preference = 25

	ttl = 10
    comment = "non-empty comment"
    ext_attrs = jsonencode({
      "Site": "None"
    })
}

data "infoblox_mx_record" "rec2" {
	dns_view = infoblox_mx_record.rec2.dns_view
	fqdn = infoblox_mx_record.rec2.fqdn
	mail_exchanger = infoblox_mx_record.rec2.mail_exchanger
}
`)
