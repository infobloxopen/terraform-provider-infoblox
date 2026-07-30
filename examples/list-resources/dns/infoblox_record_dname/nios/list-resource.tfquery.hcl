// List specific DNAME Records using filters
list "infoblox_record_dname" "list_dname_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      view = "default"
    }
  }
  limit = 10
}

// List specific DNAME Records using Extensible Attributes
list "infoblox_record_dname" "list_dname_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DNAME Records with resource details included
list "infoblox_record_dname" "list_dname_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
