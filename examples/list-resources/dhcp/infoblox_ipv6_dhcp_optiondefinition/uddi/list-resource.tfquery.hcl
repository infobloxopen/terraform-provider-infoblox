// List specific IPv6 DHCP Option definitions using filters
list "infoblox_ipv6_dhcp_optiondefinition" "list_ipv6_dhcp_optiondefinition_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_option_definition"
    }
  }
  limit = 10
}

// List specific IPv6 DHCP Option definitions using Tags
list "infoblox_ipv6_dhcp_optiondefinition" "list_ipv6_dhcp_optiondefinition_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 DHCP Option definitions with resource details included
list "infoblox_ipv6_dhcp_optiondefinition" "list_ipv6_dhcp_optiondefinition_with_resource" {
  provider         = infoblox
  include_resource = true
}
