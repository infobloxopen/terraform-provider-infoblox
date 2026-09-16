// Create a DTC Monitor SNMP with basic fields
resource "infoblox_dtc_monitor_snmp" "basic" {
  nios = {
    name = "dtc-monitor-snmp-basic"
  }
}

// Create a DTC Monitor SNMP with additional fields 
resource "infoblox_dtc_monitor_snmp" "additional_fields" {
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
        oid       = ".2"
        condition = "EXACT"
        first     = "10"
      },
      {
        oid = ".02"
      },
      {
        oid       = ".1"
        condition = "EXACT"
        first     = "20"
      }
    ]
  }
}
