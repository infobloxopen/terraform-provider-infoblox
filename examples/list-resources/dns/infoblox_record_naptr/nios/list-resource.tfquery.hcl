// List specific NAPTR Records using filters
list "infoblox_record_naptr" "list_naptr_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "naptr_record.example.com"
    }
  }
}

// List specific NAPTR Records using Extensible Attributes
list "infoblox_record_naptr" "list_naptr_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List NAPTR Records with resource details included
list "infoblox_record_naptr" "list_naptr_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
