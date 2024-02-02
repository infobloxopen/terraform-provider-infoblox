package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceDNSView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDNSViewsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dns_view.acctest", "results.0.name", "non_defaultview"),
					resource.TestCheckResourceAttr("data.infoblox_dns_view.acctest", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_dns_view.acctest", "results.0.comment", "test dns view example"),
				),
			},
		},
	})
}

var testAccDataSourceDNSViewsRead = fmt.Sprintf(`
resource "infoblox_dns_view" "foo"{
	name = "non_defaultview"
	network_view = "default" 
	comment = "test dns view example"
}

data "infoblox_dns_view" "acctest" {
	filters = {
		name = "non_defaultview"
		network_view = "default"
	}
	depends_on = [infoblox_dns_view.foo]
}
`)

func TestAccDataSourceDNSViewSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_dns_view" "dview1"{
						name = "customview"
						network_view = "default"
						comment = "new dns view example"
						ext_attrs = jsonencode({
							"Site" = "Main DNS Site"
						})
					}

					data "infoblox_dns_view" "accview" {
						filters = {
							"*Site" = "Main DNS Site"
						}
						depends_on = [infoblox_dns_view.dview1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dns_view.accview", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_dns_view.accview", "results.0.name", "customview"),
					resource.TestCheckResourceAttr("data.infoblox_dns_view.accview", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_dns_view.accview", "results.0.comment", "new dns view example"),
					resource.TestCheckResourceAttr("data.infoblox_dns_view.accview", "results.0.ext_attrs", "{\"Site\":\"Main DNS Site\"}"),
				),
			},
		},
	})
}
