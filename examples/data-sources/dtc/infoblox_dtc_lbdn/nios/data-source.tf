// Retrieve a specific DTC LBDN by name filter
data "infoblox_dtc_lbdn" "by_name" {
  filters = {
    name = "example-lbdn"
  }
}

// Retrieve DTC LBDNs using Extensible Attributes
data "infoblox_dtc_lbdn" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC LBDNs
data "infoblox_dtc_lbdn" "all" {}
