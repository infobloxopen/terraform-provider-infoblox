// List specific DTC PDP Health Checks using filters
list "infoblox_dtc_monitor_pdp" "list_dtc_health_check_pdp_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-dtc-monitor-pdp"
    }
  }
  limit = 10
}

// List specific DTC PDP Health Checks using Tag filters
list "infoblox_dtc_monitor_pdp" "list_dtc_health_check_pdp_using_tag_filters" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List DTC PDP Health Checks with resource details included
list "infoblox_dtc_monitor_pdp" "list_dtc_health_check_pdp_with_resource" {
  provider         = infoblox
  include_resource = true
}
