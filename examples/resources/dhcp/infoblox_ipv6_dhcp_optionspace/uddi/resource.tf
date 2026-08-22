resource "infoblox_ipv6_dhcp_optionspace" "example" {
  uddi = {
    name     = "example_dhcp_option_space"
    protocol = "ip4"
  }
}

resource "infoblox_ipv6_dhcp_optionspace" "example_with_options" {
  uddi = {
    name     = "example_dhcp_option_space_with_options"
    protocol = "ip6"
    //Other Optional Fields
    comment = "dhcp option space"
    tags = {
      Site = "location-1"
    }
  }
}
