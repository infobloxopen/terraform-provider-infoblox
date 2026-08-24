// List specific IPv6 DHCP Option spaces using filters
list "infoblox_ipv6_dhcp_optionspace" "list_ipv6_dhcp_optionspace_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_option_space_1"
    }
  }
  limit = 10
}

// List specific IPv6 DHCP Option spaces using Tags
list "infoblox_ipv6_dhcp_optionspace" "list_ipv6_dhcp_optionspace_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 DHCP Option spaces with resource details included
list "infoblox_ipv6_dhcp_optionspace" "list_ipv6_dhcp_optionspace_with_resource" {
  provider         = infoblox
  include_resource = true
}
