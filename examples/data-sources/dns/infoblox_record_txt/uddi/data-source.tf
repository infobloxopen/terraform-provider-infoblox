// Retrieve TXT records filtered by an attribute
data "infoblox_record_txt" "get_record_txt_using_filters" {
  filters = {
    absolute_name_spec = "abc.example.com"
  }
}

// Retrieve TXT records filtered by tag
data "infoblox_record_txt" "get_record_txt_using_tag_filters" {
  tag_filters = {
    region = "eu"
  }
}

// Retrieve all TXT records
data "infoblox_record_txt" "get_all_txt_records" {}
