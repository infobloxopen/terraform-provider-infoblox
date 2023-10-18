package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceNetworkViewReadByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_network_view" "test_network_view"{
						name = "testNetworkView"
						comment = "test comment 1"
						ext_attrs = jsonencode({
							"Location" = "AcceptanceTerraform"
						})
					}

					data "infoblox_network_view" "acctest" {
						filters = {
							"*Location" = "AcceptanceTerraform"
						}
						depends_on  = [infoblox_network_view.test_network_view]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_network_view.acctest", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_network_view.acctest", "results.0.name", "testNetworkView"),
					resource.TestCheckResourceAttr("data.infoblox_network_view.acctest", "results.0.comment", "test comment 1"),
					resource.TestCheckResourceAttr("data.infoblox_network_view.acctest", "results.0.ext_attrs", "{\"Location\":\"AcceptanceTerraform\"}"),
				),
			},
		},
	})
}
