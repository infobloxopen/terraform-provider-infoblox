// Retrieve a specific Shared SRV Record by filters
data "infoblox_sharedrecord_srv" "get_sharedrecord_srv_with_filter" {
  filters = {
    name = "sharedrecord_srv.example.com"
  }
}

// Retrieve specific Shared SRV Records using Extensible Attributes
data "infoblox_sharedrecord_srv" "get_sharedrecord_srv_with_extattr_filter" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared SRV Records
data "infoblox_sharedrecord_srv" "get_all_srv_sharedrecords" {}
