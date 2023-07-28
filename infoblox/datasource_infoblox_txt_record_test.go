package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTXTRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceTXTRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.fqdn", "test-name1.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.text", "some text for a TXT-record"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.ttl", "30"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.comment", "this is a test TXT-record"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "results.0.ext_attrs", "{\"Site\":\"Greenland\"}"),

					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "results.0.dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "results.0.fqdn", "test-name2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "results.0.zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "results.0.text", "some text for a TXT-record 2"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "results.0.ttl", fmt.Sprintf("%d", ttlUndef)),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "results.0.comment", ""),
				),
			},
		},
	})
}

var testAccDataSourceTXTRecordsRead = fmt.Sprintf(`
resource "infoblox_txt_record" "rec1"{
	dns_view = "default"
	fqdn = "test-name1.test.com"
	text = "some text for a TXT-record"
    ttl = 30
    comment = "this is a test TXT-record"
    ext_attrs = jsonencode({
      "Site": "Greenland"
    })
}

data "infoblox_txt_record" "ds1"{
	filters = {
		view= "default"
	    name= infoblox_txt_record.rec1.fqdn
	}

	depends_on = [infoblox_txt_record.rec1]
}

resource "infoblox_txt_record" "rec2"{
	dns_view = "nondefault_view"
	fqdn = "test-name2.test.com"
	text = "some text for a TXT-record 2"
}

data "infoblox_txt_record" "ds2"{
	filters = {
		view= "nondefault_view"
		name= infoblox_txt_record.rec2.fqdn
	}

	depends_on = [infoblox_txt_record.rec2]
}
`)

func TestAccDataSourceTXTRecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_txt_record" "txt1"{
						dns_view = "default"
						fqdn = "newtxt.test.com"
						text = "some text for a TXT-record"
						ttl = 30
						comment = "new sample txt-record"
						ext_attrs = jsonencode({
							"Site" = "sample text site"
						})
					}

					data "infoblox_txt_record" "dtxt1" {
						filters = {
							"*Site" = "sample text site"
						}
						depends_on = [infoblox_txt_record.txt1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_txt_record.dtxt1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.dtxt1", "results.0.fqdn", "newtxt.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.dtxt1", "results.0.text", "some text for a TXT-record"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.dtxt1", "results.0.ttl", "30"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.dtxt1", "results.0.comment", "new sample txt-record"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.dtxt1", "results.0.ext_attrs", "{\"Site\":\"sample text site\"}"),
				),
			},
		},
	})
}
