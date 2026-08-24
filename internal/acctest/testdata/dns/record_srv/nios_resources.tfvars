# Auto-generated resource acceptance-test cases for RecordSrv.
case "basic" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.name"               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      "nios.target"             = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      "nios.port"               = "80"
      "nios.priority"           = "10"
      "nios.weight"             = "360"
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ
  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
  }

}

case "comment" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      comment  = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      comment  = "This is a updated record"
    }
    check = {
      "nios.comment" = "This is a updated record"
    }
  }

}

case "creator" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      creator  = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend  = "nios"
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
      target         = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port           = 80
      priority       = 10
      weight         = 360
      ddns_principal = "dhcp/server1@CORP.LOCAL"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "dhcp/server1@CORP.LOCAL"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target         = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port           = 80
      priority       = 10
      weight         = 360
      ddns_principal = "dhcp/server2@CORP.LOCAL"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "dhcp/server2@CORP.LOCAL"
    }
  }

}

case "ddns_protected" {
  backend  = "nios"
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
      target         = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port           = 80
      priority       = 10
      weight         = 360
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target         = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port           = 80
      priority       = 10
      weight         = 360
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
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
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ
  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target    = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port      = 80
      priority  = 10
      weight    = 360
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target    = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port      = 80
      priority  = 10
      weight    = 360
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "forbid_reclamation" {
  backend  = "nios"
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
      target             = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port               = 80
      priority           = 10
      weight             = 360
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target             = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port               = 80
      priority           = 10
      weight             = 360
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "name" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.name" = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.name" = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
    }
  }

}

case "port" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.port" = "80"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 8080
      priority = 10
      weight   = 360
    }
    check = {
      "nios.port" = "8080"
    }
  }

}

case "priority" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.priority" = "10"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 1
      weight   = 360
    }
    check = {
      "nios.priority" = "1"
    }
  }

}

case "target" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.target" = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.target" = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
    }
  }

}

case "ttl" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
      ttl      = 1000
    }
    check = {
      "nios.ttl" = "1000"
    }
  }

}

case "weight" {
  backend  = "nios"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 360
    }
    check = {
      "nios.weight" = "360"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target   = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      port     = 80
      priority = 10
      weight   = 720
    }
    check = {
      "nios.weight" = "720"
    }
  }

}

case "view" {
  backend     = "nios"
  skip        = true
  skip_reason = "helper declares prerequisite resource 'nios_dns_view' which has no buildable infoblox equivalent (not in prereq_type_map.json)"
}
