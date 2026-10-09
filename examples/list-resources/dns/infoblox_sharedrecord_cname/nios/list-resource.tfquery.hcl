// List specific Shared CNAME Records using filters
list "infoblox_sharedrecord_cname" "list_shared_cname_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-shared-record-cname"
    }
  }
  limit = 10
}

// List specific Shared CNAME Records using Extensible Attributes
list "infoblox_sharedrecord_cname" "list_shared_cname_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared CNAME Records with resource details included
list "infoblox_sharedrecord_cname" "list_shared_cname_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
