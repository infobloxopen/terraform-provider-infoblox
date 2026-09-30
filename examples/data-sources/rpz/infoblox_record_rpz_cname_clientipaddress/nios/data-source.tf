// Retrieve a specific RPZ CNAME Client IP Address record by filters
data "infoblox_record_rpz_cname_clientipaddress" "get_record_rpz_cname_clientipaddress_using_filters" {
  filters = {
    name = "12.0.0.1.rpz-clientip.example.com"
  }
}

// Retrieve specific RPZ CNAME Client IP Address records using Extensible Attributes
data "infoblox_record_rpz_cname_clientipaddress" "get_record_rpz_cname_clientipaddress_using_ext_attrs" {
  ext_attr_filters = {
    Site = "datacenter-1"
  }
}

// Retrieve all RPZ CNAME Client IP Address records
data "infoblox_record_rpz_cname_clientipaddress" "get_all_rpz_cname_clientipaddress_records" {}
