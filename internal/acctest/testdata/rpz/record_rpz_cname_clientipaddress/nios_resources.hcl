# Auto-generated resource acceptance-test cases for RecordRpzCnameClientipaddress.
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
      name      = "12.0.0.1.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = "rpz-passthru"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.name"      = "12.0.0.1.{{random}}.com"
      "nios.canonical" = "rpz-passthru"
      "nios.rp_zone"   = "{{random}}.com"
      "nios.view"      = "default"
      "nios.disable"   = "false"
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
      name      = "12.0.0.2.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
    }
  }

}

case "canonical" {
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
      name      = "12.0.0.3.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.canonical" = ""
    }
  }

  step {
    nios {
      name      = "12.0.0.3.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.canonical" = "*"
    }
  }

  step {
    nios {
      name      = "12.0.0.3.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = "rpz-passthru"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
    }
    check = {
      "nios.canonical" = "rpz-passthru"
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
      name      = "12.0.0.4.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      comment   = "This is a new rpz cname client IP address record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname client IP address record"
    }
  }

  step {
    nios {
      name      = "12.0.0.4.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      comment   = "This is an updated rpz cname client IP address record"
    }
    check = {
      "nios.comment" = "This is an updated rpz cname client IP address record"
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
      name      = "12.0.0.5.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "12.0.0.5.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      disable   = true
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
      name      = "12.0.0.6.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "12.0.0.6.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
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
      name      = "12.0.0.7.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
    }
    check = {
      "nios.name" = "12.0.0.7.{{random}}.com"
    }
  }

  step {
    nios {
      name      = "12.0.0.8.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
    }
    check = {
      "nios.name" = "12.0.0.8.{{random}}.com"
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
      name      = "12.0.0.9.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
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
      name      = "12.0.0.10.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "12.0.0.10.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = "default"
      ttl       = 0
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
      name      = "12.0.0.12.${infoblox_zone_rp.test_zone.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      view      = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view" = "{{random2}}"
    }
  }

}
