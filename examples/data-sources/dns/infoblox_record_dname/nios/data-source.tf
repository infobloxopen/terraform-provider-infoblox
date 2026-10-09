// Retrieve specific DNAME records using filters
data "infoblox_record_dname" "get_dname_record_using_filters" {
  filters = {
    name = "example.com"
  }
}

// Retrieve specific DNAME records using extensible attributes
data "infoblox_record_dname" "get_dname_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DNAME records
data "infoblox_record_dname" "get_all_dname_records" {}
