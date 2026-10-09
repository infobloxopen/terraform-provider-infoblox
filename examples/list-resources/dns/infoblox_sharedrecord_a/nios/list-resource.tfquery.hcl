// List specific Shared A Records using filters
list "infoblox_sharedrecord_a" "list_shared_a_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "sharedrecord_a_basic"
    }
  }
  limit = 10
}

// List specific Shared A Records using Extensible Attributes
list "infoblox_sharedrecord_a" "list_shared_a_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared A Records with resource details included
list "infoblox_sharedrecord_a" "list_shared_a_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
