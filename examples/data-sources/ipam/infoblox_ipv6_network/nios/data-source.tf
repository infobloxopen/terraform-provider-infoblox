// Retrieve a specific IPAM IPv6 Network using filters
data "infoblox_ipv6_network" "get_ipv6network_using_filters" {
  filters = {
    network = "10::/64"
  }
}

// Retrieve specific IPAM IPv6 Networks using Extensible Attributes
data "infoblox_ipv6_network" "get_ipv6network_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM IPv6 Networks
data "infoblox_ipv6_network" "get_all_ipv6networks" {}
