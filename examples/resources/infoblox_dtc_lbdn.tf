// creating a DTC LBDN record with minimal set of parameters
resource "infoblox_dtc_lbdn" "lbdn_minimal_parameters" {
  name      = "testLbdn2"
  lb_method = "ROUND_ROBIN"
  topology  = "test-topo"
  types     = ["A", "AAAA"]
}

// creating a DTC LBDN record with full set of parameters
resource "infoblox_dtc_lbdn" "lbdn_full_set_parameters" {
  name        = "testLbdn11"
  lb_method   = "TOPOLOGY"
  patterns    = ["test.com", "info.com*"]
  ttl         = 120
  topology    = "test-topo"
  disable     = true
  types       = ["A", "AAAA", "CNAME"]
  persistence = 60
  priority    = 1
  comment     = "test"

  auth_zones {
    fqdn     = "info.com"
    dns_view = "default.view2"
  }
  auth_zones {
    fqdn     = "test.com"
    dns_view = "default"
  }
  auth_zones {
    fqdn     = "test.com"
    dns_view = "default.dnsview1"
  }

  pools {
    pool  = "pool2"
    ratio = 2
  }
  pools {
    pool  = "rrpool"
    ratio = 3
  }
  pools {
    pool  = "test-pool"
    ratio = 6
  }

  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
}
