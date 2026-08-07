data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    name = "example-srv-record.example.com"
  }
}

data "infoblox_record_srv" "get_srv_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

data "infoblox_record_srv" "get_all_srv_records" {}
