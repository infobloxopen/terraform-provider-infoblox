package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testAccDataSourceDtcServer = fmt.Sprintf(`resource "infoblox_dtc_server" "server11"{
	name = "testServer_read1"
	host = "2.3.3.5"
}
data "infoblox_dtc_server" "testServer_read1" {	
	filters = {
	    name = infoblox_dtc_server.server11.name
    }
    depends_on = [infoblox_dtc_server.server11]
}`)

func TestAccDataSourceDtcServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDtcServer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read1", "results.0.name", "testServer_read1"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read1", "results.0.host", "2.3.3.5")),
			},
		},
	})
}

func TestAccDataSourceDtcServerSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "infoblox_dtc_server" "server12"{
									name = "terraform_server.com"
									host = "2.3.4.5"
  									comment = "testing server terraform"
  									sni_hostname = "sni_host"
  									use_sni_hostname = true
  									monitors {
        								host = "2.3.4.5"
        								monitor_name = "sip"
        								monitor_type = "sip"
									}
									monitors {
        								host = "2.3.4.32"
        								monitor_name = "http"
        								monitor_type = "http"
									}
  									ext_attrs = jsonencode({
    								"Site" = "Blr"
  									})
  									auto_create_host_record= true
									disable = true
							}
							data "infoblox_dtc_server" "testServer_read" {	
							filters = {
	   			 				name = infoblox_dtc_server.server12.name
    							}
							depends_on=[infoblox_dtc_server.server12]
						}
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.name", "terraform_server.com"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.host", "2.3.4.5"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.comment", "testing server terraform"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.use_sni_hostname", "true"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.sni_hostname", "sni_host"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.monitors.0.host", "2.3.4.5"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.monitors.0.monitor_name", "sip"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.monitors.0.monitor_type", "sip"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.monitors.1.host", "2.3.4.32"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.monitors.1.monitor_name", "http"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.monitors.1.monitor_type", "http"),
					resource.TestCheckResourceAttrPair("data.infoblox_dtc_server.testServer_read", "results.0.ext_attrs.Site", "infoblox_dtc_server.server12", "ext_attrs.Site"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.auto_create_host_record", "true"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_server.testServer_read", "results.0.disable", "true"),
				),
			}},
	})
}
