// List specific DTC Topologies using filters
list "infoblox_dtc_topology" "list_dtc_topology_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-topology-basic"
    }
  }
  limit = 10
}

// List specific DTC Topologies using Extensible Attributes
list "infoblox_dtc_topology" "list_dtc_topology_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-east-1"
    }
  }
}

// List DTC Topologies with resource details included
list "infoblox_dtc_topology" "list_dtc_topology_with_resource" {
  provider         = infoblox
  include_resource = true
}
