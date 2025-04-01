package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpv4SharedNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceIpv4SharedNetwork1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.name", "shared-network13"),
					checkNetworksContains("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.networks.0", "28.11.3.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.options.0.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.options.0.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read", "results.0.options.0.use_option", "true"),
				),
			},
		},
	})
}

var testAccDataSourceIpv4SharedNetwork1 = fmt.Sprintf(`resource "infoblox_ipv4_shared_network" "record3" {
    name = "shared-network13"
    networks = ["28.11.3.0/24"]
	options {
    	name = "dhcp-lease-time"
    	value = "43200"
		vendor_class = "DHCP"
		num = 51
		use_option = true
  	}
	depends_on = [infoblox_ipv4_network.ipv4_network3]
}

resource "infoblox_ipv4_network" "ipv4_network3" {
  cidr = "28.11.3.0/24"
}

data "infoblox_ipv4_shared_network" "shared_network_read" {
	filters = {
		network_view="default"
		name=infoblox_ipv4_shared_network.record3.name
	}
}`)

var testAccDataSourceIpv4SharedNetwork2 = fmt.Sprintf(`
	resource "infoblox_ipv4_shared_network" "record4" {
	name = "shared-network14"
	comment = "test ipv4 shared network record"
	networks = ["27.12.3.0/24","27.13.3.0/24"]
	network_view = "default"
	disable = false
	ext_attrs = jsonencode({
		"Site" = "Kyoto"
	})
	use_options = true
	options {
		name = "domain-name-servers"
		value = "11.22.33.44"
		vendor_class = "DHCP"
		num = 6
		use_option = true
	}
	options {
    	name = "dhcp-lease-time"
    	value = "43200"
		vendor_class = "DHCP"
		num = 51
		use_option = true
  	}
	depends_on = [infoblox_ipv4_network.ipv4_network11, infoblox_ipv4_network.ipv4_network12]
}

resource "infoblox_ipv4_network" "ipv4_network11" {
  cidr = "27.12.3.0/24"
}

resource "infoblox_ipv4_network" "ipv4_network12" {
  cidr = "27.13.3.0/24"
}

data "infoblox_ipv4_shared_network" "shared_network_read2" {
	filters = {
		"*Site" = "Kyoto"
	}
	depends_on = [infoblox_ipv4_shared_network.record4]
}
`)

func TestAccDataSourceIpv4SharedNetworkByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceIpv4SharedNetwork2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.network_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.name", "shared-network14"),
					checkNetworksContains("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.networks.0", "27.12.3.0/24"),
					checkNetworksContains("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.networks.1", "27.13.3.0/24"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.comment", "test ipv4 shared network record"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.disable", "false"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.use_options", "true"),
					resource.TestCheckResourceAttrPair("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.ext_attrs.Site", "infoblox_ipv4_shared_network.record4", "ext_attrs.Site"),

					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.0.name", "domain-name-servers"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.0.value", "11.22.33.44"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.0.num", "6"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.0.use_option", "true"),

					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.1.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.1.value", "43200"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.1.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.1.num", "51"),
					resource.TestCheckResourceAttr("data.infoblox_ipv4_shared_network.shared_network_read2", "results.0.options.1.use_option", "true"),
				),
			},
		},
	})
}

func checkNetworksContains(resourceName, attributeName, substring string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		networks := rs.Primary.Attributes[attributeName]
		if !strings.Contains(networks, substring) {
			return fmt.Errorf("attribute %s does not contain %s", attributeName, substring)
		}

		return nil
	}
}
