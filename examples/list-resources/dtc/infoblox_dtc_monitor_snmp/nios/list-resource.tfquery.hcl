// List DTC Monitor SNMPs using filters (NIOS)
list "infoblox_dtc_monitor_snmp" "list_dtc_monitor_snmp_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dtc-monitor-snmp"
    }
  }
  limit = 10
}

// List DTC Monitor SNMPs using extensible attributes (NIOS)
list "infoblox_dtc_monitor_snmp" "list_dtc_monitor_snmp_using_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DTC Monitor SNMPs with resource details included (NIOS)
list "infoblox_dtc_monitor_snmp" "list_dtc_monitor_snmp_with_resource" {
  provider         = infoblox
  include_resource = true
}
