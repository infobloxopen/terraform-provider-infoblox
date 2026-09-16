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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.fqdn"                         = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      compartment_id     = "c4695."
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.compartment_id" = "c4695."
    }
  }

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      compartment_id     = ""
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      comment            = "Delegation zone is created by Terraform"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.comment" = "Delegation zone is created by Terraform"
    }
  }

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      comment            = "Delegation zone was created by Terraform"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.comment" = "Delegation zone was created by Terraform"
    }
  }

}

case "delegation_servers" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
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
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.1", fqdn = "ns2.com." }]
      view               = infoblox_view.test.id
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
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "", fqdn = "ns3.com." }]
      view               = infoblox_view.test.id
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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      disabled           = false
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.disabled" = "false"
    }
  }

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      disabled           = true
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.fqdn" = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
    }
  }

  step {
    uddi {
      fqdn               = "{{random4}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check = {
      "uddi.fqdn" = "{{random4}}.${infoblox_zone_auth.test.uddi.fqdn}"
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
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      tags               = { tag1 = "value1", tag2 = "value2" }
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
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
      fqdn               = "{{random4}}.${infoblox_zone_auth.test.uddi.fqdn}"
      tags               = { tag2 = "value2changed", tag3 = "value3" }
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
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
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random2}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random3}}.${infoblox_zone_auth.test.uddi.fqdn}"
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
    check_pair = {
      "uddi.view" = infoblox_view.test.id
    }
  }

}
