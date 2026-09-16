// List specific DTC PDP Monitors using filters
list "infoblox_dtc_monitor_pdp" "list_dtc_monitor_pdp_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_dtc_monitor_pdp"
    }
  }
  limit = 10
}

// List specific DTC PDP Monitors using Extensible Attributes
list "infoblox_dtc_monitor_pdp" "list_dtc_monitor_pdp_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DTC PDP Monitors with resource details included
list "infoblox_dtc_monitor_pdp" "list_dtc_monitor_pdp_with_resource" {
  provider         = infoblox
  include_resource = true
}
