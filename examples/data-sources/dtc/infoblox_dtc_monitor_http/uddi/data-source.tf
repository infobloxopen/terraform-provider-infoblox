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

// List all DTC HTTP Health Check resources
data "infoblox_dtc_monitor_http" "all" {}