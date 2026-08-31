// List specific Alias Records using filters
list "infoblox_record_alias" "list_alias_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "alias-record.example.com"
    }
  }
}

// List specific Alias Records using Extensible Attributes
list "infoblox_record_alias" "list_alias_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Alias Records with resource details included
list "infoblox_record_alias" "list_alias_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
