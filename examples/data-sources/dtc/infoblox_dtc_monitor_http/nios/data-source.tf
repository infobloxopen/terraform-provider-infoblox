// Retrieve a specific DTC HTTP Monitor by filters
data "infoblox_dtc_monitor_http" "by_name" {
  filters = {
    name = "example-monitor-http"
  }
}

// Retrieve specific DTC HTTP Monitors by extensible attributes
data "infoblox_dtc_monitor_http" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "us-east-1"
  }
}

// Retrieve all DTC HTTP Monitors
data "infoblox_dtc_monitor_http" "all" {}
