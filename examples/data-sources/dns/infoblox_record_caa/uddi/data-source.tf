// Retrieve a specific CAA record using filters
data "infoblox_record_caa" "get_caa_record_using_filters" {
  filters = {
    "absolute_name_spec" = "example.com."
  }
}

// Retrieve specific CAA records using Tags
data "infoblox_record_caa" "get_caa_record_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all CAA records
data "infoblox_record_caa" "get_all_caa_records" {}
