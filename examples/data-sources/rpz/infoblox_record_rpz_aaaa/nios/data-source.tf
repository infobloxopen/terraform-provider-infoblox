// Retrieve a specific RPZ AAAA record by filters
data "infoblox_record_rpz_aaaa" "get_record_rpz_aaaa_using_filters" {
  filters = {
    name = "blocked.rpz.example.com"
  }
}

// Retrieve specific RPZ AAAA records using Extensible Attributes
data "infoblox_record_rpz_aaaa" "get_record_rpz_aaaa_using_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ AAAA records
data "infoblox_record_rpz_aaaa" "get_all_rpz_aaaa_records" {}
