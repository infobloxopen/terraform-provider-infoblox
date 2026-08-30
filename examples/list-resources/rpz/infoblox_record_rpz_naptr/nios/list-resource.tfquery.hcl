// List specific Substitute (NAPTR Record) Rules using filters
list "infoblox_record_rpz_naptr" "list_record_rpz_naptr_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "naptr.rpz-zone.example.com"
    }
  }
}

// List specific Substitute (NAPTR Record) Rules using Extensible Attributes
list "infoblox_record_rpz_naptr" "list_record_rpz_naptr_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Substitute (NAPTR Record) Rules with resource details included
list "infoblox_record_rpz_naptr" "list_record_rpz_naptr_with_resource" {
  provider         = infoblox
  include_resource = true
}
