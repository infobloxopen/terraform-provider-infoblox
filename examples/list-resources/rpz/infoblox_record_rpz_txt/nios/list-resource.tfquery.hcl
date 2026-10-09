// List specific RPZ TXT Records using filters
list "infoblox_record_rpz_txt" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "blocked.rpz.example.com"
    }
  }
}

// List specific RPZ TXT Records using Extensible Attributes
list "infoblox_record_rpz_txt" "list_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List RPZ TXT Records with resource details included
list "infoblox_record_rpz_txt" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
