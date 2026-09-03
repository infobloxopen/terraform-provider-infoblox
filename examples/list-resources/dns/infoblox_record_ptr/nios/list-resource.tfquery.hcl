// List specific PTR Records using filters
list "infoblox_record_ptr" "list_ptr_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}

// List specific PTR Records using Extensible Attributes
list "infoblox_record_ptr" "list_ptr_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List PTR Records with resource details included
list "infoblox_record_ptr" "list_ptr_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
