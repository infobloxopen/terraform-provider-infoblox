// List specific RPZ CNAME IP address records using filters
list "infoblox_record_rpz_cname_ipaddress" "by_name" {
  provider = infoblox
  config {
    filters = {
      name = "11.0.0.1.rpzip.example.com"
    }
  }
}

// List specific RPZ CNAME IP address records using Extensible Attributes
list "infoblox_record_rpz_cname_ipaddress" "by_ext_attr" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List RPZ CNAME IP address records with resource details included
list "infoblox_record_rpz_cname_ipaddress" "all_with_details" {
  provider         = infoblox
  include_resource = true
}
