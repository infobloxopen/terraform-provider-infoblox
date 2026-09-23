// List specific DTC Topologies using filters
list "infoblox_dtc_topology" "list_dtc_topology_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Topology with geographic rules"
    }
  }
  limit = 10
}

// List specific DTC Topologies using Tags
list "infoblox_dtc_topology" "list_dtc_topology_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "us-east-1"
    }
  }
}

// List DTC Topologies with resource details included
list "infoblox_dtc_topology" "list_dtc_topology_with_resource" {
  provider         = infoblox
  include_resource = true
}
