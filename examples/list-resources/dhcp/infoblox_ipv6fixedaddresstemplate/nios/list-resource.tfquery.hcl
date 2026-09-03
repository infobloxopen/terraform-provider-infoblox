// List specific IPv6 Fixed Address Templates using filters
list "infoblox_ipv6fixedaddresstemplate" "list_ipv6_fixed_address_template_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ipv6_fixed_address_template"
    }
  }
}

// List specific IPv6 Fixed Address Templates using Extensible Attributes
list "infoblox_ipv6fixedaddresstemplate" "list_ipv6_fixed_address_template_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Fixed Address Templates with resource details included
list "infoblox_ipv6fixedaddresstemplate" "list_ipv6_fixed_address_template_with_resource" {
  provider         = infoblox
  include_resource = true
  limit            = 10
}
