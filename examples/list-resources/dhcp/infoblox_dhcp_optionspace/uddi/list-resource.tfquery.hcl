// List specific DHCP Option Spaces using filters
list "infoblox_dhcp_optionspace" "list_dhcp_optionspace_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_option_space_1"
    }
  }
  limit = 10
}

// List specific DHCP Option Spaces using Tags
list "infoblox_dhcp_optionspace" "list_dhcp_optionspace_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List DHCP Option Spaces with resource details included
list "infoblox_dhcp_optionspace" "list_dhcp_optionspace_with_resource" {
  provider         = infoblox
  include_resource = true
}
