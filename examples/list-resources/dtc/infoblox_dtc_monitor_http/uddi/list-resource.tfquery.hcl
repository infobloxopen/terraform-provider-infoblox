// List specific DTC Monitor HTTPs using filters
list "infoblox_dtc_monitor_http" "by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-http-monitor"
    }
  }
  limit = 10
}

// List specific DTC Monitor HTTPs using Tags
list "infoblox_dtc_monitor_http" "by_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "us-east-1"
    }
  }
}

// List DTC Monitor HTTPs with resource details included
list "infoblox_dtc_monitor_http" "with_resource" {
  provider         = infoblox
  include_resource = true
}
