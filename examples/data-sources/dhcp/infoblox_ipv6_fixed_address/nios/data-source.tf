// Retrieve a specific IPv6 Fixed Address by filters
data "infoblox_ipv6_fixed_address" "get_ipv6_fixed_address_using_filters" {
  filters = {
    ipv6addr = "2001:db8:abcd:1231::2"
  }
}

// Retrieve specific IPv6 Fixed Address using Extensible Attributes
data "infoblox_ipv6_fixed_address" "get_ipv6_fixed_address_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 Fixed Addresses
data "infoblox_ipv6_fixed_address" "get_all_ipv6_fixed_address" {}
