// Create an IPv6 Filter Option with Basic Fields
resource "infoblox_ipv6_filteroption" "ipv6_filter_option_with_basic_fields" {
  nios = {
    name = "example_ipv6_filter_option_1"
  }
}

// Create an IPv6 Filter Option with Additional Fields
resource "infoblox_ipv6_filteroption" "ipv6_filter_option_with_additional_fields" {
  nios = {
    name = "example_ipv6_filter_option_2"

    // Additional Fields
    comment    = "IPv6 Filter Option created via Terraform"
    expression = "(option dhcp6.server-id=\"server-id\" AND option dhcp6.vendor-class=\"DHCPv6\")"
    lease_time = 7200
    option_list = [
      {
        name         = "dhcp6.name-servers"
        num          = 23
        value        = "fc00::,2001:db8::"
        vendor_class = "DHCPv6"
      },
      {
        name         = "dhcp6.remote-id"
        num          = 37
        value        = "remote-id"
        vendor_class = "DHCPv6"
      }
    ]

    //Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
