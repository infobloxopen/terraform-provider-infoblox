// Retrieve a specific IPAM network container using filters
data "infoblox_network_container" "get_network_containers_using_filters" {
  filters = {
    "network" = "10.0.0.0/24"
  }
}

// Retrieve specific IPAM network containers using Extensible Attributes
data "infoblox_network_container" "get_network_containers_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM network containers
data "infoblox_network_container" "get_all_network_containers" {}
