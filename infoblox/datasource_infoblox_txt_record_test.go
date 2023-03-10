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
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "fqdn", "test-name1.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "text", "some text for a TXT-record"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "ttl", "30"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "comment", "this is a test TXT-record"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds1", "ext_attrs", "{\"Site\":\"Greenland\"}"),

					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "fqdn", "test-name2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "text", "some text for a TXT-record 2"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "ttl", fmt.Sprintf("%d", ttlUndef)),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.ds2", "comment", ""),
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
	dns_view="default"
	fqdn=infoblox_txt_record.rec1.fqdn
}

resource "infoblox_txt_record" "rec2"{
	dns_view = "nondefault_view"
	fqdn = "test-name2.test.com"
	text = "some text for a TXT-record 2"
}

data "infoblox_txt_record" "ds2"{
	dns_view="nondefault_view"
	fqdn=infoblox_txt_record.rec2.fqdn
}
`)
