// Retrieve MX records filtered by an attribute
data "infoblox_record_mx" "get_record_mx_using_filters" {
  filters = {
    absolute_name_spec = "abc.example.com"
  }
}

// Retrieve MX records filtered by tag
data "infoblox_record_mx" "get_record_mx_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all MX records
data "infoblox_record_mx" "get_all_mx_records" {}
