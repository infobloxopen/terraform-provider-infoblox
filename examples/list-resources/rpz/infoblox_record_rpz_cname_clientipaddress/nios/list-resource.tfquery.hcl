// List specific RPZ CNAME Client IP Address Records using filters
list "infoblox_record_rpz_cname_clientipaddress" "list_records_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "12.0.0.1.rpz-clientip.example.com"
    }
  }
}

// List specific RPZ CNAME Client IP Address Records using Extensible Attributes
list "infoblox_record_rpz_cname_clientipaddress" "list_records_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "datacenter-1"
    }
  }
}

// List RPZ CNAME Client IP Address Records with resource details included
list "infoblox_record_rpz_cname_clientipaddress" "list_records_with_resource" {
  provider         = infoblox
  include_resource = true
}
