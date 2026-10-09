// Retrieve a specific PTR record using filters
data "infoblox_record_ptr" "get_ptr_record_using_filters" {
  filters = {
    ipv4addr = "10.20.1.2"
  }
}

// Retrieve specific PTR records using Extensible Attributes
data "infoblox_record_ptr" "get_ptr_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all PTR records
data "infoblox_record_ptr" "get_all_ptr_records" {}
