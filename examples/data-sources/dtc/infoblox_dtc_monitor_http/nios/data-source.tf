// Filter by name
data "infoblox_dtc_monitor_http" "by_name" {
  filters = {
    name = "example-monitor-http"
  }
}

// Filter by extensible attribute
data "infoblox_dtc_monitor_http" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "us-east-1"
  }
}

// List all DTC Monitor HTTP resources
data "infoblox_dtc_monitor_http" "all" {}