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
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.name", "_http._tcp.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.target", "www.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.port", "80"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.priority", "50"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.weight", "40"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.ttl", "10"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec1", "results.0.comment", "testing SRV-record datasource"),
					resource.TestCheckResourceAttrPair("data.infoblox_srv_record.srec1", "results.0.ext_attrs.Site", "infoblox_srv_record.rec1", "ext_attrs.Site"),

					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.name", "_http._tcp.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.target", "www2.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.port", "8080"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.priority", "50"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.weight", "40"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.ttl", fmt.Sprintf("%d", ttlUndef)),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.srec2", "results.0.comment", ""),
				),
			},
		},
	})
}

var testAccDataSourceSRVRecordsRead = fmt.Sprintf(`
resource "infoblox_dns_view" "view" {
	name = "nondefault_view"
}
resource "infoblox_zone_auth" "test" {
	fqdn = "test.com"
	view = "nondefault_view"
	depends_on = [infoblox_dns_view.view]
}
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
	depends_on = [infoblox_zone_auth.test]
}

data "infoblox_srv_record" "srec1" {
	filters = {
		view = infoblox_srv_record.rec1.dns_view
		name = infoblox_srv_record.rec1.name
		target = infoblox_srv_record.rec1.target
		port = infoblox_srv_record.rec1.port
	}

	depends_on = [infoblox_srv_record.rec1]
}

resource "infoblox_zone_auth" "test1" {
	fqdn = "test.com"
}
resource "infoblox_srv_record" "rec2" {
	name = "_http._tcp.test.com"
	priority = 50
	weight = 40
	port = 8080
	target = "www2.test.com"
	depends_on = [infoblox_zone_auth.test1]
}

data "infoblox_srv_record" "srec2" {
	filters = {
		view = infoblox_srv_record.rec2.dns_view
		name = infoblox_srv_record.rec2.name
		target = infoblox_srv_record.rec2.target
		port = infoblox_srv_record.rec2.port
	}

	depends_on = [infoblox_srv_record.rec2]
}
`)

func TestAccDataSourceSRVRecordSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test" {
						fqdn = "test.com"
					}
					resource "infoblox_srv_record" "srv1"{
						name = "_http._udp.test.com"
						priority = 40
						weight = 55
						port = 5080
						target = "random.test.com"
						dns_view = "default"
						comment = "new sample srv-record"
						ext_attrs = jsonencode({
							"Site" = "test srv site"
						})
						depends_on = [infoblox_zone_auth.test]
					}

					data "infoblox_srv_record" "dsrv1" {
						filters = {
							"*Site" = "test srv site"
						}
						depends_on = [infoblox_srv_record.srv1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.name", "_http._udp.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.priority", "40"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.weight", "55"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.port", "5080"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.target", "random.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.dsrv1", "results.0.comment", "new sample srv-record"),
					resource.TestCheckResourceAttrPair("data.infoblox_srv_record.dsrv1", "results.0.ext_attrs.Site", "infoblox_srv_record.srv1", "ext_attrs.Site"),
				),
			},
		},
	})
}
