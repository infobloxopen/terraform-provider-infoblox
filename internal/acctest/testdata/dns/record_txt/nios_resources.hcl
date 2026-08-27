# Auto-generated resource acceptance-test cases for RecordTxt.
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
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
      view = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.text"               = "Record Text"
      "nios.name"               = "{{random2}}.{{random}}.com"
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
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
      view = infoblox_zone_auth.test.nios.view
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
      name    = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text    = "Record Text"
      comment = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      name    = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text    = "Record Text"
      comment = "This is an updated record"
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
      name    = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text    = "Record Text"
      creator = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name    = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text    = "Record Text"
      creator = "DYNAMIC"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text           = "Record Text"
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
      text           = "Record Text"
      ddns_principal = "dhcp/server2@CORP.LOCAL"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "dhcp/server2@CORP.LOCAL"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text           = "Record Text"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text           = "Record Text"
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
      name    = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text    = "Record Text"
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name    = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text    = "Record Text"
      disable = true
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text      = "Record Text"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text      = "Record Text"
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
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
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text               = "Record Text"
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text               = "Record Text"
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
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
    }
    check = {
      "nios.name" = "{{random3}}.{{random}}.com"
    }
  }

}

case "text" {
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
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
    }
    check = {
      "nios.text" = "Record Text"
    }
  }

  step {
    nios {
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Updated Text"
    }
    check = {
      "nios.text" = "Record Updated Text"
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
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
      ttl  = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      text = "Record Text"
      ttl  = 1000
    }
    check = {
      "nios.ttl" = "1000"
    }
  }

}
