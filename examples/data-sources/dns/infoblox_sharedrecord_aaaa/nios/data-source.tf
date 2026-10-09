// Retrieve a specific Shared AAAA Record by filters
data "infoblox_sharedrecord_aaaa" "get_sharedrecord_aaaa_with_filter" {
  filters = {
    name = "sharedrecord_aaaa_basic"
  }
}

// Retrieve specific Shared AAAA Records using Extensible Attributes
data "infoblox_sharedrecord_aaaa" "get_sharedrecord_aaaa_with_extattr_filter" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared AAAA Records
data "infoblox_sharedrecord_aaaa" "get_all_shared_aaaa_records" {}
