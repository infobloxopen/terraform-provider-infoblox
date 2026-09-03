// Retrieve a specific Shared A Record by filters
data "infoblox_sharedrecord_a" "get_sharedrecord_a_with_filter" {
  filters = {
    name = "sharedrecord_a_basic"
  }
}

// Retrieve specific Shared A Records using Extensible Attributes
data "infoblox_sharedrecord_a" "get_sharedrecord_a_with_extattr_filter" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared A Records
data "infoblox_sharedrecord_a" "get_all_shared_a_records" {}
