// List specific RPZ AAAA IP Address records using filters
list "infoblox_record_rpz_aaaa_ipaddress" "list_rpz_aaaa_ipaddress_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "2001:db8::/64.rpz.example.com"
    }
  }
}

// List specific RPZ AAAA IP Address records using Extensible Attributes
list "infoblox_record_rpz_aaaa_ipaddress" "list_rpz_aaaa_ipaddress_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "headquarters"
    }
  }
}

// List RPZ AAAA IP Address records with resource details included
list "infoblox_record_rpz_aaaa_ipaddress" "list_rpz_aaaa_ipaddress_with_resource" {
  provider         = infoblox
  include_resource = true
}
