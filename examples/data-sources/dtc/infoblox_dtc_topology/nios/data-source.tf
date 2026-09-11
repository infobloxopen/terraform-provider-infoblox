// Retrieve a specific DTC Topology using filters
data "infoblox_dtc_topology" "get_dtc_topology_using_filters" {
  filters = {
    name = "example-topology-basic"
  }
}

// Retrieve specific DTC Topologies using Extensible Attributes
data "infoblox_dtc_topology" "get_dtc_topology_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "us-east-1"
  }
}

// Retrieve all DTC Topologies
data "infoblox_dtc_topology" "get_all_dtc_topologies" {}
