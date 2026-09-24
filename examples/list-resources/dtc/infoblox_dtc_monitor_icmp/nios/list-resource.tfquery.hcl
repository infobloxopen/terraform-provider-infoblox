// List specific DTC Monitor ICMPs using filters
list "infoblox_dtc_monitor_icmp" "list_dtc_monitor_icmp_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-icmp-monitor"
    }
  }
  limit = 10
}

// List specific DTC Monitor ICMPs using Extensible Attributes
list "infoblox_dtc_monitor_icmp" "list_dtc_monitor_icmp_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "us-east-1"
    }
  }
}

// List DTC Monitor ICMPs with resource details included
list "infoblox_dtc_monitor_icmp" "list_dtc_monitor_icmp_with_resource" {
  provider         = infoblox
  include_resource = true
}
