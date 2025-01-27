// creating a DTC LBDN record with minimal set of parameters
resource "infoblox_dtc_lbdn" "lbdn_minimal_parameters" {
    name = "testLbdn2"
    lb_method = "ROUND_ROBIN"
  topology = "test-topo"
}

// creating a DTC LBDN record with full set of parameters
resource "infoblox_dtc_lbdn" "lbdn_full_set_parameters" {
  name = "testLbdn11"
  auth_zones {
    fqdn = "test.com"
    dns_view = "default"
  }
  auth_zones {
    fqdn = "test.com"
    dns_view = "default.dnsview1"
  }
  auth_zones {
    fqdn = "info.com"
    dns_view = "default.view2"
  }
  comment = "test"
  ext_attrs = jsonencode({
    "Location" = "65.8665701230204, -37.00791763398113"
  })
  lb_method = "TOPOLOGY"
  patterns = ["test.com","info.com*"]
  pools {
    pool = "pool2"
    ratio = 2
  }
  pools {
    pool = "rrpool"
    ratio = 3
  }
  pools {
    pool = "test-pool"
    ratio = 6
  }
  ttl = 120
  topology = "test-topo"
  disable = true
  types = ["A", "AAAA", "CNAME"]
  persistence = 60
  priority = 1
}

data "infoblox_dtc_lbdn" "readlbdn" {
  filters = {
    name = "testLbdn11"
#    comment = "test LBDN"
  }
  // This is just to ensure that the record has been be created
  // using 'infoblox_dtc_lbdn' resource block before the data source will be queried.
#  depends_on = [infoblox_dtc_lbdn.lbdn_record]
}

// returns matching LBDN with name and comment, if any
output "lbdn_res" {
  value = data.infoblox_dtc_lbdn.readlbdn
}