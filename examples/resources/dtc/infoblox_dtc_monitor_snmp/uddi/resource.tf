// Create a basic DTC Monitor SNMP on UDDI (BloxOne)
resource "infoblox_dtc_monitor_snmp" "basic" {
  uddi = {
    name    = "dtc-monitor-snmp-basic"
    version = "v2c"
  }
}

// Create a DTC Monitor SNMP with additional fields on UDDI
resource "infoblox_dtc_monitor_snmp" "full" {
  uddi = {
    name       = "dtc-monitor-snmp-full"
    version    = "v2c"
    comment    = "Example DTC SNMP monitor"
    community  = "private"
    interval   = 20
    port       = 10161
    retry_down = 2
    retry_up   = 3
    timeout    = 5
    disabled   = false
    tags = {
      Site = "location-1"
    }
  }
}
