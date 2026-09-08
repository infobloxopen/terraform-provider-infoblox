// Create a DTC PDP Monitor with Basic Fields
resource "infoblox_dtc_monitor_pdp" "monitor_pdp_basic" {
  nios = {
    name = "example_dtc_monitor_pdp"
  }
}

// Create a DTC PDP Monitor with Additional Fields
resource "infoblox_dtc_monitor_pdp" "monitor_pdp_additional" {
  nios = {
    name       = "example_dtc_monitor_pdp_additional"
    comment    = "This is a DTC Monitor PDP"
    ext_attrs  = { Site = "location-1" }
    interval   = 5
    port       = 80
    retry_down = 2
    retry_up   = 2
    timeout    = 10
  }
}
