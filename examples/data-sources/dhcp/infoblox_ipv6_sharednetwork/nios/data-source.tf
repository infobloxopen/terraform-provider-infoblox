// Retrieve a specific IPv6 Shared Network using filters
data "infoblox_ipv6_sharednetwork" "example_ipv6_sharednetwork_by_filters" {
  filters = {
    name = "example_ipv6_shared_network1"
  }
}

// Retrieve a specific IPv6 Shared Network using Extensible Attributes
data "infoblox_ipv6_sharednetwork" "example_ipv6_sharednetwork_by_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 Shared Networks
data "infoblox_ipv6_sharednetwork" "example_ipv6_sharednetworks" {}
