// List specific Networks using filters
list "infoblox_network" "list_networks_using_filters" {
  provider = infoblox
  config {
    filters = {
      network = "10.0.0.0/24"
    }
  }
  limit = 10
}

// List specific Networks using Extensible Attributes
list "infoblox_network" "list_networks_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Networks with resource details included
list "infoblox_network" "list_networks_with_resource" {
  provider         = infoblox
  include_resource = true
}
