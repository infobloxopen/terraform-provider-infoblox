// Retrieve a specific DTC LBDN by name filter
data "infoblox_dtc_lbdn" "by_name" {
  filters = {
    name = "example-lbdn."
  }
}

// Retrieve DTC LBDNs using tag filters
data "infoblox_dtc_lbdn" "by_tags" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC LBDNs
data "infoblox_dtc_lbdn" "all" {}
