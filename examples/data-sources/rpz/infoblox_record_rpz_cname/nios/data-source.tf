// Retrieve a specific RPZ CNAME record by filters
data "infoblox_record_rpz_cname" "get_record_rpz_cname_using_filters" {
  filters = {
    name = "blocked.rpz.example.com"
  }
}

// Retrieve specific RPZ CNAME records using Extensible Attributes
data "infoblox_record_rpz_cname" "get_record_rpz_cname_using_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME records
data "infoblox_record_rpz_cname" "get_all_rpz_cname_records" {}
