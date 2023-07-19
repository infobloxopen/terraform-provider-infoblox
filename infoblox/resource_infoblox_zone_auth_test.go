package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceZoneAuthBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCNAMERecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone" {
						fqdn = "test.com"
						comment = "Zone Auth created by terraform acceptance test"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "comment", "Zone Auth created by terraform acceptance test"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "fqdn", "test.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "view", "default"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "ext_attrs", "{\"Location\":\"AcceptanceTerraform\"}"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_zone_auth" "test_zone" {
						fqdn = "test.com"
						comment = "Comment updated by terraform acceptance test"
						ext_attrs = jsonencode({
							Location = "AcceptanceTerraform"
						})
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "fqdn", "test.com"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "comment", "Comment updated by terraform acceptance test"),
					resource.TestCheckResourceAttr("infoblox_zone_auth.test_zone", "ext_attrs", "{\"Location\":\"AcceptanceTerraform\"}"),
				),
			},
		},
	},
	)
}
