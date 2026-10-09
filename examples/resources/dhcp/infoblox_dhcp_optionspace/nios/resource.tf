// Create a DHCP Option Space with Basic Fields
resource "infoblox_dhcp_optionspace" "dhcp_option_space_with_basic_fields" {
  nios = {
    name = "example_option_space_1"
  }
}

// Create a DHCP Option Space with Additional Fields 
resource "infoblox_dhcp_optionspace" "dhcp_option_space_with_additional_fields" {
  nios = {
    name    = "example_option_space_2"
    comment = "Example Option Space"
  }
}
