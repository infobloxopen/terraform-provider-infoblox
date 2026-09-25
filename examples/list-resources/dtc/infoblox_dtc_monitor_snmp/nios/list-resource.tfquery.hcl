// List DTC Monitor SNMPs using filters 
list "infoblox_dtc_monitor_snmp" "list_dtc_monitor_snmp_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "dtc-monitor-snmp-basic"
    }
  }
  limit = 10
}

// List DTC Monitor SNMPs using extensible attributes 
list "infoblox_dtc_monitor_snmp" "list_dtc_monitor_snmp_using_ext_attrs" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DTC Monitor SNMPs with resource details included 
list "infoblox_dtc_monitor_snmp" "list_dtc_monitor_snmp_with_resource" {
  provider         = infoblox
  include_resource = true
}
