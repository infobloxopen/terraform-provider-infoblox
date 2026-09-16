# Hand-authored resource acceptance-test cases for RecordRpzAaaaIpaddress.
//
// TODO : Objects to be present in the grid before running the test cases
// Response Policy Zone - rpz-test.infoblox.com
//
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr  = "{{random_ipv6}}"
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
      ipv6addr  = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6}}"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "{{random_ipv6_2}}"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6_2}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random_ipv6_network}}.rpz-test.infoblox.com"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = "rpz-test.infoblox.com"
    }
    check = {
      "nios.name" = "{{random_ipv6_network}}.rpz-test.infoblox.com"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network2}}.rpz-test.infoblox.com"
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
      ipv6addr = "{{random_ipv6}}"
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
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "test" {
      nios = {
        name = "{{random3}}"
      }
    }
    resource "infoblox_zone_rp" "test" {
      nios = {
        fqdn = "{{random}}.com"
        view = infoblox_view.test.nios.name
      }
    }
    PREREQ
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      view     = infoblox_view.test.nios.name
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

}
