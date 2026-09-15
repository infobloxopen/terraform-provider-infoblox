// List specific DTC Monitor TCPs using filters
list "infoblox_dtc_monitor_tcp" "list_dtc_monitor_tcp_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_tcp_monitor"
    }
  }
  limit = 10
}

// List specific DTC Monitor TCPs using Extensible Attributes
list "infoblox_dtc_monitor_tcp" "list_dtc_monitor_tcp_using_ext_attr_filters" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DTC Monitor TCPs with resource details included
list "infoblox_dtc_monitor_tcp" "list_dtc_monitor_tcp_with_resource" {
  provider         = infoblox
  include_resource = true
}
