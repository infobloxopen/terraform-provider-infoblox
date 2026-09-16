// Create a DTC Monitor TCP with Basic Fields
resource "infoblox_dtc_monitor_tcp" "dtc_monitor_tcp_basic" {
  uddi = {
    name = "example_tcp_monitor"
    port = 8080
  }
}

// Create a DTC Monitor TCP with Additional Fields
resource "infoblox_dtc_monitor_tcp" "dtc_monitor_tcp_additional" {
  uddi = {
    name       = "example_tcp_monitor2"
    port       = 9090
    comment    = "DTC TCP monitor creation"
    interval   = 45
    timeout    = 10
    retry_up   = 2
    retry_down = 3
    disabled   = false
    tags = {
      Site = "location-1"
    }
  }
}
