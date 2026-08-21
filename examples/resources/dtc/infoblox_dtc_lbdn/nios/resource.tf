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
    fqdn = "example.com"
    view = "default"
    grid_primary = [
      { name = "infoblox.localdomain" }
    ]
  }
}

// Create DTC LBDN with Auth Zones and Patterns
// Pools are referenced by their NIOS ref (dtc:pool/<ref>:<name>)
resource "infoblox_dtc_lbdn" "lbdn_with_zones" {
  nios = {
    name      = "example-lbdn-2"
    lb_method = "RATIO"
    auth_zones = [
      infoblox_zone_auth.example_zone.id,
    ]
    patterns = ["*.example.com"]
    // Reference an existing DTC pool by its NIOS ref
    pools = [
      { pool = "dtc:pool/<base64>:<pool-name>", ratio = 2 }
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
      { pool = "dtc:pool/<base64>:<pool-name>", ratio = 1 }
    ]
    disable = true
  }
}
