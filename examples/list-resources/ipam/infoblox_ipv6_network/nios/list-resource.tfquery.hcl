// List specific IPv6 Networks using filters
list "infoblox_ipv6_network" "list_ipv6networks_using_filters" {
  provider = infoblox
  config {
    filters = {
      network = "2001:db8::/32"
    }
  }
  limit = 10
}

// List specific IPv6 Networks using Extensible Attributes
list "infoblox_ipv6_network" "list_ipv6networks_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Networks with resource details included
list "infoblox_ipv6_network" "list_ipv6networks_with_resource" {
  provider         = infoblox
  include_resource = true
}
