// Create a DHCP IPv6 Fixed Address Template with Basic Fields
resource "infoblox_ipv6_fixed_address_template" "ipv6_fixed_address_template_basic_fields" {
  nios = {
    name = "example_ipv6_fixed_address_template"
  }
}

// Create a DHCP IPv6 Fixed Address Template with Additional Fields
resource "infoblox_ipv6_fixed_address_template" "ipv6_fixed_address_template_additional_fields" {
  nios = {
    name = "example_ipv6_fixed_address_template_2"
    // Additional Fields
    comment     = "IPv6 Fixed Address Template Created by Terraform"
    domain_name = "example.com"
    options = [
      {
        name  = "dhcp-lease-time"
        num   = "51"
        value = "5000"
      }
    ]
    valid_lifetime = 5000
    ext_attrs = {
      Site = "location-1"
    }
  }
}
