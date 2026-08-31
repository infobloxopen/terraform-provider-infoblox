// List specific Dtc Topologies using filters
list "infoblox_dtc_topology" "list_dtc_topology_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Dtc Topologies using Tags
list "infoblox_dtc_topology" "list_dtc_topology_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dtc Topologies with resource details included
list "infoblox_dtc_topology" "list_dtc_topology_with_resource" {
  provider         = infoblox
  include_resource = true
}
