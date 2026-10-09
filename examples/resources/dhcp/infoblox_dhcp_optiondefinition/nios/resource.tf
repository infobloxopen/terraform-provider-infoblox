// Create a DHCP Option Definition in the default DHCP Option Space
resource "infoblox_dhcp_optiondefinition" "option_definition_in_default_option_space" {
  nios = {
    code = 120
    name = "example_option_definition_default_space"
    type = "string"
  }
}

// Create a DHCP Option Space (Required as Parent)
resource "infoblox_dhcp_optionspace" "option_space_with_basic_fields" {
  nios = {
    name = "example_option_space"
  }
}

// Create a DHCP Option Definition in the above created DHCP Option Space
resource "infoblox_dhcp_optiondefinition" "option_definition_with_basic_fields" {
  nios = {
    code  = 10
    name  = "example_option_definition"
    type  = "string"
    space = infoblox_dhcp_optionspace.option_space_with_basic_fields.nios.name
  }
}

// Create a DHCP Option Definition with an array data type
resource "infoblox_dhcp_optiondefinition" "option_definition_with_array_type" {
  nios = {
    code  = 12
    name  = "example_option_definition_array"
    type  = "array of ip-address"
    space = infoblox_dhcp_optionspace.option_space_with_basic_fields.nios.name
  }
}
