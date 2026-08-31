// Create an IPv6 DHCP Option Definition with Basic Fields in default "DHCPv6" option space.
resource "infoblox_ipv6_dhcp_optiondefinition" "ipv6_dhcp_option_definition_with_basic_fields" {
  nios = {
    code = 250
    name = "dhcp6.example_option_definition_1"
    type = "string"
  }
}

// Create an IPv6 Option Space(required as parent)
resource "infoblox_ipv6_dhcp_optionspace" "ipv6_dhcp_option_space" {
  nios = {
    name              = "example_option_space"
    enterprise_number = 5473
  }
}

// Create an IPv6 DHCP Option Definition in the above created Option Space
resource "infoblox_ipv6_dhcp_optiondefinition" "ipv6_dhcp_option_definition_with_additional_fields" {
  nios = {
    code = 251
    name = "example_option_definition_2"
    type = "32-bit unsigned integer"

    space = infoblox_ipv6_dhcp_optionspace.ipv6_dhcp_option_space.nios.name
  }
}
