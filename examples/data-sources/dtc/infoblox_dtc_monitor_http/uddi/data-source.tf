// List all DTC HTTP Health Check resources
data "infoblox_dtc_monitor_http" "all" {}

// Filter by name
data "infoblox_dtc_monitor_http" "by_name" {
  filters = {
    name = "example-http-monitor"
  }
}

// Filter by tag
data "infoblox_dtc_monitor_http" "by_tags" {
  tag_filters = {
    Site = "us-east-1"
  }
}
