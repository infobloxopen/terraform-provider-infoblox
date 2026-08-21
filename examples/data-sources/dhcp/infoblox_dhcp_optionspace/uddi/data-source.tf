// Get DHCP Option spaces filtered by an attribute
data "infoblox_dhcp_optionspace" "example_by_name" {
  filters = {
    name = "example-space"
  }
}

// Get DHCP Option space/s by tag
data "infoblox_dhcp_optionspace" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all DHCP Option spaces
data "infoblox_dhcp_optionspace" "example_all" {}
