// Retrieve a specific DTC Monitor TCP by filters
data "infoblox_dtc_monitor_tcp" "get_dtc_monitor_tcp_using_filters" {
  filters = {
    name = "example_tcp_monitor"
  }
}

// Retrieve specific DTC Monitor TCPs using Tags
data "infoblox_dtc_monitor_tcp" "get_dtc_monitor_tcp_using_tags" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all DTC Monitor TCPs
data "infoblox_dtc_monitor_tcp" "get_all_dtc_monitor_tcps" {}
