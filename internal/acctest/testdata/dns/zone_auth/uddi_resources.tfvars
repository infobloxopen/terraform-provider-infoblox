# Auto-generated resource acceptance-test cases for ZoneAuth.
case "basic" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.fqdn"                        = "{{random}}.com."
      "uddi.primary_type"                = "cloud"
      "uddi.disabled"                    = "false"
      "uddi.gss_tsig_enabled"            = "false"
      "uddi.initial_soa_serial"          = "1"
      "uddi.notify"                      = "false"
      "uddi.use_forwarders_for_subzones" = "true"
    }
  }

}

case "disappears" {
  backend = "uddi"
  disappears = true
  expect_non_empty_plan = true
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
  }

}

case "fqdn" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.fqdn"         = "{{random}}.com."
      "uddi.primary_type" = "cloud"
    }
  }

  step {
    uddi {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.fqdn"         = "{{random2}}.com."
      "uddi.primary_type" = "cloud"
    }
  }

}

case "primary_type" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.fqdn"         = "{{random}}.com."
      "uddi.primary_type" = "cloud"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "external"
    }
    check = {
      "uddi.fqdn"         = "{{random}}.com."
      "uddi.primary_type" = "external"
    }
  }

}

case "comment" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      comment      = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      comment      = "test comment update"
    }
    check = {
      "uddi.comment" = "test comment update"
    }
  }

}

case "disabled" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      disabled     = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      disabled     = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

}

case "external_primaries" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn               = "{{random}}.com."
      primary_type       = "external"
      external_primaries = [{ fqdn = "tf-infoblox-test.com.", address = "192.168.10.10", type = "primary" }]
    }
    check = {
      "uddi.external_primaries.0.fqdn"    = "tf-infoblox-test.com."
      "uddi.external_primaries.0.address" = "192.168.10.10"
      "uddi.external_primaries.0.type"    = "primary"
    }
  }

  step {
    uddi {
      fqdn               = "{{random}}.com."
      primary_type       = "external"
      external_primaries = [{ fqdn = "tf-infoblox.com.", address = "192.168.11.11", type = "primary" }]
    }
    check = {
      "uddi.external_primaries.0.fqdn"    = "tf-infoblox.com."
      "uddi.external_primaries.0.address" = "192.168.11.11"
      "uddi.external_primaries.0.type"    = "primary"
    }
  }

  step {
    uddi {
      fqdn               = "{{random}}.com."
      primary_type       = "external"
      external_primaries = [{ fqdn = "tf-infoblox-test.com.", address = "192.168.10.10", type = "nsg" }]
    }
  }

}

case "external_secondaries" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn                 = "{{random}}.com."
      primary_type         = "external"
      external_secondaries = [{ fqdn = "tf-infoblox-test.com.", address = "192.168.10.10" }]
    }
    check = {
      "uddi.external_secondaries.0.fqdn"    = "tf-infoblox-test.com."
      "uddi.external_secondaries.0.address" = "192.168.10.10"
    }
  }

  step {
    uddi {
      fqdn                 = "{{random}}.com."
      primary_type         = "external"
      external_secondaries = [{ fqdn = "tf-infoblox.com.", address = "192.168.11.11" }]
    }
    check = {
      "uddi.external_secondaries.0.fqdn"    = "tf-infoblox.com."
      "uddi.external_secondaries.0.address" = "192.168.11.11"
    }
  }

}

case "gss_tsig_enabled" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn             = "{{random}}.com."
      primary_type     = "cloud"
      gss_tsig_enabled = false
    }
    check = {
      "uddi.gss_tsig_enabled" = "false"
    }
  }

  step {
    uddi {
      fqdn             = "{{random}}.com."
      primary_type     = "cloud"
      gss_tsig_enabled = true
    }
    check = {
      "uddi.gss_tsig_enabled" = "true"
    }
  }

}

case "inheritance_sources" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn                        = "{{random}}.com."
      primary_type                = "cloud"
      inheritance_sources         = { gss_tsig_enabled = { action = "inherit" }, notify = { action = "inherit" }, transfer_acl = { action = "inherit" }, useforwardersforsubzones = { action = "inherit" } }
      gss_tsig_enabled            = true
      notify                      = true
      transfer_acl                = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
      use_forwarders_for_subzones = true
    }
    check = {
      "uddi.inheritance_sources.gss_tsig_enabled.action" = "inherit"
    }
  }

  step {
    uddi {
      fqdn                        = "{{random}}.com."
      primary_type                = "cloud"
      inheritance_sources         = { gss_tsig_enabled = { action = "override" }, notify = { action = "override" }, transfer_acl = { action = "override" }, useforwardersforsubzones = { action = "override" } }
      gss_tsig_enabled            = true
      notify                      = true
      transfer_acl                = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
      use_forwarders_for_subzones = true
    }
    check = {
      "uddi.inheritance_sources.gss_tsig_enabled.action" = "override"
    }
  }

}

