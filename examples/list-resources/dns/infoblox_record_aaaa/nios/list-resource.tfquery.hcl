// List specific AAAA Records using filters
list "infoblox_record_aaaa" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_record.example.com"
    }
  }
}

// List specific AAAA Records using Extensible Attributes
list "infoblox_record_aaaa" "list_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List AAAA Records with resource details included
list "infoblox_record_aaaa" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
