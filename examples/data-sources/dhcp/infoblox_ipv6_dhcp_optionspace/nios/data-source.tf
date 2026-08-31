// Retrieve a specific IPv6 DHCP Option Space by filters
data "infoblox_ipv6_dhcp_optionspace" "get_ipv6_dhcp_option_space_using_filters" {
  filters = {
    name = "example_option_space_1"
  }
}

// Retrieve all IPv6 DHCP Option Spaces
data "infoblox_ipv6_dhcp_optionspace" "get_all_ipv6_dhcp_option_spaces" {}
