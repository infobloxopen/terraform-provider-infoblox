// List specific RPZ A Records using filters
list "infoblox_record_rpz_a" "list_record_rpz_a_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "a-record.rpz.example.com"
    }
  }
}

// List specific RPZ A Records using Extensible Attributes
list "infoblox_record_rpz_a" "list_record_rpz_a_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List RPZ A Records with resource details included
list "infoblox_record_rpz_a" "list_record_rpz_a_with_resource" {
  provider         = infoblox
  include_resource = true
}
