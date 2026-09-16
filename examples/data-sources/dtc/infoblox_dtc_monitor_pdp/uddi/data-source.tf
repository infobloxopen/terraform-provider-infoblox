// Retrieve a specific DTC PDP Health Check by filters
data "infoblox_dtc_monitor_pdp" "by_filters" {
  filters = {
    name = "example-dtc-monitor-pdp"
  }
}

// Retrieve DTC PDP Health Checks using Tag filters
data "infoblox_dtc_monitor_pdp" "by_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC PDP Health Checks
data "infoblox_dtc_monitor_pdp" "all" {}
