data "infoblox_record_a" "get_a_record_using_filters" {
  filters = {
    name = "rec-1.example.com"
  }
}

data "infoblox_record_a" "get_a_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

data "infoblox_record_a" "get_all_a_records" {}
