# Auto-generated resource acceptance-test cases for ZoneDelegated.
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.fqdn"                         = "test.123."
      "uddi.delegation_servers.#"         = "1"
      "uddi.delegation_servers.0.address" = "12.0.0.0"
      "uddi.delegation_servers.0.fqdn"    = "ns1.com."
      "uddi.disabled"                     = "false"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  skip                  = true
  skip_reason           = "t.Skip: Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
  }

}

case "compartment_id" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      compartment_id     = "c4695."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.compartment_id" = "c4695."
    }
  }

  step {
    uddi {
      fqdn               = "test.123."
      compartment_id     = ""
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.compartment_id" = ""
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      comment            = "Delegation zone is created by Terraform"
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.comment" = "Delegation zone is created by Terraform"
    }
  }

  step {
    uddi {
      fqdn               = "test.123."
      comment            = "Delegation zone was created by Terraform"
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.comment" = "Delegation zone was created by Terraform"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
#   "fqdn": %q
case "delegation_servers" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn = "test.123."
      view = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.delegation_servers.#"         = "1"
      "uddi.delegation_servers.0.address" = "12.0.0.0"
      "uddi.delegation_servers.0.fqdn"    = "ns1.com."
    }
  }

  step {
    uddi {
      fqdn = "test.123."
      view = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.delegation_servers.#"         = "1"
      "uddi.delegation_servers.0.address" = "12.0.0.1"
      "uddi.delegation_servers.0.fqdn"    = "ns2.com."
    }
  }

  step {
    uddi {
      fqdn = "test.123."
      view = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.delegation_servers.#"         = "1"
      "uddi.delegation_servers.0.address" = ""
      "uddi.delegation_servers.0.fqdn"    = "ns3.com."
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      disabled           = false
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.disabled" = "false"
    }
  }

  step {
    uddi {
      fqdn               = "test.123."
      disabled           = true
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.disabled" = "true"
    }
  }

}

case "fqdn" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.fqdn" = "test.123."
    }
  }

  step {
    uddi {
      fqdn               = "test1.123."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.fqdn" = "test1.123."
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      tags               = { tag1 = "value1", tag2 = "value2" }
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      fqdn               = "test.123."
      tags               = { tag2 = "value2changed", tag3 = "value3" }
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "view" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "%[1]q" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "test.123."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
    }
    depends_on = [infoblox_zone_auth.test]
  }

  step {
    uddi {
      fqdn               = "test.123."
      delegation_servers = { address = "12.0.0.0", fqdn = "ns1.com." }
    }
    depends_on = [infoblox_zone_auth.test]
  }

}
