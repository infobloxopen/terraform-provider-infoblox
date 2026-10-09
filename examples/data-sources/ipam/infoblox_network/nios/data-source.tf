// Retrieve a specific IPAM network container using filters
data "infoblox_network" "get_network_using_filters" {
  filters = {
    "network" = "10.0.0.0/24"
  }
}

// Retrieve specific IPAM network using Extensible Attributes
data "infoblox_network" "get_network_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM network
data "infoblox_network" "get_all_network" {}
