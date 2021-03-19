package infoblox

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNetworkCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_network.acctest", "network_name", "acctest-network"),
				),
			},
		},
	})
}

var testAccDataSourceNetworkCreate = `
resource "infoblox_network" "test_network"{
  network_name      = "acctest-network"
  cidr              = "10.4.20.0/24"
  tenant_id         = "test_tenant_id"
}

data "infoblox_network" "acctest" {
  cidr              = infoblox_network.test_network.cidr
  tenant_id         = infoblox_network.test_network.tenant_id
}
`
