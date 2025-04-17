package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

var testAccDataSourceRange = fmt.Sprintf(`
								resource "infoblox_ipv4_range" "range" {
  									start_addr = "17.0.0.221"
  									end_addr   = "17.0.0.240"
  									options {
    									name         = "dhcp-lease-time"
    									value        = "43200"
    									vendor_class = "DHCP"
    									num          = 51
    									use_option   = true
  									}
									options {
    									name = "routers"
    									num = "3"
    									use_option = true
    									value = "17.0.0.2"
    									vendor_class = "DHCP"
  								}
  									network              = "17.0.0.0/24"
  									network_view = "default"
  									comment              = "test comment"
  									name                 = "test_range"
  									disable              = false
  									member = jsonencode({
    									name = "infoblox.localdomain"
		})
  									server_association_type= "MEMBER"
 									 ext_attrs = jsonencode({
    									"Site" = "Blr"
  									})
  									use_options = true
}

								data "infoblox_ipv4_range" "range_rec_temp" {
  								filters = {
    								start_addr = infoblox_ipv4_range.range.start_addr
  								}
  								depends_on = [infoblox_ipv4_range.range]
								}
								`)

// minimal parameters
func TestAccDataSourceRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRange,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.start_addr", "17.0.0.221"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.0.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.0.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.0.use_option", "true"),
					resource.TestCheckResourceAttrPair("data.infoblox_ipv4_range.range_rec_temp", "results.0.ext_attrs.Site", "infoblox_ipv4_range.range", "ext_attrs.Site"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.comment", "test comment"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.name", "test_range"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.disable", "false"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.server_association_type", "MEMBER"),
					resource.TestMatchResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.member", regexp.MustCompile(`"name":"infoblox.localdomain"`)),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.use_options", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.end_addr", "17.0.0.240"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.network", "17.0.0.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.1.name", "routers"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.1.value", "17.0.0.2"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.1.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.1.num", "3"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_range.range_rec_temp", "results.0.options.1.use_option", "true")),
			},
		},
	})
}
