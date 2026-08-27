// Retrieve a specific DHCP Option Space by filters
data "infoblox_dhcp_optionspace" "get_dhcp_option_space_using_filters" {
  filters = {
    name = "example_option_space_1"
  }
}

// Retrieve all DHCP Option Spaces
data "infoblox_dhcp_optionspace" "get_all_dhcp_option_spaces" {}
