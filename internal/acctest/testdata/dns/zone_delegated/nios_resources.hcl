// Objects to be present on the grid for testing
// delegation_group, delegation_group_1 - nsgroup:delegation

# Auto-generated resource acceptance-test cases for ZoneDelegated.
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
    }
    check = {
      "nios.disable"                  = "false"
      "nios.enable_rfc2317_exclusion" = "false"
      "nios.locked"                   = "false"
      "nios.ms_ad_integrated"         = "false"
      "nios.ms_ddns_mode"             = "NONE"
      "nios.view"                     = "default"
      "nios.zone_format"              = "FORWARD"
    }
  }

}

case "disappears" {
  backend               = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      comment     = "This is a delegated zone"
    }
    check = {
      "nios.comment" = "This is a delegated zone"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      comment     = "This is an updated delegated zone"
    }
    check = {
      "nios.comment" = "This is an updated delegated zone"
    }
  }

}

case "delegate_to" {
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.3" }]
    }
    check = {
      "nios.delegate_to.0.name"    = "{{random3}}.com"
      "nios.delegate_to.0.address" = "10.0.0.3"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.4" }]
    }
    check = {
      "nios.delegate_to.0.name"    = "{{random3}}.com"
      "nios.delegate_to.0.address" = "10.0.0.4"
    }
  }

}

case "delegated_ttl" {
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
      fqdn          = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to   = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      delegated_ttl = 3600
    }
    check = {
      "nios.delegated_ttl" = "3600"
    }
  }

  step {
    nios {
      fqdn          = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to   = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      delegated_ttl = 7200
    }
    check = {
      "nios.delegated_ttl" = "7200"
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      disable     = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      disable     = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "enable_rfc2317_exclusion" {
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
      fqdn                     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to              = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      enable_rfc2317_exclusion = true
    }
    check = {
      "nios.enable_rfc2317_exclusion" = "true"
    }
  }

  step {
    nios {
      fqdn                     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to              = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      enable_rfc2317_exclusion = false
    }
    check = {
      "nios.enable_rfc2317_exclusion" = "false"
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ext_attrs   = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ext_attrs   = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "locked" {
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      locked      = true
    }
    check = {
      "nios.locked" = "true"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      locked      = false
    }
    check = {
      "nios.locked" = "false"
    }
  }

}

case "ms_ad_integrated" {
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
      fqdn             = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to      = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ms_ad_integrated = true
    }
    check = {
      "nios.ms_ad_integrated" = "true"
    }
  }

  step {
    nios {
      fqdn             = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to      = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ms_ad_integrated = false
    }
    check = {
      "nios.ms_ad_integrated" = "false"
    }
  }

}

case "ms_ddns_mode" {
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
      fqdn         = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to  = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ms_ddns_mode = "NONE"
    }
    check = {
      "nios.ms_ddns_mode" = "NONE"
    }
  }

  step {
    nios {
      fqdn         = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to  = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ms_ddns_mode = "ANY"
    }
    check = {
      "nios.ms_ddns_mode" = "ANY"
    }
  }

  step {
    nios {
      fqdn             = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to      = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ms_ddns_mode     = "SECURE"
      ms_ad_integrated = true
    }
    check = {
      "nios.ms_ddns_mode" = "SECURE"
    }
  }

}

case "ns_group" {
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ns_group    = "delegation_group"
    }
    check = {
      "nios.ns_group" = "delegation_group"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      ns_group    = "delegation_group_1"
    }
    check = {
      "nios.ns_group" = "delegation_group_1"
    }
  }

}

case "prefix" {
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
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      prefix      = "{{random4}}"
    }
    check = {
      "nios.prefix" = "{{random4}}"
    }
  }

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}.com", address = "10.0.0.1" }]
      prefix      = "{{random4}}"
    }
    check = {
      "nios.prefix" = "{{random4}}"
    }
  }

}

case "zone_format_ipv4" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
resource "infoblox_zone_auth" "parent_auth_reverse_zone_ipv4" {
  nios = {
    fqdn = "111.0.1.0/24"
    view = "default"
    zone_format = "IPV4"
  }
}
PREREQ

  step {
    nios {
      fqdn        = "111.0.1.10/32"
      delegate_to = [{ name = "{{random}}.com", address = "10.0.0.1" }]
      zone_format = "IPV4"
    }
    depends_on = [infoblox_zone_auth.parent_auth_reverse_zone_ipv4]
    check = {
      "nios.zone_format" = "IPV4"
    }
  }

}

case "zone_format_ipv6" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
resource "infoblox_zone_auth" "parent_auth_reverse_zone_ipv6" {
  nios = {
    fqdn = "2001::/64"
    view = "default"
    zone_format = "IPV6"
  }
}
PREREQ

  step {
    nios {
      fqdn        = "2001::1/128"
      delegate_to = [{ name = "{{random}}.com", address = "10.0.0.1" }]
      zone_format = "IPV6"
    }
    depends_on = [infoblox_zone_auth.parent_auth_reverse_zone_ipv6]
    check = {
      "nios.zone_format" = "IPV6"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_dns_view" {
    nios = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random2}}.com"
      view = infoblox_view.test_dns_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      fqdn        = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random4}}.com", address = "10.0.0.1" }]
      view        = infoblox_view.test_dns_view.nios.name
    }
    check = {
      "nios.view" = "{{random}}"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: step has no Config helper call
case "import" {
  backend     = "nios"
  skip        = true
  skip_reason = "step has no Config helper call"
  parallel    = true

  step {
    nios {
      fqdn        = "{{random}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random2}}.com", address = "10.0.0.1" }]
    }
  }

}
