// Retrieve a specific Substitute (NAPTR Record) Rule by filters
data "infoblox_record_rpz_naptr" "get_record_using_filters" {
  filters = {
    name = "naptr.rpz.example.com"
  }
}

// Retrieve specific Substitute (NAPTR Record) Rules using Extensible Attributes
data "infoblox_record_rpz_naptr" "get_record_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all Substitute (NAPTR Record) Rules
data "infoblox_record_rpz_naptr" "get_all_record_rpz_naptr" {}
