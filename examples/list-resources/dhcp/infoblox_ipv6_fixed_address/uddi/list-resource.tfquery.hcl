// List specific IPv6 Fixed Addresses using filters
list "infoblox_ipv6_fixed_address" "list_ipv6_fixed_address_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_fixed_address"
    }
  }
  limit = 10
}

// List specific IPv6 Fixed Addresses using Tags
list "infoblox_ipv6_fixed_address" "list_ipv6_fixed_address_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Fixed Addresses with resource details included
list "infoblox_ipv6_fixed_address" "list_ipv6_fixed_address_with_resource" {
  provider         = infoblox
  include_resource = true
}
