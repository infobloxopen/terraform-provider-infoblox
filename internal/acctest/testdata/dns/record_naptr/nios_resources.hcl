# Auto-generated resource acceptance-test cases for RecordNaptr.
case "basic" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.order"              = "10"
      "nios.preference"         = "10"
      "nios.replacement"        = "."
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
      "nios.view"               = "default"
    }
  }

}

case "disappears" {
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
  }

}

case "comment" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      comment     = "comment"
    }
    check = {
      "nios.comment" = "comment"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      comment     = "updated comment"
    }
    check = {
      "nios.comment" = "updated comment"
    }
  }

}

case "creator" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      creator     = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      creator     = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order          = 10
      preference     = 10
      replacement    = "."
      ddns_principal = "ddns_principal"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "ddns_principal"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order          = 10
      preference     = 10
      replacement    = "."
      ddns_principal = "updated_ddns_principal"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "updated_ddns_principal"
    }
  }

}

case "ddns_protected" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order          = 10
      preference     = 10
      replacement    = "."
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order          = 10
      preference     = 10
      replacement    = "."
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      disable     = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      disable     = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "flags" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "U"
    }
    check = {
      "nios.flags" = "U"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "S"
    }
    check = {
      "nios.flags" = "S"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "A"
    }
    check = {
      "nios.flags" = "A"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "P"
    }
    check = {
      "nios.flags" = "P"
    }
  }

}

case "forbid_reclamation" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order              = 10
      preference         = 10
      replacement        = "."
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order              = 10
      preference         = 10
      replacement        = "."
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

}

case "name" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
  }

  step {
    nios {
      name        = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
  }

}

case "order" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.order" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 20
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.order" = "20"
    }
  }

}

case "preference" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.preference" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 20
      replacement = "."
    }
    check = {
      "nios.preference" = "20"
    }
  }

}

case "regexp" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      regexp      = "!^.*$!sip:jdoe@corpabc.com!"
    }
    check = {
      "nios.regexp" = "!^.*$!sip:jdoe@corpabc.com!"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
    }
    check = {
      "nios.regexp" = "!^.*$!sip:jdoe@corpxyz.com!"
    }
  }

}

case "replacement" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.replacement" = "."
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 20
      replacement = "test.com"
    }
    check = {
      "nios.replacement" = "test.com"
    }
  }

}

case "services" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      services    = "http+E2U"
    }
    check = {
      "nios.services" = "http+E2U"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 20
      replacement = "."
      services    = "SIPS+D2T"
    }
    check = {
      "nios.services" = "SIPS+D2T"
    }
  }

}

case "ttl" {
  backend = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 10
      replacement = "."
      ttl         = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      order       = 10
      preference  = 20
      replacement = "."
      ttl         = 20
    }
    check = {
      "nios.ttl" = "20"
    }
  }

}
