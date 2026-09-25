// List specific RPZ A IP Address records using filters
list "infoblox_record_rpz_a_ipaddress" "list_rpz_a_ipaddress_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "10.10.0.0/16.rpz.example.com"
    }
  }
}

// List specific RPZ A IP Address records using Extensible Attributes
list "infoblox_record_rpz_a_ipaddress" "list_rpz_a_ipaddress_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "headquarters"
    }
  }
}

// List RPZ A IP Address records with resource details included
list "infoblox_record_rpz_a_ipaddress" "list_rpz_a_ipaddress_with_resource" {
  provider         = infoblox
  include_resource = true
}
