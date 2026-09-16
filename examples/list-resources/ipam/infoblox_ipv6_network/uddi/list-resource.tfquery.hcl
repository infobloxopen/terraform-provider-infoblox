// List specific Ipv6 Networks using filters
list "infoblox_ipv6_network" "list_ipv6network_using_filters" {
  provider = infoblox
  config {
    filters = {
      address = "2001:db8:1ef8:e4ee::"
    }
  }
  limit = 10
}

// List specific Ipv6 Networks using Tags
list "infoblox_ipv6_network" "list_ipv6network_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6 Networks with resource details included
list "infoblox_ipv6_network" "list_ipv6network_with_resource" {
  provider         = infoblox
  include_resource = true
}
