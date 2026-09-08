list "infoblox_dtc_monitor_http" "by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-monitor-http"
    }
  }
  limit = 10
}

list "infoblox_dtc_monitor_http" "by_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-east-1"
    }
  }
}

list "infoblox_dtc_monitor_http" "with_resource" {
  provider         = infoblox
  include_resource = true
}
