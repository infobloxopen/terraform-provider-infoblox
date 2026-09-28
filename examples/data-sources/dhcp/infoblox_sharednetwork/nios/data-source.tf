// Retrieve a specific Shared Network record by filters
data "infoblox_sharednetwork" "get_shared_network_using_filters" {
  filters = {
    name = "example_shared_network1"
  }
}

// Retrieve specific Shared Networks using Extensible Attributes
data "infoblox_sharednetwork" "get_shared_network_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared Networks
data "infoblox_sharednetwork" "get_all_shared_networks" {}
