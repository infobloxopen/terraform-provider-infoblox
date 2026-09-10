// Retrieve a specific Host Record by filters
data "infoblox_record_host" "get_host_record_using_filters" {
  filters = {
    name = "host-1.example.com"
  }
}

// Retrieve a specific Host Record by Extensible Attributes
data "infoblox_record_host" "get_host_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Host Records
data "infoblox_record_host" "get_all_host_records" {}
