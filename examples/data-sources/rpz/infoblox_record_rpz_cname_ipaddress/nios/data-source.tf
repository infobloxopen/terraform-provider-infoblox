// Retrieve a specific RPZ CNAME IP address record by filters
data "infoblox_record_rpz_cname_ipaddress" "by_name" {
  filters = {
    name = "11.0.0.1.rpzip.example.com"
  }
}

// Retrieve specific RPZ CNAME IP address records using Extensible Attributes
data "infoblox_record_rpz_cname_ipaddress" "by_ext_attr" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME IP address records
data "infoblox_record_rpz_cname_ipaddress" "all" {}
