// Retrieve a specific TXT record by filters
data "infoblox_record_txt" "get_record_using_filters" {
  filters = {
    name = "example-txt-record.example.com"
  }
}

// Retrieve specific TXT records using Extensible Attributes
data "infoblox_record_txt" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all TXT records
data "infoblox_record_txt" "get_all_txt_records" {}
