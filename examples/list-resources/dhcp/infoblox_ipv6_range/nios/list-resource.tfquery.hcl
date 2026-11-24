// List specific Ipv6 Ranges using filters
list "infoblox_ipv6_range" "list_ipv6_range_using_filters" {
  provider = infoblox
  config {
    filters = {
      start_addr = "15::10"
    }
  }
  limit = 10
}

// List specific Ipv6 Ranges using Extensible Attributes
list "infoblox_ipv6_range" "list_ipv6_range_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6 Ranges with resource details included
list "infoblox_ipv6_range" "list_ipv6_range_with_resource" {
  provider         = infoblox
  include_resource = true
}
