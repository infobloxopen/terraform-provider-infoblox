// Retrieve a specific AAAA record by filters
data "infoblox_record_aaaa" "get_record_aaaa_using_filters" {
  filters = {
    name = "example_record.example.com"
  }
}

// Retrieve specific AAAA records using Extensible Attributes
data "infoblox_record_aaaa" "get_record_aaaa_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all AAAA records
data "infoblox_record_aaaa" "get_all_aaaa_records" {}
