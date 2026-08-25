// Retrieve a specific IPAM IPv6 Network using filters
data "infoblox_ipv6_network" "example_by_attribute" {
  filters = {
    "address" = "2001:db8:1ef8:e4ee::"
    "cidr"    = "64"
  }
}

// Retrieve specific IPAM IPv6 Networks using tags
data "infoblox_ipv6_network" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM IPv6 Networks
data "infoblox_ipv6_network" "example_all" {}
