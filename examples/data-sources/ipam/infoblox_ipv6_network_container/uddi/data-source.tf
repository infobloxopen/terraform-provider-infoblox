// Get IPv6 Network Container by Name
data "infoblox_ipv6_network_container" "example_by_attribute" {
  filters = {
    "name" = "example_ipv6_network_container"
  }
}

// Get IPv6 Network Container by Address
data "infoblox_ipv6_network_container" "example_by_address" {
  filters = {
    "address" = "2001:db8::"
    "cidr"    = "64"
  }
}

// Get IPv6 Network Container by tags
data "infoblox_ipv6_network_container" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all IPv6 Network Containers
data "infoblox_ipv6_network_container" "example_all" {}
