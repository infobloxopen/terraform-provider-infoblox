# Auto-generated resource acceptance-test cases for RecordRpzCnameIpaddress.
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
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name"      = "10.0.0.1.{{random}}.com"
      "nios.canonical" = "10.0.0.1"
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
  }

}

case "canonical" {
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
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "10.0.0.1"
    }
  }

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "*"
    }
  }

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = ""
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
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      comment   = "This is a new rpz cname ipaddress record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname ipaddress record"
    }
  }

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      comment   = "This is an updated rpz cname ipaddress record"
    }
    check = {
      "nios.comment" = "This is an updated rpz cname ipaddress record"
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
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "10.0.0.1.{{random}}.com"
    }
  }

  step {
    nios {
      name      = "10.0.0.2.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.2"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "10.0.0.2.{{random}}.com"
    }
  }

}

case "rp_zone" {
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
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.custom_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      view      = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view" = "{{random2}}"
    }
  }

}
