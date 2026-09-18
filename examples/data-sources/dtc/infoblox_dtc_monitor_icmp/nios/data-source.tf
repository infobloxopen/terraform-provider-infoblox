// Retrieve a specific DTC Monitor ICMP using filters
data "infoblox_dtc_monitor_icmp" "get_dtc_monitor_icmp_using_filters" {
  filters = {
    name = "example-icmp-monitor"
  }
}

// Retrieve specific DTC Monitor ICMPs using Extensible Attributes
data "infoblox_dtc_monitor_icmp" "get_dtc_monitor_icmp_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "us-east-1"
  }
}

// Retrieve all DTC Monitor ICMPs
data "infoblox_dtc_monitor_icmp" "get_all_dtc_monitor_icmps" {}
