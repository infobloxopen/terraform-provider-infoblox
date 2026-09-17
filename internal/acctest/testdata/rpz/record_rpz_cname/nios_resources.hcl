# Auto-generated resource acceptance-test cases for RecordRpzCname.

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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name"      = "{{random2}}.{{random}}.com"
      "nios.canonical" = ""
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = ""
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

  # "" = Block Domain (NXDOMAIN); "*" = Block Domain (No Data);
  # leading-label = Passthru; FQDN = Substitution; "infoblox-passthru" (wildcard name) = Wildcard Passthru.
  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = ""
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = ""
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "*"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random3}}.com"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "{{random3}}.com"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "sub.{{random4}}.com"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "sub.{{random4}}.com"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      comment   = "This is a new rpz cname record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname record"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      comment   = "This is an updated rpz cname record"
    }
    check = {
      "nios.comment" = "This is an updated rpz cname record"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name      = "{{random3}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random3}}.{{random}}.com"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ttl       = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

  # ttl removed from config after being set: Optional+Computed must absorb
  # whatever the API reports instead of raising a "was null, but now N" diff.
  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "custom_view" {
    nios = {
      name = "{{random3}}"
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
      name      = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "*"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      view      = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view"    = "{{random3}}"
      "nios.rp_zone" = "{{random}}.com"
    }
  }

}
