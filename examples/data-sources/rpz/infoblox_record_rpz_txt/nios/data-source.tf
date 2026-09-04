// Retrieve a specific RPZ TXT record by filters
data "infoblox_record_rpz_txt" "get_record_rpz_txt_using_filters" {
  filters = {
    name = "blocked.rpz.example.com"
  }
}

// Retrieve specific RPZ TXT records using Extensible Attributes
data "infoblox_record_rpz_txt" "get_record_rpz_txt_using_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ TXT records
data "infoblox_record_rpz_txt" "get_all_rpz_txt_records" {}
