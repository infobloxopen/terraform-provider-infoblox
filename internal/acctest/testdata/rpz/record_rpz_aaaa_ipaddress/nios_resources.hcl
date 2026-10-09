# Hand-authored resource acceptance-test cases for RecordRpzAaaaIpaddress.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6}}"
      "nios.name"     = "{{random_ipv6_network}}.{{random}}.com"
      "nios.rp_zone"  = "{{random}}.com"
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      comment  = "test comment"
    }
    check = {
      "nios.comment" = "test comment"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "value1" }
    }
    check = {
      "nios.ext_attrs.Site" = "value1"
    }
  }

  step {
    nios {
      name      = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6}}"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6_2}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6_2}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random_ipv6_network}}.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network2}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random_ipv6_network2}}.{{random}}.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      ttl      = 600
    }
    check = {
      "nios.ttl" = "600"
    }
  }

  step {
    nios {
      name     = "{{random_ipv6_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
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
