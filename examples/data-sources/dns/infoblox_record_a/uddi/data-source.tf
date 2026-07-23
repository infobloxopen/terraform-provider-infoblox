data "infoblox_record_a" "get_a_record_using_filters" {
  filters = {
    "name_in_zone" = "record_a.example.com"
  }
}

data "infoblox_record_a" "get_a_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_record_a" "get_all_a_records" {}
