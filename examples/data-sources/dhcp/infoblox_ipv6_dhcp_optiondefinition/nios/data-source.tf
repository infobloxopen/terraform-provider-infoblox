// Retrieve a specific IPv6 DHCP Option Definition by filters
data "infoblox_ipv6_dhcp_optiondefinition" "get_ipv6_dhcp_option_definition_using_filters" {
  filters = {
    name = "dhcp6.example_option_definition_1"
  }
}

// Retrieve all IPv6 DHCP Option Definitions
data "infoblox_ipv6_dhcp_optiondefinition" "get_all_ipv6_dhcp_option_definitions" {}
