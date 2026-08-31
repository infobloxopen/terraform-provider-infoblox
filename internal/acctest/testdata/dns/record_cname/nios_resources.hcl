# Auto-generated resource acceptance-test cases for RecordCname.
case "basic" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.canonical"          = "{{random3}}.{{random}}.com"
      "nios.name"               = "{{random2}}.{{random}}.com"
      "nios.view"               = "default"
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
    }
  }

}

case "canonical" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.canonical" = "{{random3}}.{{random}}.com"
    }
  }

  step {
    nios {
      canonical = "{{random4}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.canonical" = "{{random4}}.{{random}}.com"
    }
  }

}

case "comment" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      comment   = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      comment   = "This is an updated record"
    }
    check = {
      "nios.comment" = "This is an updated record"
    }
  }

}

case "creator" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      creator   = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      creator   = "DYNAMIC"
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical      = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_principal = "DDNS_PRINCIPAL_1"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_1"
    }
  }

  step {
    nios {
      canonical      = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_principal = "DDNS_PRINCIPAL_2"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_2"
    }
  }

}

case "ddns_protected" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical      = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      canonical      = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view           = infoblox_zone_auth.test.nios.view
      ddns_protected = true
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "extattrs" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random5}}" }
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical          = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view               = infoblox_zone_auth.test.nios.view
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      canonical          = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view               = infoblox_zone_auth.test.nios.view
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "name" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random4}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random4}}.{{random}}.com"
    }
  }

}

case "ttl" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ttl       = 1000
    }
    check = {
      "nios.ttl" = "1000"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      view      = infoblox_zone_auth.test.nios.view
      ttl       = 3200
    }
    check = {
      "nios.ttl" = "3200"
    }
  }

}
