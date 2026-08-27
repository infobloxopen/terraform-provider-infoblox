# Hand-authored resource acceptance-test cases for RecordRpzAaaaIpaddress.
# rp_zone is hardcoded to "rpz-test.infoblox.com" (persistent zone on the test NIOS grid)
# because infoblox_zone_rp is not yet registered in the unified provider.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.ipv6addr" = "2001:db8::10"
      "nios.name"     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      "nios.rp_zone"  = "rpz-test.infoblox.com"
      "nios.view"     = "default"
      "nios.disable"  = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
      comment  = "test comment"
    }
    check = {
      "nios.comment" = "test comment"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
      comment  = "test comment update"
    }
    check = {
      "nios.comment" = "test comment update"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr  = "2001:db8::10"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "value1" }
    }
    check = {
      "nios.ext_attrs.Site" = "value1"
    }
  }

  step {
    nios {
      name      = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr  = "2001:db8::10"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "value2" }
    }
    check = {
      "nios.ext_attrs.Site" = "value2"
    }
  }

}

case "ipv6addr" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.ipv6addr" = "2001:db8::10"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::20"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.ipv6addr" = "2001:db8::20"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.name" = "{{random_ipv6_network}}.rpz-test.infoblox.com"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network2}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.name" = "{{random_ipv6_network2}}.rpz-test.infoblox.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
      ttl      = 600
    }
    check = {
      "nios.ttl" = "600"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
      ttl      = 3600
    }
    check = {
      "nios.ttl" = "3600"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "2001:db8::10"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.view" = "default"
    }
  }

}
