// Get IPV6 DHCP Option Definition filtered by an attribute
data "infoblox_ipv6_dhcp_optiondefinition" "example_by_name" {
  filters = {
    name = "example-definition"
  }
}

// Get all DHCP Option Definition
data "infoblox_ipv6_dhcp_optiondefinition" "example_all" {}
