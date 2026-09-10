list "infoblox_dtc_monitor_http" "by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-http-monitor"
    }
  }
  limit = 10
}

list "infoblox_dtc_monitor_http" "by_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "us-east-1"
    }
  }
}

list "infoblox_dtc_monitor_http" "with_resource" {
  provider         = infoblox
  include_resource = true
}
