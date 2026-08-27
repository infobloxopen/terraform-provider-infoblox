# Hand-authored datasource acceptance-test cases for RecordRpzAaaaIpaddress.
# rp_zone is hardcoded to "rpz-test.infoblox.com" (persistent zone on the test NIOS grid)
# because infoblox_zone_rp is not yet registered in the unified provider.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv6addr", "nios.name", "nios.rp_zone", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.ipv6addr", "nios.name", "nios.rp_zone", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr  = "2001:db8::10"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "value1" }
    }
  }

}
