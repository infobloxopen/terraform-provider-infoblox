// List specific IPv6 Fixed Addresses using filters
list "infoblox_ipv6_fixed_address" "list_ipv6fixedaddress_using_filters" {
  provider = infoblox
  config {
    filters = {
      ipv6addr = "2001:db8:abcd:1231::2"
    }
  }
  limit = 10
}

// List specific IPv6 Fixed Addresses using Extensible Attributes
list "infoblox_ipv6_fixed_address" "list_ipv6fixedaddress_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Fixed Addresses with resource details included
list "infoblox_ipv6_fixed_address" "list_ipv6fixedaddress_with_resource" {
  provider         = infoblox
  include_resource = true
}
