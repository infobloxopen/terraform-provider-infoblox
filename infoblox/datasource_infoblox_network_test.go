package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccDataSourceNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNetworkRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_network.acctest", "cidr", "10.4.20.0/24"),
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
	network_view = "default"
  	cidr = infoblox_ipv4_network.test_network.cidr
}
`)
