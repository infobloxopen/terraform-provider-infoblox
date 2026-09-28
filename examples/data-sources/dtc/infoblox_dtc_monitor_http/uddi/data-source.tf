// Retrieve a specific DTC HTTP Health Check by filters
data "infoblox_dtc_monitor_http" "by_name" {
  filters = {
    name = "example-http-monitor"
  }
}

// Retrieve DTC HTTP Health Checks using Tag filters
data "infoblox_dtc_monitor_http" "by_tags" {
  tag_filters = {
    Site = "us-east-1"
  }
}

// List all DTC HTTP Health Check resources
data "infoblox_dtc_monitor_http" "all" {}
