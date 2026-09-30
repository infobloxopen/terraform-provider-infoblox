// Retrieve a specific DHCP Option Group by name
data "infoblox_option_group" "example_by_name" {
  filters = {
    name = "example_dhcp_option_group"
  }
}

// Retrieve DHCP Option Groups using tag filters
data "infoblox_option_group" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Option Groups
data "infoblox_option_group" "example_all" {}
