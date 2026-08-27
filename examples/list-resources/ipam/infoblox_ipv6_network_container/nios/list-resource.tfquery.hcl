// List specific IPv6 Network Containers using filters
list "infoblox_ipv6_network_container" "list_containers_using_filters" {
  provider = infoblox
  config {
    filters = {
      network = "10::/64"
    }
  }
  limit = 10
}

// List specific IPv6 Network Containers using Extensible Attributes
list "infoblox_ipv6_network_container" "list_containers_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Network Containers with resource details included
list "infoblox_ipv6_network_container" "list_containers_with_resource" {
  provider         = infoblox
  include_resource = true
}
