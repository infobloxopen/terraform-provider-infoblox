// List DHCP Option Groups using name filter
list "infoblox_option_group" "list_option_group_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_dhcp_option_group"
    }
  }
  limit = 10
}

// List DHCP Option Groups using tag filters
list "infoblox_option_group" "list_option_group_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List DHCP Option Groups and include the full resource in output
list "infoblox_option_group" "list_option_group_with_resource" {
  provider         = infoblox
  include_resource = true
}
