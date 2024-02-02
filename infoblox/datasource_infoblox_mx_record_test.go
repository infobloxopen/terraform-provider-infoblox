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
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "results.0.fqdn", "test-name.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "results.0.mail_exchanger", "mx-test.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "results.0.preference", "25"),
					//resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "ttl", ""),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec1", "results.0.comment", ""),

					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.fqdn", "test-name2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.mail_exchanger", "mx-test2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.preference", "25"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.ttl", "10"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.comment", "non-empty comment"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.rec2", "results.0.ext_attrs", "{\"Site\":\"None\"}"),
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
	filters = {
		view = infoblox_mx_record.rec1.dns_view
		name = infoblox_mx_record.rec1.fqdn
		mail_exchanger = infoblox_mx_record.rec1.mail_exchanger
		preference = infoblox_mx_record.rec1.preference
	}

	depends_on = [infoblox_mx_record.rec1]
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
	filters = {
		view = infoblox_mx_record.rec2.dns_view
		name = infoblox_mx_record.rec2.fqdn
		mail_exchanger = infoblox_mx_record.rec2.mail_exchanger
		preference = infoblox_mx_record.rec2.preference
	}

	depends_on = [infoblox_mx_record.rec2]
}
`)

func TestAccDataSourceMXRecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_mx_record" "mx1"{
						fqdn = "samplemx.test.com"
						mail_exchanger = "exdemo.test.com"
						preference = 30
						ttl = 10
						dns_view = "default"
						comment = "new sample mx-record"
						ext_attrs = jsonencode({
							"Site": "Some automated site"
						})
					}

					data "infoblox_mx_record" "dmx1" {
						filters = {
							"*Site" = "Some automated site"
						}
						depends_on = [infoblox_mx_record.mx1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_mx_record.dmx1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.dmx1", "results.0.mail_exchanger", "exdemo.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.dmx1", "results.0.preference", "30"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.dmx1", "results.0.fqdn", "samplemx.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.dmx1", "results.0.comment", "new sample mx-record"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.dmx1", "results.0.ext_attrs", "{\"Site\":\"Some automated site\"}"),
				),
			},
		},
	})
}
