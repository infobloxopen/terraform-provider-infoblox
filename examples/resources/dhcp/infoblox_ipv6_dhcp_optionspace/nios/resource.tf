// Create an IPv6 DHCP Option Space with Basic Fields
resource "infoblox_ipv6_dhcp_optionspace" "ipv6_dhcp_option_space_with_basic_fields" {
  nios = {
    name              = "example_option_space_1"
    enterprise_number = 5473
  }
}

// Create an IPv6 DHCP Option Space with Additional Fields
resource "infoblox_ipv6_dhcp_optionspace" "ipv6_dhcp_option_space_with_additional_fields" {
  nios = {
    name              = "example_option_space_2"
    enterprise_number = 5473
    comment           = "Example Ipv6 Option Space"
  }
}
