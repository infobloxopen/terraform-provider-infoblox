// List specific IPv6 Shared Networks using filters
list "infoblox_ipv6_sharednetwork" "list_ipv6_sharednetwork_by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ipv6_shared_network1"
    }
  }
  limit = 10
}

// List specific IPv6 Shared Networks using Extensible Attributes
list "infoblox_ipv6_sharednetwork" "list_ipv6_sharednetwork_by_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Shared Networks with resource details included
list "infoblox_ipv6_sharednetwork" "list_ipv6_sharednetwork_with_resource" {
  provider         = infoblox
  include_resource = true
}
