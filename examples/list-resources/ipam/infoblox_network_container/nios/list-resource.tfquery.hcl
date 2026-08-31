// List specific Network Containers using filters
list "infoblox_network_container" "list_containers_using_filters" {
  provider = infoblox
  config {
    filters = {
      network = "10.0.0.0/24"
    }
  }
  limit = 10
}

// List specific Network Containers using Extensible Attributes
list "infoblox_network_container" "list_containers_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Network Containers with resource details included
list "infoblox_network_container" "list_containers_with_resource" {
  provider         = infoblox
  include_resource = true
}
