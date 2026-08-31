// Retrieve a specific DHCP Option Space by filters
data "infoblox_dhcp_optionspace" "get_option_space_by_name" {
  filters = {
    name = "example_option_space_1"
  }
}

// Retrieve a specific DHCP Option Space by tags
data "infoblox_dhcp_optionspace" "get_option_space_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Option Spaces
data "infoblox_dhcp_optionspace" "get_all_option_spaces" {}
