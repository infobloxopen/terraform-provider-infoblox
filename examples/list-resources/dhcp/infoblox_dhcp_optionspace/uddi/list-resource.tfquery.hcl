// List specific Dhcp Optionspaces using filters
list "infoblox_dhcp_optionspace" "list_dhcp_optionspace_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-space"
    }
  }
  limit = 10
}

// List specific Dhcp Optionspaces using Tags
list "infoblox_dhcp_optionspace" "list_dhcp_optionspace_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dhcp Optionspaces with resource details included
list "infoblox_dhcp_optionspace" "list_dhcp_optionspace_with_resource" {
  provider         = infoblox
  include_resource = true
}
