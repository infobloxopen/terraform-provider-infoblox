// List specific CAA Records using filters
list "infoblox_record_caa" "list_caa_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "caa-record.example.com"
    }
  }
  limit = 10
}

// List specific CAA Records using Extensible Attributes
list "infoblox_record_caa" "list_caa_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List CAA Records with resource details included
list "infoblox_record_caa" "list_caa_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
