// List specific Dtc Monitor Icmps using filters
list "infoblox_dtc_monitor_icmp" "list_dtc_monitor_icmp_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "DTC ICMP monitor for health checks"
    }
  }
  limit = 10
}

// List specific Dtc Monitor Icmps using Tags
list "infoblox_dtc_monitor_icmp" "list_dtc_monitor_icmp_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Dtc Monitor Icmps with resource details included
list "infoblox_dtc_monitor_icmp" "list_dtc_monitor_icmp_with_resource" {
  provider         = infoblox
  include_resource = true
}
