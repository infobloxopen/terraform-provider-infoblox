# Auto-generated resource acceptance-test cases for ZoneAuth.
// Auth-nsgs, ACL and Tsig key has to be created before running the test cases.

case "basic" {
  backend  = "uddi"
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
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  skip                  = true
  skip_reason           = "Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]"

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
  }

}

case "fqdn" {
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend  = "uddi"
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
}

case "external_secondaries" {
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend  = "uddi"
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
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_AUTH_NSG_ID_1", "UDDI_AUTH_NSG_ID_2"]
  skip_reason       = "UDDI_AUTH_NSG_ID_1 and UDDI_AUTH_NSG_ID_2 environment variables must be set for this test to run"

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      nsgs         = ["{{uddi_auth_nsg_id_1}}"]
    }
    check = {
      "uddi.nsgs.0" = "{{uddi_auth_nsg_id_1}}"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      nsgs         = ["{{uddi_auth_nsg_id_2}}"]
    }
    check = {
      "uddi.nsgs.0" = "{{uddi_auth_nsg_id_2}}"
    }
  }
}

case "query_acl" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_namedacl" "test" {
    uddi = {
      name = "{{random}}"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_tsig_key" "test" {
    uddi = {
      name = "tsig-key-{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
  }
  PREREQ

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
      query_acl    = [{ element = "acl", acl = infoblox_namedacl.test.id }]
    }
    check = {
      "uddi.query_acl.0.element" = "acl"
      "uddi.query_acl.0.acl"     = "${infoblox_namedacl.test.id}"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      query_acl = [
        { element = "tsig_key", access = "deny",
          tsig_key = {
            key = "${infoblox_tsig_key.test.id}"
          }
      }]
    }
    depends_on = [infoblox_tsig_key.test]
    check = {
      "uddi.query_acl.0.access"         = "deny"
      "uddi.query_acl.0.element"        = "tsig_key"
      "uddi.query_acl.0.tsig_key.0.key" = "${infoblox_tsig_key.test.id}"
    }
  }

}

case "tags" {
  backend  = "uddi"
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
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_namedacl" "test" {
    uddi = {
      name = "{{random}}"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_tsig_key" "test" {
    uddi = {
      name = "tsig-key-{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
  }
  PREREQ

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
      transfer_acl = [{ element = "acl", acl = infoblox_namedacl.test.id }]
    }
    check = {
      "uddi.transfer_acl.0.element" = "acl"
      "uddi.transfer_acl.0.acl"     = "${infoblox_namedacl.test.id}"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      transfer_acl = [
        { element = "tsig_key", access = "deny",
          tsig_key = {
            key = "${infoblox_tsig_key.test.id}"
          }
      }]
    }
    depends_on = [infoblox_tsig_key.test]
    check = {
      "uddi.transfer_acl.0.access"         = "deny"
      "uddi.transfer_acl.0.element"        = "tsig_key"
      "uddi.transfer_acl.0.tsig_key.0.key" = "${infoblox_tsig_key.test.id}"
    }
  }

}

case "update_acl" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_namedacl" "test" {
    uddi = {
      name = "{{random}}"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_tsig_key" "test" {
    uddi = {
      name = "tsig-key-{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
  }
  PREREQ

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
      update_acl   = [{ element = "acl", acl = infoblox_namedacl.test.id }]
    }
    check = {
      "uddi.update_acl.0.element" = "acl"
      "uddi.update_acl.0.acl"     = "${infoblox_namedacl.test.id}"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      update_acl = [
        { element = "tsig_key", access = "deny",
          tsig_key = {
            key = "${infoblox_tsig_key.test.id}"
          }
      }]
    }
    depends_on = [infoblox_tsig_key.test]
    check = {
      "uddi.update_acl.0.access"         = "deny"
      "uddi.update_acl.0.element"        = "tsig_key"
      "uddi.update_acl.0.tsig_key.0.key" = "${infoblox_tsig_key.test.id}"
    }
  }

}

case "use_forwarders_for_subzones" {
  backend  = "uddi"
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
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_VIEW_ID_1"]
  skip_reason       = "UDDI_VIEW_ID_1 environment variable must be set for this test to run"

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      view         = "{{uddi_view_id_1}}"
    }
    check = {
      "uddi.view" = "{{uddi_view_id_1}}"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      view         = "{{uddi_view_id_1}}"
    }
    check = {
      "uddi.view" = "{{uddi_view_id_1}}"
    }
  }
}
