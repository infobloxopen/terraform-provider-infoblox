package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSRVRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceSRVRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "name", "_http._tcp.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "target", "www.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "port", "80"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "priority", "50"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "weight", "40"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "ttl", "10"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "comment", "testing SRV-record datasource"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec1", "ext_attrs", "{\"Site\":\"Moon\"}"),

					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "name", "_http._tcp.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "target", "www2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "port", "8080"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "priority", "50"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "weight", "40"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "ttl", fmt.Sprintf("%d", ttlUndef)),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.rec2", "comment", ""),
				),
			},
		},
	})
}

var testAccDataSourceSRVRecordsRead = fmt.Sprintf(`
resource "infoblox_srv_record" "rec1" {
	dns_view = "nondefault_view"
	name = "_http._tcp.test.com"
	priority = 50
	weight = 40
	port = 80
	target = "www.test.com"
    ttl = 10
    comment = "testing SRV-record datasource"
    ext_attrs = jsonencode({
      "Site": "Moon"
    })
}

data "infoblox_srv_record" "rec1" {
	dns_view = infoblox_srv_record.rec1.dns_view
	name = infoblox_srv_record.rec1.name
	target = infoblox_srv_record.rec1.target
	port = infoblox_srv_record.rec1.port
}

resource "infoblox_srv_record" "rec2" {
	name = "_http._tcp.test.com"
	priority = 50
	weight = 40
	port = 8080
	target = "www2.test.com"
}

data "infoblox_srv_record" "rec2" {
	dns_view = infoblox_srv_record.rec2.dns_view
	name = infoblox_srv_record.rec2.name
	target = infoblox_srv_record.rec2.target
	port = infoblox_srv_record.rec2.port
}
`)
