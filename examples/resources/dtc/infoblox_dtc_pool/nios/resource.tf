// Create a DTC Pool with Basic Fields
resource "infoblox_dtc_pool" "dtc_pool_basic" {
  nios = {
    name                = "dtc_pool"
    lb_preferred_method = "ROUND_ROBIN"
  }
}

// Create DTC Servers and assign them to a Pool.
resource "infoblox_dtc_server" "dtc_server1" {
  nios = {
    name = "dtc_server1"
    host = "10.0.0.1"
  }
}

resource "infoblox_dtc_server" "dtc_server2" {
  nios = {
    name = "dtc_server2"
    host = "10.0.0.2"
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
    servers = [
      {
        server = infoblox_dtc_server.dtc_server1.id
        ratio  = 34
      },
      {
        server = infoblox_dtc_server.dtc_server2.id
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
    monitors            = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"]
    quorum              = 1
  }
}
