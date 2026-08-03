data "infoblox_record_srv" "get_srv_record_using_filters" {
  filters = {
    "name_in_zone" = "record_srv.example.com"
  }
}

data "infoblox_record_srv" "get_srv_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_record_srv" "get_all_srv_records" {}
