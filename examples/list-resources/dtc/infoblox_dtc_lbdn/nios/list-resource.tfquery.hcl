// List specific DTC LBDNs using filters
list "infoblox_dtc_lbdn" "list_dtc_lbdn_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-lbdn-1"
    }
  }
  limit = 10
}

// List DTC LBDNs using Extensible Attribute filters
list "infoblox_dtc_lbdn" "list_dtc_lbdn_using_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DTC LBDNs with resource details included
list "infoblox_dtc_lbdn" "list_dtc_lbdn_with_resource" {
  provider         = infoblox
  include_resource = true
}
