// Retrieve a specific CNAME record by filters
data "infoblox_record_cname" "get_record_using_filters" {
  filters = {
    name = "example_record.example.com"
  }
}

// Retrieve specific CNAME records using Extensible Attributes
data "infoblox_record_cname" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all CNAME records
data "infoblox_record_cname" "get_all_cname_records" {}
