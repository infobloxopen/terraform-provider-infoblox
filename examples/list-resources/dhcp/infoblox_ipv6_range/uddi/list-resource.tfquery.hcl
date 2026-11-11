// List specific Ipv6 Ranges using filters
list "infoblox_ipv6_range" "list_ipv6_range_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Example IPv6 Range created by the terraform provider"
    }
  }
  limit = 10
}

// List specific Ipv6 Ranges using Tags
list "infoblox_ipv6_range" "list_ipv6_range_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6 Ranges with resource details included
list "infoblox_ipv6_range" "list_ipv6_range_with_resource" {
  provider         = infoblox
  include_resource = true
}
