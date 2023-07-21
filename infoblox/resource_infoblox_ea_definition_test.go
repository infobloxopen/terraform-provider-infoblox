package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"testing"
)

func TestAcc_resourceEADefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEADefinitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_ea_definition" "ea_def" {
						name = "AcceptanceTestEA"
						type = "STRING"
						comment = "Acceptance test extensible attribute"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_ea_definition.ea_def", "name", `AcceptanceTestEA`),
					resource.TestCheckResourceAttr("infoblox_ea_definition.ea_def", "type", `STRING`),
					resource.TestCheckResourceAttr("infoblox_ea_definition.ea_def", "comment",
						`Acceptance test extensible attribute`),
				),
			},
			{
				Config: `
					resource "infoblox_ea_definition" "ea_def" {
						name = "AcceptanceTestEA"
						type = "STRING"
						comment = "Acceptance test updated extensible attribute"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_ea_definition.ea_def", "name", `AcceptanceTestEA`),
					resource.TestCheckResourceAttr("infoblox_ea_definition.ea_def", "type", `STRING`),
					resource.TestCheckResourceAttr("infoblox_ea_definition.ea_def", "comment",
						`Acceptance test updated extensible attribute`),
				),
			},
		},
	})
}

func testAccCheckEADefinitionDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(ibclient.IBConnector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ea_definition" {
			continue
		}

		eaDef := &ibclient.EADefinition{}
		eaDefResult := &ibclient.EADefinition{}

		err := conn.GetObject(eaDef, rs.Primary.ID, nil, eaDefResult)
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}
		if eaDefResult != nil {
			return fmt.Errorf("object with ID '%s' remains", rs.Primary.ID)
		}
	}

	return nil
}
