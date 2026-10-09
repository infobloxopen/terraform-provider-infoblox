// Retrieve a specific NAPTR record by filters
data "infoblox_record_naptr" "get_record_using_filters" {
  filters = {
    name = "naptr_record.example.com"
  }
}

// Retrieve specific NAPTR records using Extensible Attributes
data "infoblox_record_naptr" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all NAPTR records
data "infoblox_record_naptr" "get_all_naptr_records" {}
