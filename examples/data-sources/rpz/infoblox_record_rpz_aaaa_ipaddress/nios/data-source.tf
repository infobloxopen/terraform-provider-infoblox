// Retrieve a specific RPZ AAAA IP Address record by filters
data "infoblox_record_rpz_aaaa_ipaddress" "get_rpz_aaaa_ipaddress_using_filters" {
  filters = {
    name = "2001:db8::/64.rpz.example.com"
  }
}

// Retrieve specific RPZ AAAA IP Address records using Extensible Attributes
data "infoblox_record_rpz_aaaa_ipaddress" "get_rpz_aaaa_ipaddress_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "headquarters"
  }
}

// Retrieve all RPZ AAAA IP Address records
data "infoblox_record_rpz_aaaa_ipaddress" "get_all_rpz_aaaa_ipaddress_records" {}
