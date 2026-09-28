// Create a DTC Monitor TCP with Basic Fields
resource "infoblox_dtc_monitor_tcp" "dtc_monitor_tcp_basic" {
  nios = {
    name = "example_tcp_monitor"
    port = 8080
  }
}

// Create a DTC Monitor TCP with Additional Fields
resource "infoblox_dtc_monitor_tcp" "dtc_monitor_tcp_additional" {
  nios = {
    name    = "example_tcp_monitor2"
    port    = 9090
    comment = "DTC TCP monitor creation"
    ext_attrs = {
      Site = "location-1"
    }
    interval   = 45
    timeout    = 10
    retry_up   = 2
    retry_down = 3
  }
}
