// List specific Shared MX Records using filters
list "infoblox_sharedrecord_mx" "list_shared_mx_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "sharedrecord-mx-basic"
    }
  }
  limit = 10
}

// List specific Shared MX Records using Extensible Attributes
list "infoblox_sharedrecord_mx" "list_shared_mx_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Shared MX Records with resource details included
list "infoblox_sharedrecord_mx" "list_shared_mx_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
