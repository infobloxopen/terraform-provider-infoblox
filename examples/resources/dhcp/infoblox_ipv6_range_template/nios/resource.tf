// Create an IPV6 DHCP Range Template with Basic Fields
resource "infoblox_ipv6_range_template" "ipv6_range_template_basic_fields" {
  nios = {
    name                = "example_range_template"
    number_of_addresses = 10
    offset              = 20
    // add cloud_api_compatible = true if Terraform Internal ID extensible attribute has cloud access
    cloud_api_compatible = false
  }
}

// Create an IPV6 DHCP Range Template with Additional Fields
resource "infoblox_ipv6_range_template" "ipv6_range_template_additional_fields" {
  nios = {
    name                = "example_range_template_additional_fields"
    number_of_addresses = 100
    offset              = 200
    // add cloud_api_compatible = true if Terraform Internal ID extensible attribute has cloud access
    cloud_api_compatible = true
    comment              = "Example comment for ipv6 range template"
    exclude = [
      {
        number_of_addresses = 10
        offset              = 20
        comment             = "Example comment for range template exclude"
      }
    ]
    logic_filter_rules = [
      {
        filter = "ipv6_option_filter"
        type   = "Option"
      }
    ]
    option_filter_rules = [
      {
        filter     = "ipv6_option_filter1"
        permission = "Deny"
      }
    ]
    recycle_leases = false
  }
}
