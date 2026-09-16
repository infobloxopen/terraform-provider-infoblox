# Auto-generated resource acceptance-test cases for RecordMx.
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
    }
    check = {
      "nios.mail_exchanger"     = "mail.example.com"
      "nios.preference"         = "10"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
      "nios.view"               = "default"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      comment        = "This is a comment."
    }
    check = {
      "nios.comment" = "This is a comment."
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      comment        = "This is an updated comment."
    }
    check = {
      "nios.comment" = "This is an updated comment."
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      creator        = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      creator        = "DYNAMIC"
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
      mail_exchanger = "mail.example.com"
      preference     = 10
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
      mail_exchanger = "mail.example.com"
      preference     = 10
      ddns_principal = "updated_ddns_principal"
      creator        = "DYNAMIC"
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      disable        = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      disable        = true
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      ext_attrs      = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      ext_attrs      = { Site = "{{random4}}" }
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
      mail_exchanger     = "mail.example.com"
      preference         = 10
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger     = "mail.example.com"
      preference         = 10
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

}

case "mail_exchanger" {
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
      mail_exchanger = "mail1.example.com"
      preference     = 10
    }
    check = {
      "nios.mail_exchanger" = "mail1.example.com"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail2.example.com"
      preference     = 10
    }
    check = {
      "nios.mail_exchanger" = "mail2.example.com"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
    }
  }

  step {
    nios {
      name           = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
    }
  }

}

case "preference" {
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
      mail_exchanger = "mail.example.com"
      preference     = 10
    }
    check = {
      "nios.preference" = "10"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 20
    }
    check = {
      "nios.preference" = "20"
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
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      ttl            = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      mail_exchanger = "mail.example.com"
      preference     = 10
      ttl            = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}

case "view" {
  backend     = "nios"
  skip        = true
  skip_reason = "helper declares prerequisite resource 'nios_dns_view' which has no buildable infoblox equivalent (not in prereq_type_map.json)"
}
