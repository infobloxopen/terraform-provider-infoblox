// List specific Shared AAAA Records using filters
list "infoblox_sharedrecord_aaaa" "list_shared_aaaa_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "sharedrecord_aaaa_basic"
    }
  }
  limit = 10
}

// List specific Shared AAAA Records using Extensible Attributes
list "infoblox_sharedrecord_aaaa" "list_shared_aaaa_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared AAAA Records with resource details included
list "infoblox_sharedrecord_aaaa" "list_shared_aaaa_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
