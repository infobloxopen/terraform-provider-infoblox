// Create a DTC PDP Health Check with Basic Fields
resource "infoblox_dtc_monitor_pdp" "health_check_pdp_basic" {
  uddi = {
    name = "example-dtc-monitor-pdp"
  }
}

// Create a DTC PDP Health Check with Additional Fields
resource "infoblox_dtc_monitor_pdp" "health_check_pdp_additional" {
  uddi = {
    name       = "example-dtc-monitor-pdp-additional"
    comment    = "This is a DTC PDP Health Check"
    disabled   = false
    interval   = 15
    port       = 80
    retry_down = 2
    retry_up   = 2
    timeout    = 10
    tags       = { Site = "location-1" }
  }
}
