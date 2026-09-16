// Retrieve a specific RPZ A Record by filters
data "infoblox_record_rpz_a" "get_record_using_filters" {
  filters = {
    name = "a-record.rpz.example.com"
  }
}

// Retrieve specific RPZ A Records using Extensible Attributes
data "infoblox_record_rpz_a" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ A Records
data "infoblox_record_rpz_a" "get_all_rpz_a_records" {}
