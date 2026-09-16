// Create a DTC Pool with Basic Fields
resource "infoblox_dtc_pool" "dtc_pool_basic" {
  uddi = {
    name   = "dtc_pool"
    method = "round_robin"
  }
}

// Create DTC Servers and assign them to a Pool.
resource "infoblox_dtc_server" "dtc_server1" {
  uddi = {
    name    = "dtc_server1"
    address = "10.0.0.1"
  }
}

resource "infoblox_dtc_server" "dtc_server2" {
  uddi = {
    name    = "dtc_server2"
    address = "10.0.0.2"
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
    servers = [
      {
        server_id = infoblox_dtc_server.dtc_server1.id
        weight    = 1
      },
      {
        server_id = infoblox_dtc_server.dtc_server2.id
        weight    = 2
      }
    ]
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

    tags = {
      Site = "location-1"
    }
  }
}
