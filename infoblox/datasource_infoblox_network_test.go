package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNetworkRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.0.cidr", "10.4.20.0/24"),
				),
			},
		},
	})
}

var testAccDataSourceNetworkRead = fmt.Sprintf(`
resource "infoblox_ipv4_network" "test_network"{
  	cidr = "10.4.20.0/24"
}

data "infoblox_ipv4_network" "acctest" {
	filters = {
		network_view = "default"
		network = infoblox_ipv4_network.test_network.cidr
	}
}
`)

func TestAccDataSourceNetworkReadByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_network" "test_network" {
						cidr = "10.4.20.0/24"
						comment = "Created by terraform-provider-infoblox acceptance test"
						ext_attrs = jsonencode({
							Building = "AcceptanceTerraform"
						})
					}

					data "infoblox_ipv4_network" "acctest" {
						filters = {
							"*Building" = "AcceptanceTerraform"
						}
						depends_on  = [infoblox_ipv4_network.test_network]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.0.cidr", "10.4.20.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.0.comment", "Created by terraform-provider-infoblox acceptance test"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "results.0.ext_attrs", "{\"Building\":\"AcceptanceTerraform\"}"),
				),
			},
		},
	})
}
