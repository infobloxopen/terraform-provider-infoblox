// List specific IPv6 Filter Options using filters
list "infoblox_ipv6_filteroption" "list_ipv6_filter_options_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ipv6_filter_option_1"
    }
  }
}

// List specific IPv6 Filter Options using Extensible Attributes
list "infoblox_ipv6_filteroption" "list_ipv6_filter_options_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Filter Options with resource details included
list "infoblox_ipv6_filteroption" "list_ipv6_filter_options_with_resource" {
  provider         = infoblox
  include_resource = true
}
