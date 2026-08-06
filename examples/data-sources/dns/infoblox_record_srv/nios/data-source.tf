data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    name = "example-srv-record.example.com"
  }
}

output "infoblox_record_srv" {
  value = data.infoblox_record_srv.get_srv_record_using_filters.results
}

data "infoblox_record_srv" "get_srv_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

output "infoblox_record_srv_ext_attr_filters" {
  value = data.infoblox_record_srv.get_srv_record_using_extensible_attributes.results
}

data "infoblox_record_srv" "get_all_srv_records" {}
