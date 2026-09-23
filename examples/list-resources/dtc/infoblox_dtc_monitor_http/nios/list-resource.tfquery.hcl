// List DTC Monitor HTTPs using filters
list "infoblox_dtc_monitor_http" "by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-monitor-http"
    }
  }
  limit = 10
}

// List DTC Monitor HTTPs using extensible attributes
list "infoblox_dtc_monitor_http" "by_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-east-1"
    }
  }
}

// List DTC Monitor HTTPs with resource details included
list "infoblox_dtc_monitor_http" "with_resource" {
  provider         = infoblox
  include_resource = true
}
