data "infoblox_record_ns" "get_ns_record_using_filters" {
  filters = {
    "absolute_name_spec" = "ns.example.com."
  }
}

data "infoblox_record_ns" "get_ns_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

data "infoblox_record_ns" "get_all_ns_records" {}
