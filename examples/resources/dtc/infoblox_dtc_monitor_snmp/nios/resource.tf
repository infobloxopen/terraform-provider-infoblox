// Create a basic DTC Monitor SNMP on NIOS
resource "infoblox_dtc_monitor_snmp" "basic" {
  nios = {
    name = "dtc-monitor-snmp-basic"
  }
}

// Create a DTC Monitor SNMP with additional fields on NIOS
resource "infoblox_dtc_monitor_snmp" "full" {
  nios = {
    name       = "dtc-monitor-snmp-full"
    comment    = "Example DTC SNMP monitor"
    community  = "private"
    interval   = 20
    port       = 10161
    retry_down = 2
    retry_up   = 3
    timeout    = 5
    version    = "V2C"
    ext_attrs = {
      Site = "location-1"
    }
    oids = [
      {
        oid       = ".1.3.6.1.2.1.1.1.0"
        type      = "STRING"
        condition = "EXACT"
        first     = "Linux"
      }
    ]
  }
}
