// Retrieve a specific IPv6 DHCP Option Space by filters
data "infoblox_ipv6_dhcp_optionspace" "get_ipv6_option_space_by_name" {
  filters = {
    name = "example_option_space_1"
  }
}

// Retrieve a specific IPv6 DHCP Option Space by tags
data "infoblox_ipv6_dhcp_optionspace" "get_option_space_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 DHCP Option Spaces
data "infoblox_ipv6_dhcp_optionspace" "get_all_option_spaces" {}