case "initial_soa_serial" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn               = "{{random}}.com."
      primary_type       = "cloud"
      initial_soa_serial = 1
    }
    check = {
      "uddi.initial_soa_serial" = "1"
    }
  }

  step {
    uddi {
      fqdn               = "{{random}}.com."
      primary_type       = "cloud"
      initial_soa_serial = 2
    }
    check = {
      "uddi.initial_soa_serial" = "2"
    }
  }

}

case "notify" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      notify       = false
    }
    check = {
      "uddi.notify" = "false"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      notify       = true
    }
    check = {
      "uddi.notify" = "true"
    }
  }

}

case "nsgs" {
  backend = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ns_group_unknown" "one" {
    uddi = {
      name = "one"
    }
  }
  resource "infoblox_ns_group_unknown" "two" {
    uddi = {
      name = "two"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      nsgs         = [infoblox_ns_group_unknown.one.id]
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      nsgs         = [infoblox_ns_group_unknown.two.id]
    }
  }

}

case "query_acl" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      query_acl    = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.query_acl.0.access"  = "allow"
      "uddi.query_acl.0.element" = "ip"
      "uddi.query_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      query_acl    = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.query_acl.0.access"  = "deny"
      "uddi.query_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      query_acl    = [{ element = "acl", acl = infoblox_acl_unknown.test.id }]
    }
    check = {
      "uddi.query_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.query_acl.0.access"  = "deny"
      "uddi.query_acl.0.element" = "tsig_key"
    }
  }

}

case "tags" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      tags         = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      tags         = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "transfer_acl" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      transfer_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.transfer_acl.0.access"  = "allow"
      "uddi.transfer_acl.0.element" = "ip"
      "uddi.transfer_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      transfer_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.transfer_acl.0.access"  = "deny"
      "uddi.transfer_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      transfer_acl = [{ element = "acl", acl = infoblox_acl_unknown.test.id }]
    }
    check = {
      "uddi.transfer_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.transfer_acl.0.access"  = "deny"
      "uddi.transfer_acl.0.element" = "tsig_key"
    }
  }

}

case "update_acl" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      update_acl   = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.update_acl.0.access"  = "allow"
      "uddi.update_acl.0.element" = "ip"
      "uddi.update_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      update_acl   = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.update_acl.0.access"  = "deny"
      "uddi.update_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      update_acl   = [{ element = "acl", acl = infoblox_acl_unknown.test.id }]
    }
    check = {
      "uddi.update_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
    check = {
      "uddi.update_acl.0.access"  = "deny"
      "uddi.update_acl.0.element" = "tsig_key"
    }
  }

}

case "use_forwarders_for_subzones" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      fqdn                        = "{{random}}.com."
      primary_type                = "cloud"
      use_forwarders_for_subzones = true
    }
    check = {
      "uddi.use_forwarders_for_subzones" = "true"
    }
  }

  step {
    uddi {
      fqdn                        = "{{random}}.com."
      primary_type                = "cloud"
      use_forwarders_for_subzones = false
    }
    check = {
      "uddi.use_forwarders_for_subzones" = "false"
    }
  }

}

case "view" {
  backend = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_view" "two" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      view         = infoblox_view.one.id
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      view         = infoblox_view.two.id
    }
  }

}

case "zone_authority" {
  backend = "uddi"
  expect_non_empty_plan = true
  parallel = true

  step {
    uddi {
      fqdn           = "{{random}}.com."
      primary_type   = "cloud"
      zone_authority = { default_ttl = 28800, expire = 2419200, mname = "ns.b1ddi", negative_ttl = 900, refresh = 10800, retry = 3600, rname = "hostmaster", use_default_mname = false }
    }
    check = {
      "uddi.zone_authority.default_ttl"       = "28800"
      "uddi.zone_authority.expire"            = "2419200"
      "uddi.zone_authority.mname"             = "ns.b1ddi"
      "uddi.zone_authority.negative_ttl"      = "900"
      "uddi.zone_authority.refresh"           = "10800"
      "uddi.zone_authority.retry"             = "3600"
      "uddi.zone_authority.rname"             = "hostmaster"
      "uddi.zone_authority.use_default_mname" = "false"
    }
  }

  step {
    uddi {
      fqdn           = "{{random}}.com."
      primary_type   = "cloud"
      zone_authority = { default_ttl = 30000, expire = 2519200, mname = "ns.b1ddi", negative_ttl = 800, refresh = 11800, retry = 3700, rname = "hostmaster", use_default_mname = false }
    }
    check = {
      "uddi.zone_authority.default_ttl"       = "30000"
      "uddi.zone_authority.expire"            = "2519200"
      "uddi.zone_authority.mname"             = "ns.b1ddi"
      "uddi.zone_authority.negative_ttl"      = "800"
      "uddi.zone_authority.refresh"           = "11800"
      "uddi.zone_authority.retry"             = "3700"
      "uddi.zone_authority.rname"             = "hostmaster"
      "uddi.zone_authority.use_default_mname" = "false"
    }
  }

}
