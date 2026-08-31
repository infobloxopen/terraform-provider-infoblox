// List specific DHCP Option Definitions using filters
list "infoblox_dhcp_optiondefinition" "list_dhcp_optiondefinition_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_option_definition"
    }
  }
  limit = 10
}

// List DHCP Option Definitions with resource details included
list "infoblox_dhcp_optiondefinition" "list_dhcp_optiondefinition_with_resource" {
  provider         = infoblox
  include_resource = true
}
