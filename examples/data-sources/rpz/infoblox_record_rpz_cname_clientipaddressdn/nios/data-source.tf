// Retrieve a specific RPZ CNAME Client IP Address DN record by filters
data "infoblox_record_rpz_cname_clientipaddressdn" "get_record_rpz_cname_clientipaddressdn_using_filters" {
  filters = {
    name = "10.10.0.0/24.rpz.example.com"
  }
}

// Retrieve specific RPZ CNAME Client IP Address DN records using Extensible Attributes
data "infoblox_record_rpz_cname_clientipaddressdn" "get_record_rpz_cname_clientipaddressdn_using_ext_attrs" {
  ext_attr_filters = {
    Site = "datacenter-1"
  }
}

// Retrieve all RPZ CNAME Client IP Address DN records
data "infoblox_record_rpz_cname_clientipaddressdn" "get_all_rpz_cname_clientipaddressdn_records" {}
