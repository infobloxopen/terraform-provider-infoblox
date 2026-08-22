// Get DHCP Option codes filtered by an attribute
data "infoblox_ipv6_dhcp_optiondefinition" "example_by_name" {
  filters = {
    name = "example-code"
  }
}

// Get DHCP Option code/s by tag
data "infoblox_ipv6_dhcp_optiondefinition" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all DHCP Option codes
data "infoblox_ipv6_dhcp_optiondefinition" "example_all" {}
