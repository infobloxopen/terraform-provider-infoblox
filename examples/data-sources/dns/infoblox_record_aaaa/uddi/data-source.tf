// Retrieve AAAA records filtered by an attribute
data "infoblox_record_aaaa" "get_record_aaaa_using_filters" {
  filters = {
    absolute_name_spec = "abc.example.com"
  }
}

// Retrieve AAAA records filtered by tag
data "infoblox_record_aaaa" "get_record_aaaa_using_tag_filters" {
  tag_filters = {
    region = "eu"
  }
}

// Retrieve all AAAA records
data "infoblox_record_aaaa" "get_all_aaaa_records" {}
