// List specific Dtc Monitor Tcps using filters
list "infoblox_dtc_monitor_tcp" "list_dtc_monitor_tcp_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "DTC TCP monitor creation"
    }
  }
  limit = 10
}

// List specific Dtc Monitor Tcps using Tags
list "infoblox_dtc_monitor_tcp" "list_dtc_monitor_tcp_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dtc Monitor Tcps with resource details included
list "infoblox_dtc_monitor_tcp" "list_dtc_monitor_tcp_with_resource" {
  provider         = infoblox
  include_resource = true
}
