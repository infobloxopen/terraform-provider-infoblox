// Create a DTC Monitor ICMP with required fields only
resource "infoblox_dtc_monitor_icmp" "example" {
  uddi = {
    name = "example-icmp-monitor"
  }
}

// Create a DTC Monitor ICMP with all optional fields
resource "infoblox_dtc_monitor_icmp" "example_all_fields" {
  uddi = {
    name       = "example-icmp-monitor-full"
    comment    = "DTC ICMP monitor for health checks"
    disabled   = false
    interval   = 45
    timeout    = 10
    retry_up   = 2
    retry_down = 3
    tags = {
      env  = "production"
      Site = "us-east-1"
    }
  }
}
