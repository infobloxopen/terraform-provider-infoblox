package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testAccDataSourceFixedAddress = fmt.Sprintf(`
							resource "infoblox_ipv4_fixed_address" "fix3"{
									ipv4addr        = "17.0.0.9"
									mac = "00:0C:24:2E:8F:2A"
									options {
										name         = "dhcp-lease-time"
										value        = "43200"
										vendor_class = "DHCP"
										num          = 51
										use_option   = true
									}
							depends_on=[infoblox_ipv4_network.net3]
							}
							data "infoblox_ipv4_fixed_address" "testFixedAddress_read1" {	
							filters = {
	   	 							ipv4addr = infoblox_ipv4_fixed_address.fix3.ipv4addr
    						}
    						depends_on = [infoblox_ipv4_fixed_address.fix3]
							}
							resource "infoblox_ipv4_network" "net3" {
									cidr = "17.0.0.0/24"
							}`)

var testAccDataSourceFixedAddressEASearch = fmt.Sprintf(`
							resource "infoblox_ipv4_fixed_address" "fix4"{
  								client_identifier_prepend_zero=true
  								comment= "fixed address"
  								dhcp_client_identifier="23"
  								disable= true
  								ext_attrs = jsonencode({
    								"Site": "Blr"
  								})
 	 							match_client = "CLIENT_ID"
  								name = "fixed_address_1"
								network = "18.0.0.0/24"
  								network_view = "default"
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
    								value = "18.0.0.2"
    								vendor_class = "DHCP"
  								}
  								use_option = true
  								depends_on=[infoblox_ipv4_network.net2]
								}
								data "infoblox_ipv4_fixed_address" "testFixedAddress_read1" {
  								filters = {
    								"*Site" = "Blr"
 	 							}
  								depends_on = [infoblox_ipv4_fixed_address.fix4]
								}
								output "fa_rec_temp1" {
  									value = data.infoblox_ipv4_fixed_address.testFixedAddress_read1
								}
								resource "infoblox_ipv4_network" "net2" {
  									cidr = "18.0.0.0/24"
								}
`)

// minimal parameters
func TestAccDataSourceFixedAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFixedAddress,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.ipv4addr", "17.0.0.9"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.mac", "00:0c:24:2e:8f:2a"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.use_option", "true")),
			},
		},
	})
}
func TestAccDataSourceFixedAddressSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceFixedAddressEASearch,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.ipv4addr", "18.0.0.1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.0.use_option", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.1.name", "routers"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.1.value", "18.0.0.2"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.1.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.1.num", "3"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.options.1.use_option", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.comment", "fixed address"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.match_client", "CLIENT_ID"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.dhcp_client_identifier", "23"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.client_identifier_prepend_zero", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.disable", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.name", "fixed_address_1"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.network", "18.0.0.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.use_option", "true"),
					resource.TestCheckResourceAttrPair("data.infoblox_ipv4_fixed_address.testFixedAddress_read1", "results.0.ext_attrs.Site", "infoblox_ipv4_fixed_address.fix4", "ext_attrs.Site"),
				),
			},
		},
	})
}
