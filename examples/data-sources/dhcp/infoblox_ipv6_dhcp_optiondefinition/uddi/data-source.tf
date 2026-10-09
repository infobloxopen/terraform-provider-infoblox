// Get IPv6 DHCP Option Definition filtered by an attribute
data "infoblox_ipv6_dhcp_optiondefinition" "example_by_name" {
  filters = {
    name = "example-definition"
  }
}

// Get all IPv6 DHCP Option Definitions
data "infoblox_ipv6_dhcp_optiondefinition" "example_all" {}
