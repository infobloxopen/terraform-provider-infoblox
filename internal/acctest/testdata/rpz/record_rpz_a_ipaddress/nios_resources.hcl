case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.ipv4addr" = "10.10.0.1"
      "nios.name"     = "10.10.0.0/16.{{random}}.com"
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
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
      comment  = "This is a new rpz a ipaddress record"
    }
    check = {
      "nios.comment" = "This is a new rpz a ipaddress record"
    }
  }

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
      comment  = "This is an updated rpz a ipaddress record"
    }
    check = {
      "nios.comment" = "This is an updated rpz a ipaddress record"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
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
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr  = "10.10.0.1"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr  = "10.10.0.1"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ipv4addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.ipv4addr" = "10.10.0.1"
    }
  }

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.2"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.ipv4addr" = "10.10.0.2"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.name" = "10.10.0.0/16.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "10.15.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.name" = "10.15.0.0/16.{{random}}.com"
    }
  }

}

case "rp_zone" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.rp_zone" = "{{random}}.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
      ttl      = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "custom_view" {
    nios = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.custom_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.10.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.10.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
      view     = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view" = "{{random2}}"
    }
  }

}
