package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceNetworkContainer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNetworkContainerRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.0.cidr", "10.4.20.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.0.comment", "network container #1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.0.ext_attrs", "{\"Location\":\"North Pole\"}"),
				),
			},
		},
	})
}

var testAccDataSourceNetworkContainerRead = fmt.Sprintf(`
resource "infoblox_ipv4_network_container" "nc1"{
  	cidr = "10.4.20.0/24"
    comment = "network container #1"
    ext_attrs = jsonencode({
      "Location": "North Pole"
    })
}

data "infoblox_ipv4_network_container" "ds1" {
	filters = {
		network_view = infoblox_ipv4_network_container.nc1.network_view
		network = infoblox_ipv4_network_container.nc1.cidr
	}
}
`)

func TestAccDataSourceNetworkCOntainerSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_network_container" "nc1"{
  						cidr = "10.4.20.0/24"
    					comment = "network container #1"
    					ext_attrs = jsonencode({
      						"Location": "North Pole"
    					})
					}

					data "infoblox_ipv4_network_container" "ds1" {
						filters = {
							"*Location" = "North Pole"
						}
						depends_on = [infoblox_ipv4_network_container.nc1]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.0.cidr", "10.4.20.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.0.comment", "network container #1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network_container.ds1", "results.0.ext_attrs", "{\"Location\":\"North Pole\"}"),
				),
			},
		},
	})
}
