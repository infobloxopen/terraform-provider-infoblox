// Retrieve a specific Shared MX Record by filters
data "infoblox_sharedrecord_mx" "get_sharedrecord_mx_with_filter" {
  filters = {
    name = "sharedrecord-mx-basic"
  }
}

// Retrieve specific Shared MX Records using Extensible Attributes
data "infoblox_sharedrecord_mx" "get_sharedrecord_mx_with_extattr_filter" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared MX Records
data "infoblox_sharedrecord_mx" "get_all_mx_sharedrecords" {}
