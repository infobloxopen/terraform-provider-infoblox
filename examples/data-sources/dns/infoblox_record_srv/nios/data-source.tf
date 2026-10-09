// Retrieve a specific SRV record by filters
data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    name = "example-srv-record.example-auth-zone.com"
  }
}

// Retrieve specific SRV records using Extensible Attributes
data "infoblox_record_srv" "get_srv_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all SRV records
data "infoblox_record_srv" "get_all_srv_records" {}
