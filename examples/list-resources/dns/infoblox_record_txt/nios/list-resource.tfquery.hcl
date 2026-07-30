// List specific TXT Records using filters
list "infoblox_record_txt" "list_txt_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "record-txt.example.com"
    }
  }
}

// List specific TXT Records using Extensible Attributes
list "infoblox_record_txt" "list_txt_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List TXT Records with resource details included
list "infoblox_record_txt" "list_txt_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
