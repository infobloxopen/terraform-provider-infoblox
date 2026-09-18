// List specific Host Records using filters
list "infoblox_record_host" "list_host_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "host-1.example.com"
    }
  }
  limit = 10
}

// List specific Host Records using Extensible Attributes
list "infoblox_record_host" "list_host_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List Host Records with resource details included
list "infoblox_record_host" "list_host_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
