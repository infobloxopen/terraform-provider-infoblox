// Retrieve a specific Shared CNAME Record by filters
data "infoblox_sharedrecord_cname" "get_shared_cname_record_using_filters" {
  filters = {
    name = "example-shared-record-cname"
  }
}

// Retrieve specific Shared CNAME Records using Extensible Attributes
data "infoblox_sharedrecord_cname" "get_shared_cname_records_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Shared CNAME records
data "infoblox_sharedrecord_cname" "get_all_shared_cname_records" {}
