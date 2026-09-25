// Retrieve a specific Record RPZ PTR Rule by filters
data "infoblox_record_rpz_ptr" "get_record_using_filters" {
  filters = {
    ptrdname = "record1.rpz.example.com"
  }
}

// Retrieve specific Record RPZ PTR Rules using Extensible Attributes
data "infoblox_record_rpz_ptr" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Record RPZ PTR Rules
data "infoblox_record_rpz_ptr" "get_all_record_rpz_ptr" {}
