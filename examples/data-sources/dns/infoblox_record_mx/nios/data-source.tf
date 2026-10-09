// Retrieve a specific MX record by filters
data "infoblox_record_mx" "get_record_using_filters" {
  filters = {
    name = "mx_record.example.com"
  }
}

// Retrieve specific MX records using Extensible Attributes
data "infoblox_record_mx" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all MX records
data "infoblox_record_mx" "get_all_mx_records" {}
