# Auto-generated resource acceptance-test cases for RecordCaa.
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
    check = {
      "nios.ca_flag"            = "0"
      "nios.ca_tag"             = "issue"
      "nios.ca_value"           = "digicert.com"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
  }

}

case "ca_flag" {
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
    check = {
      "nios.ca_flag" = "0"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 1
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
    check = {
      "nios.ca_flag" = "1"
    }
  }

}

case "ca_tag" {
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
    check = {
      "nios.ca_tag" = "issue"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issuewild"
      ca_value = "digicert.com"
    }
    check = {
      "nios.ca_tag" = "issuewild"
    }
  }

}

case "ca_value" {
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
    check = {
      "nios.ca_value" = "digicert.com"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "letsencrypt.org"
    }
    check = {
      "nios.ca_value" = "letsencrypt.org"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      comment  = "comment"
    }
    check = {
      "nios.comment" = "comment"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      comment  = "updated comment"
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      creator  = "DYNAMIC"
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
      ca_flag        = 0
      ca_tag         = "issue"
      ca_value       = "digicert.com"
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
      ca_flag        = 0
      ca_tag         = "issue"
      ca_value       = "digicert.com"
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
      ca_flag        = 0
      ca_tag         = "issue"
      ca_value       = "digicert.com"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag        = 0
      ca_tag         = "issue"
      ca_value       = "digicert.com"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      disable  = true
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
      ca_flag   = 0
      ca_tag    = "issue"
      ca_value  = "digicert.com"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag   = 0
      ca_tag    = "issue"
      ca_value  = "digicert.com"
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
      ca_flag            = 0
      ca_tag             = "issue"
      ca_value           = "digicert.com"
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag            = 0
      ca_tag             = "issue"
      ca_value           = "digicert.com"
      forbid_reclamation = true
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
    }
  }

  step {
    nios {
      name     = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      ttl      = 20
    }
    check = {
      "nios.ttl" = "20"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "test" {
      nios = {
        name = "{{random3}}"
      }
    }
    resource "infoblox_zone_auth" "test" {
      nios = {
        fqdn = "{{random}}.com"
        view = infoblox_view.test.nios.name
      }
    }
    PREREQ
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      view     = "{{random3}}"
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

  step {
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "test" {
      nios = {
        name = "{{random4}}"
      }
    }
    resource "infoblox_zone_auth" "test" {
      nios = {
        fqdn = "{{random}}.com"
        view = infoblox_view.test.nios.name
      }
    }
    PREREQ
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
      view     = "{{random4}}"
    }
    check = {
      "nios.view" = "{{random4}}"
    }
  }

}
