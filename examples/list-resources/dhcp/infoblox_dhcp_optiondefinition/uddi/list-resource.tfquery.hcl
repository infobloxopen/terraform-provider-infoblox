// List specific Dhcp Optiondefinitions using filters
list "infoblox_dhcp_optiondefinition" "list_dhcp_optiondefinition_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-code"
    }
  }
  limit = 10
}

// List specific Dhcp Optiondefinitions using Tags
list "infoblox_dhcp_optiondefinition" "list_dhcp_optiondefinition_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dhcp Optiondefinitions with resource details included
list "infoblox_dhcp_optiondefinition" "list_dhcp_optiondefinition_with_resource" {
  provider         = infoblox
  include_resource = true
}
