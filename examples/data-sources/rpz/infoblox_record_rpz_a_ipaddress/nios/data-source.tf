// Retrieve a specific RPZ A IP Address record by filters
data "infoblox_record_rpz_a_ipaddress" "get_rpz_a_ipaddress_using_filters" {
  filters = {
    name = "10.10.0.0/16.rpz.example.com"
  }
}

// Retrieve specific RPZ A IP Address records using Extensible Attributes
data "infoblox_record_rpz_a_ipaddress" "get_rpz_a_ipaddress_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "headquarters"
  }
}

// Retrieve all RPZ A IP Address records
data "infoblox_record_rpz_a_ipaddress" "get_all_rpz_a_ipaddress_records" {}
