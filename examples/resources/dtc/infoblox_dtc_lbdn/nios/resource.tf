// Create DTC LBDN with Basic Fields
resource "infoblox_dtc_lbdn" "lbdn_basic" {
  nios = {
    name      = "example-lbdn-1"
    lb_method = "ROUND_ROBIN"
  }
}

// Create an Authoritative Zone to link to the LBDN
resource "infoblox_zone_auth" "example_zone" {
  nios = {
    fqdn = "example1.com"
    view = "default"
    grid_primary = [
      { name = "infoblox.localdomain" }
    ]
  }
}

// Create DTC Servers to assign to the pool
resource "infoblox_dtc_server" "example_server1" {
  nios = {
    name = "example-server-1"
    host = "10.0.0.1"
  }
}

resource "infoblox_dtc_server" "example_server2" {
  nios = {
    name = "example-server-2"
    host = "10.0.0.2"
  }
}

// Create a DTC Pool to associate with the LBDN
resource "infoblox_dtc_pool" "example_pool" {
  nios = {
    name                = "example-pool"
    lb_preferred_method = "ROUND_ROBIN"
    comment             = "Pool for example LBDN"
    servers = [
      { server = infoblox_dtc_server.example_server1.id, ratio = 1 },
      { server = infoblox_dtc_server.example_server2.id, ratio = 1 },
    ]
  }
}

// Create DTC LBDN with Auth Zones and Patterns
resource "infoblox_dtc_lbdn" "lbdn_with_zones" {
  nios = {
    name      = "example-lbdn-3"
    lb_method = "RATIO"
    auth_zones = [
      infoblox_zone_auth.example_zone.id,
    ]
    patterns = ["*.example1.com"]
    pools = [
      { pool = infoblox_dtc_pool.example_pool.id, ratio = 2 }
    ]
    comment     = "LBDN with zones and pools"
    ext_attrs   = { Site = "location-1" }
    ttl         = 300
    disable     = false
    types       = ["A", "AAAA"]
    persistence = 0
    priority    = 1
  }
}

// Create DTC LBDN with TOPOLOGY load-balancing method
// The topology must use pool-destination rules (not server-destination)
resource "infoblox_dtc_lbdn" "lbdn_topology" {
  nios = {
    name      = "example-lbdn-topology"
    lb_method = "TOPOLOGY"
    // Reference an existing DTC topology by its NIOS ref
    topology = "dtc:topology/<base64>:<topology-name>"
    pools = [
      { pool = infoblox_dtc_pool.example_pool.id, ratio = 1 }
    ]
    disable = true
  }
}
