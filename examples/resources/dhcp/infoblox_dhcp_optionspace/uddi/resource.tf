// Create a DHCP Option Space with Basic Fields
resource "infoblox_dhcp_optionspace" "example" {
  uddi = {
    name = "example_option_space_1"
  }
}

// Create a DHCP Option Space with Additional Fields 
resource "infoblox_dhcp_optionspace" "example_with_options" {
  uddi = {
    name = "example_option_space_2"

    //Other Optional Fields
    comment = "DHCP Option Space"
    tags = {
      Site = "location-1"
    }
  }
}
