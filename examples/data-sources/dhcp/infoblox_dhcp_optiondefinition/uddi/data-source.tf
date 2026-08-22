// Retrieve a specific DHCP Option Definitions by name
data "infoblox_dhcp_optiondefinition" "example_by_name" {
  filters = {
    name = "example_option_definition"
  }
}

// Retrieve all DHCP Option Definitions
data "infoblox_dhcp_optiondefinition" "example_all" {}
