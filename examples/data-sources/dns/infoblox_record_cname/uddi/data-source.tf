// Retrieve CNAME records filtered by an attribute
data "infoblox_record_cname" "get_record_cname_using_filters" {
  filters = {
    absolute_name_spec = "abc.example.com"
  }
}

// Retrieve CNAME records filtered by tag
data "infoblox_record_cname" "get_record_cname_using_tag_filters" {
  tag_filters = {
    region = "eu"
  }
}

// Retrieve all CNAME records
data "infoblox_record_cname" "get_all_cname_records" {}
