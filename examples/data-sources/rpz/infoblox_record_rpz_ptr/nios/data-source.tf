// Retrieve a specific Substitute (PTR Record) Rule by filters
data "infoblox_record_rpz_ptr" "get_record_using_filters" {
  filters = {
    ptrdname = "record1.rpz.example.com"
  }
}

// Retrieve specific Substitute (PTR Record) Rules using Extensible Attributes
data "infoblox_record_rpz_ptr" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Substitute (PTR Record) Rules
data "infoblox_record_rpz_ptr" "get_all_record_rpz_ptr" {}
