// Retrieve a specific IPv6 Filter Option by filters
data "infoblox_ipv6_filteroption" "get_ipv6_filter_option_using_filters" {
  filters = {
    name = "ipv6_filter_option_example"
  }
}

// Retrieve specific IPv6 Filter Options using Tags
data "infoblox_ipv6_filteroption" "get_ipv6_filter_option_using_tag_filters" {
  tag_filters = {
    location = "site1"
  }
}

// Retrieve all IPv6 Filter Options
data "infoblox_ipv6_filteroption" "get_all_ipv6_filter_options" {}
