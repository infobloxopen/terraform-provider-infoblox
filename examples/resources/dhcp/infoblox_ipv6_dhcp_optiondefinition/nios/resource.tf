// Create an Ipv6 DHCP Option Definition with Basic Fields
// The option definition is added to the default "DHCPv6" option space.
resource "infoblox_ipv6_dhcp_optiondefinition" "ipv6_dhcp_option_definition_with_basic_fields" {
  nios = {
    code = 250
    name = "dhcp6.example_option_definition_1"
    type = "string"
  }
}

// Create an Ipv6 DHCP Option Definition in a custom Option Space
resource "infoblox_ipv6_dhcp_optionspace" "ipv6_dhcp_option_space" {
  nios = {
    name              = "example_option_space"
    enterprise_number = 5473
  }
}

resource "infoblox_ipv6_dhcp_optiondefinition" "ipv6_dhcp_option_definition_with_additional_fields" {
  nios = {
    code = 251
    name = "example_option_definition_2"
    type = "32-bit unsigned integer"

    // Other optional fields
    space = infoblox_ipv6_dhcp_optionspace.ipv6_dhcp_option_space.nios.name
  }
}
