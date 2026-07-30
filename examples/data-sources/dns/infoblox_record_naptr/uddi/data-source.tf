// Retrieve NAPTR records filtered by an attribute
data "infoblox_record_naptr" "get_record_naptr_using_filters" {
  filters = {
    absolute_name_spec = "abc.example.com"
  }
}

// Retrieve NAPTR records filtered by tag
data "infoblox_record_naptr" "get_record_naptr_using_tag_filters" {
  tag_filters = {
    region = "eu"
  }
}

// Retrieve all NAPTR records
data "infoblox_record_naptr" "get_all_naptr_records" {}
