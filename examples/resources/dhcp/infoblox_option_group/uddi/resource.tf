resource "infoblox_option_group" "example" {
  uddi = {
    name     = "example_dhcp_option_group"
    protocol = "ip4"
  }
}

resource "infoblox_dhcp_optionspace" "option_space" {
  uddi = {
    name = "option_space"
  }
}

resource "infoblox_dhcp_optiondefinition" "option_code" {
  uddi = {
    code         = 234
    name         = "option_code"
    option_space = infoblox_dhcp_optionspace.option_space.id
    type         = "boolean"
  }
}

resource "infoblox_option_group" "example_with_options" {
  uddi = {
    name     = "example_dhcp_option_group_with_options"
    protocol = "ip4"

    dhcp_options = [
      {
        type         = "option"
        option_code  = infoblox_dhcp_optiondefinition.option_code.id
        option_value = "true"
      }
    ]
    comment = "dhcp option group"
    tags = {
      Site = "location-1"
    }
  }
}
