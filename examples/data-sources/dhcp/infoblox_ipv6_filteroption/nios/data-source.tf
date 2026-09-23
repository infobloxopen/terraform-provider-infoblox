// Retrieve a specific IPv6 Filter Option by filters
data "infoblox_ipv6_filteroption" "get_ipv6_filter_options_using_filters" {
  filters = {
    name = "example_ipv6_filter_option_1"
  }
}

// Retrieve specific IPv6 Filter Options using Extensible Attributes
data "infoblox_ipv6_filteroption" "get_ipv6_filter_options_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 Filter Options
data "infoblox_ipv6_filteroption" "get_all_ipv6_filter_options" {}
