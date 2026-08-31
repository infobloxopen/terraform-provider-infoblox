// List specific IPv6 DHCP Option Definitions using filters
list "infoblox_ipv6_dhcp_optiondefinition" "list_ipv6_dhcp_optiondefinition_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dhcp6.example_option_definition_1"
    }
  }
  limit = 10
}

// List IPv6 DHCP Option Definitions with resource details included
list "infoblox_ipv6_dhcp_optiondefinition" "list_ipv6_dhcp_optiondefinition_with_resource" {
  provider         = infoblox
  include_resource = true
}
