// Read a specific DTC Monitor SNMP by filters (NIOS)
data "infoblox_dtc_monitor_snmp" "by_name" {
  filters = {
    name = "dtc-monitor-snmp-basic"
  }
}

// Read DTC Monitor SNMPs by extensible attributes (NIOS)
data "infoblox_dtc_monitor_snmp" "by_ext_attrs" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Read all DTC Monitor SNMPs (NIOS)
data "infoblox_dtc_monitor_snmp" "all" {}
