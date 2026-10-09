// Retrieve a specific DTC PDP Monitor by filters
data "infoblox_dtc_monitor_pdp" "by_filters" {
  filters = {
    name = "example_dtc_monitor_pdp"
  }
}

// Retrieve DTC PDP Monitors using Extensible Attributes
data "infoblox_dtc_monitor_pdp" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC PDP Monitors
data "infoblox_dtc_monitor_pdp" "all" {}
