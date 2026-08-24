// Create an IPv6 DHCP Option Space with Basic Fields
resource "infoblox_ipv6_dhcp_optionspace" "example" {
  uddi = {
    name = "example_option_space_1"
  }
}

// Create an IPv6 DHCP Option Space with Additional Fields
resource "infoblox_ipv6_dhcp_optionspace" "example_with_options" {
  uddi = {
    name = "example_option_space_2"

    //Other Optional Fields
    comment = "IPv6 DHCP option space"
    tags = {
      Site = "location-1"
    }
  }
}
