// Retrieve a specific RPZ CNAME Client IP Address DN Record by filters
data "infoblox_record_rpz_cname_clientipaddressdn" "get_by_filters" {
  filters = {
    name = "10.10.0.0/24.rpz-test.infoblox.com"
  }
}

// Retrieve specific RPZ CNAME Client IP Address DN Records using Extensible Attributes
data "infoblox_record_rpz_cname_clientipaddressdn" "get_by_ext_attr_filters" {
  ext_attr_filters = {
    Site = "datacenter-1"
  }
}

// Retrieve all RPZ CNAME Client IP Address DN Records
data "infoblox_record_rpz_cname_clientipaddressdn" "get_all" {}
