// Create a DTC Pool with Basic Fields
resource "infoblox_dtc_pool" "dtc_pool_basic" {
  uddi = {
    name   = "dtc_pool"
    method = "round_robin"
  }
}

// Create a DTC Pool with Additional Fields
resource "infoblox_dtc_pool" "dtc_pool_advanced" {
  uddi = {
    name                = "dtc_pool2"
    method              = "ratio"
    comment             = "DTC pool creation"
    pool_availability   = "any"
    server_availability = "all"
    ttl                 = 30
    disabled            = false

    tags = {
      Site = "location-1"
    }
  }
}

// Create a DTC Pool with Quorum Availability
resource "infoblox_dtc_pool" "dtc_pool_quorum" {
  uddi = {
    name                        = "dtc_pool_quorum"
    method                      = "global_availability"
    pool_availability           = "quorum"
    pool_servers_quorum         = 1
    server_availability         = "quorum"
    server_health_checks_quorum = 1
    consolidated_health_enabled = false

    tags = {
      Site = "location-1"
    }
  }
}
