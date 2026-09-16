# Auto-generated resource acceptance-test cases for RecordDname.
case "basic" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random2}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target"             = "{{random}}.com"
      "nios.name"               = "{{random2}}.com"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random2}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.example.com"
      view   = infoblox_zone_auth.test.nios.view
    }
  }

}

case "comment" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = infoblox_zone_auth.test.nios.fqdn
      target  = "{{random}}.com"
      comment = "comment"
      view    = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.comment" = "comment"
    }
  }

  step {
    nios {
      name    = infoblox_zone_auth.test.nios.fqdn
      target  = "{{random}}.com"
      comment = "updated comment"
      view    = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.comment" = "updated comment"
    }
  }

}

case "creator" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = infoblox_zone_auth.test.nios.fqdn
      target  = "{{random}}.com"
      creator = "STATIC"
      view    = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name    = infoblox_zone_auth.test.nios.fqdn
      target  = "{{random}}.com"
      creator = "DYNAMIC"
      view    = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = infoblox_zone_auth.test.nios.fqdn
      target         = "{{random}}.com"
      ddns_principal = "ddns_principal"
      creator        = "DYNAMIC"
      view           = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ddns_principal" = "ddns_principal"
    }
  }

  step {
    nios {
      name           = infoblox_zone_auth.test.nios.fqdn
      target         = "{{random}}.com"
      ddns_principal = "updated_ddns_principal"
      creator        = "DYNAMIC"
      view           = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ddns_principal" = "updated_ddns_principal"
    }
  }

}

case "ddns_protected" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = infoblox_zone_auth.test.nios.fqdn
      target         = "{{random}}.com"
      ddns_protected = false
      view           = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = infoblox_zone_auth.test.nios.fqdn
      target         = "{{random}}.com"
      ddns_protected = true
      view           = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = infoblox_zone_auth.test.nios.fqdn
      target  = "{{random}}.com"
      disable = false
      view    = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name    = infoblox_zone_auth.test.nios.fqdn
      target  = "{{random}}.com"
      disable = true
      view    = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = infoblox_zone_auth.test.nios.fqdn
      target    = "{{random}}.com"
      ext_attrs = { Site = "{{random4}}" }
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

  step {
    nios {
      name      = infoblox_zone_auth.test.nios.fqdn
      target    = "{{random}}.com"
      ext_attrs = { Site = "{{random5}}" }
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ext_attrs.Site" = "{{random5}}"
    }
  }

}

case "forbid_reclamation" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name               = infoblox_zone_auth.test.nios.fqdn
      target             = "{{random}}.com"
      forbid_reclamation = false
      view               = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

  step {
    nios {
      name               = infoblox_zone_auth.test.nios.fqdn
      target             = "{{random}}.com"
      forbid_reclamation = true
      view               = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

}

case "name" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  resource "infoblox_zone_auth" "updated_zone" {
    nios = {
      fqdn = "{{random4}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random3}}.com"
    }
  }

  step {
    nios {
      name   = infoblox_zone_auth.updated_zone.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.updated_zone.nios.view
    }
    check = {
      "nios.name" = "{{random4}}.com"
    }
  }

}

case "target" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random4}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target" = "{{random}}.com"
    }
  }

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random2}}.com"
      view   = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target" = "{{random2}}.com"
    }
  }

}

case "ttl" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      ttl    = 10
      view   = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      ttl    = 20
      view   = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.ttl" = "20"
    }
  }

}
