// Create a DTC Monitor ICMP with required fields only
resource "infoblox_dtc_monitor_icmp" "example" {
  nios = {
    name = "example-icmp-monitor"
  }
}

// Create a DTC Monitor ICMP with all optional fields
resource "infoblox_dtc_monitor_icmp" "example_all_fields" {
  nios = {
    name       = "example-icmp-monitor-full"
    comment    = "DTC ICMP monitor for health checks"
    interval   = 45
    timeout    = 10
    retry_up   = 2
    retry_down = 3
    ext_attrs = {
      Site = "us-east-1"
    }
  }
}
