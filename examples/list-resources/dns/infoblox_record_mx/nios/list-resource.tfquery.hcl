// List specific MX Records using filters
list "infoblox_record_mx" "list_mx_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "mx_record.example.com"
    }
  }
}

// List specific MX Records using Extensible Attributes
list "infoblox_record_mx" "list_mx_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List MX Records with resource details included
list "infoblox_record_mx" "list_mx_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
