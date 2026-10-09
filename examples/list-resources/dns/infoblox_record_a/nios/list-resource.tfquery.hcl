// List specific A Records using filters
list "infoblox_record_a" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}

// List specific A Records using Extensible Attributes
list "infoblox_record_a" "list_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List A Records with resource details included
list "infoblox_record_a" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
