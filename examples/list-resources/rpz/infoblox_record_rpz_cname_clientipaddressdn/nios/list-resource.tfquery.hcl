// List specific RPZ CNAME Client IP Address DN Records using filters
list "infoblox_record_rpz_cname_clientipaddressdn" "list_by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "10.10.0.0/24.rpz-test.infoblox.com"
    }
  }
}

// List specific RPZ CNAME Client IP Address DN Records using Extensible Attributes
list "infoblox_record_rpz_cname_clientipaddressdn" "list_by_ext_attr_filters" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "datacenter-1"
    }
  }
}

// List RPZ CNAME Client IP Address DN Records with resource details included
list "infoblox_record_rpz_cname_clientipaddressdn" "list_with_resource" {
  provider         = infoblox
  include_resource = true
}
