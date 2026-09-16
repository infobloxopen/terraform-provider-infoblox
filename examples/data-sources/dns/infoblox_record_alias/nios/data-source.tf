// Retrieve a specific alias record by filters
data "infoblox_record_alias" "get_alias_record_using_filters" {
  filters = {
    name = "alias-record.example.com"
  }
}

// Retrieve specific alias records using Extensible Attributes
data "infoblox_record_alias" "get_alias_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all alias records
data "infoblox_record_alias" "get_all_alias_records" {}
