// Create a DHCP Option Space (Required as Parent)
resource "infoblox_dhcp_optionspace" "option_space" {
  uddi = {
    name = "example_option_space"
  }
}

// Create a DHCP Option Definition with Basic Fields
resource "infoblox_dhcp_optiondefinition" "option_definition_with_basic_fields" {
  uddi = {
    code         = 250
    name         = "example_option_definition"
    type         = "int32"
    option_space = infoblox_dhcp_optionspace.option_space.id
  }
}

// Create a DHCP Option Definition with Additional Fields
resource "infoblox_dhcp_optiondefinition" "option_definition_with_additional_fields" {
  uddi = {
    code         = 251
    name         = "example_option_definition_with_options"
    type         = "int32"
    option_space = infoblox_dhcp_optionspace.option_space.id

    // Other optional fields
    array   = true
    comment = "Option Definition Example"
  }
}
