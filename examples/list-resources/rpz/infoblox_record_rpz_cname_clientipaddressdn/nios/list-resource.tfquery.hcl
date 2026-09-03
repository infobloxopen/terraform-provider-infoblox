// List specific RPZ CNAME Client IP Address DN Records using filters
list "infoblox_record_rpz_cname_clientipaddressdn" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "10.10.0.0/24.rpz.example.com"
    }
  }
}

// List specific RPZ CNAME Client IP Address DN Records using Extensible Attributes
list "infoblox_record_rpz_cname_clientipaddressdn" "list_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "datacenter-1"
    }
  }
}

// List RPZ CNAME Client IP Address DN Records with resource details included
list "infoblox_record_rpz_cname_clientipaddressdn" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
