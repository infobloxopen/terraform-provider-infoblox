// Create a DTC Pool with Basic Fields
resource "infoblox_dtc_pool" "dtc_pool_basic" {
  nios = {
    name                = "dtc_pool"
    lb_preferred_method = "ROUND_ROBIN"
  }
}

// Create a DTC Pool with Additional Fields
resource "infoblox_dtc_pool" "dtc_pool_advanced" {
  nios = {
    name                = "dtc_pool2"
    lb_preferred_method = "RATIO"
    comment             = "DTC pool creation"
    ext_attrs = {
      Site = "location-1"
    }
    // servers must reference existing DTC server refs on NIOS
    servers = [
      {
        server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHNlcnZlcjE:server1"
        ratio  = 34
      },
      {
        server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHNlcnZlcjI:server2"
        ratio  = 55
      }
    ]
    monitors            = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"]
    lb_alternate_method = "NONE"
    disable             = false
    availability        = "ANY"
    ttl                 = 23
    use_ttl             = true
  }
}

// Create a DTC Pool with Quorum Availability
resource "infoblox_dtc_pool" "dtc_pool_quorum" {
  nios = {
    name                = "dtc_pool_quorum"
    lb_preferred_method = "ROUND_ROBIN"
    availability        = "QUORUM"
    quorum              = 1
  }
}
