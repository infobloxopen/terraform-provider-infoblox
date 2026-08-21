resource "infoblox_dhcp_optiondefinition" "option_code" {
  uddi = {
    code = 250
    name = "example_option_code"
    type = "int32"
  }
}

resource "infoblox_dhcp_optiondefinition" "option_code_with_options" {
  uddi = {
    code = 251
    name = "example_option_code_with_options"
    type = "int32"

    // Other optional fields
    array   = true
    comment = "Option code example"
  }
}
