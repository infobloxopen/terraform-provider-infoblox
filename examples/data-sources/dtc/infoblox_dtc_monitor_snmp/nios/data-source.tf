// Read a specific DTC Monitor SNMP by filters 
data "infoblox_dtc_monitor_snmp" "by_name" {
  filters = {
    name = "dtc-monitor-snmp-basic"
  }
}

// Read DTC Monitor SNMPs by extensible attributes 
data "infoblox_dtc_monitor_snmp" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Read all DTC Monitor SNMPs
data "infoblox_dtc_monitor_snmp" "all" {}
