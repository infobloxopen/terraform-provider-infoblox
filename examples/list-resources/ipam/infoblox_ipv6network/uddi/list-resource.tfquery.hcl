// List specific Ipv6Networks using filters
list "infoblox_ipv6network" "list_ipv6network_using_filters" {
  provider = infoblox
  config {
    filters = {
      address = "10.0.0.0"
    }
  }
  limit = 10
}

// List specific Ipv6Networks using Tags
list "infoblox_ipv6network" "list_ipv6network_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6Networks with resource details included
list "infoblox_ipv6network" "list_ipv6network_with_resource" {
  provider         = infoblox
  include_resource = true
}
