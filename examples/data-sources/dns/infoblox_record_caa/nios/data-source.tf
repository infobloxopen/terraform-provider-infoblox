// Retrieve a specific CAA record using filters
data "infoblox_record_caa" "get_caa_record_using_filters" {
  filters = {
    name = "caa-record.example.com"
  }
}

// Retrieve specific CAA records using Extensible Attributes
data "infoblox_record_caa" "get_caa_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all CAA records
data "infoblox_record_caa" "get_all_caa_records" {}
