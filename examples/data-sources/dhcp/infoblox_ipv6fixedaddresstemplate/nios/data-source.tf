// Retrieve a specific IPv6 Fixed Address Template by filters
data "infoblox_ipv6fixedaddresstemplate" "get_ipv6_fixed_address_template_using_filters" {
  filters = {
    name = "example_ipv6_fixed_address_template"
  }
}

// Retrieve specific IPv6 Fixed Address Templates using Extensible Attributes
data "infoblox_ipv6fixedaddresstemplate" "get_ipv6_fixed_address_template_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 Fixed Address Templates
data "infoblox_ipv6fixedaddresstemplate" "get_all_ipv6_fixed_address_templates" {}
