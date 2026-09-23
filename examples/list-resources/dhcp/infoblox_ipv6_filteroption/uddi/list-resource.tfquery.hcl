// List specific Ipv6 Filter Options using filters
list "infoblox_ipv6_filteroption" "list_ipv6_filter_option_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "ipv6_filter_option_example"
    }
  }
  limit = 10
}

// List specific Ipv6 Filter Options using Tags
list "infoblox_ipv6_filteroption" "list_ipv6_filter_option_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      location = "site1"
    }
  }
}

// List Ipv6 Filter Options with resource details included
list "infoblox_ipv6_filteroption" "list_ipv6_filter_option_with_resource" {
  provider         = infoblox
  include_resource = true
}
