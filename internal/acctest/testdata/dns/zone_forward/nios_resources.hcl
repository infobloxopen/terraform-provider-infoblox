
// Objects to be present on the grid for testing
// ns_group1, ns_group2 - Forwarding Member
// ensg1, ensg2 - Forward/Stub Server

# Auto-generated resource acceptance-test cases for ZoneForward.
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
    }
    check = {
      "nios.fqdn"                  = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      "nios.disable"               = "false"
      "nios.disable_ns_generation" = "false"
      "nios.forwarders_only"       = "false"
      "nios.locked"                = "false"
      "nios.ms_ad_integrated"      = "false"
      "nios.ms_ddns_mode"          = "NONE"
      "nios.view"                  = "default"
      "nios.zone_format"           = "FORWARD"
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      comment           = "Zone forward comment"
    }
    check = {
      "nios.comment" = "Zone forward comment"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      comment           = "Zone forward comment updated"
    }
    check = {
      "nios.comment" = "Zone forward comment updated"
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      disable           = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      disable           = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "disable_ns_generation" {
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
      fqdn                  = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group     = "ensg1"
      disable_ns_generation = true
    }
    check = {
      "nios.disable_ns_generation" = "true"
    }
  }

  step {
    nios {
      fqdn                  = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group     = "ensg1"
      disable_ns_generation = false
    }
    check = {
      "nios.disable_ns_generation" = "false"
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ext_attrs         = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ext_attrs         = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "external_ns_group" {
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
    }
    check = {
      "nios.external_ns_group" = "ensg1"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg2"
    }
    check = {
      "nios.external_ns_group" = "ensg2"
    }
  }

}

case "forward_to" {
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
      fqdn       = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      forward_to = [{ name = "example1.org", address = "10.1.0.1" }]
    }
    check = {
      "nios.forward_to.#"         = "1"
      "nios.forward_to.0.name"    = "example1.org"
      "nios.forward_to.0.address" = "10.1.0.1"
    }
  }

  step {
    nios {
      fqdn       = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      forward_to = [{ name = "example2.org", address = "10.1.0.2" }]
    }
    check = {
      "nios.forward_to.#"         = "1"
      "nios.forward_to.0.name"    = "example2.org"
      "nios.forward_to.0.address" = "10.1.0.2"
    }
  }

}

case "forwarders_only" {
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      forwarders_only   = false
    }
    check = {
      "nios.forwarders_only" = "false"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      forwarders_only   = true
    }
    check = {
      "nios.forwarders_only" = "true"
    }
  }

}

case "forwarding_servers" {
  backend      = "nios"
  parallel     = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      forwarding_servers = [{
        name                    = "{{grid_master_hostname}}"
        forwarders_only         = true
        use_override_forwarders = true
        forward_to              = [{ name = "example1.org", address = "11.10.1.2" }]
      }]
    }
    check = {
      "nios.forwarding_servers.#"                         = "1"
      "nios.forwarding_servers.0.name"                    = "{{grid_master_hostname}}"
      "nios.forwarding_servers.0.forwarders_only"         = "true"
      "nios.forwarding_servers.0.use_override_forwarders" = "true"
      "nios.forwarding_servers.0.forward_to.0.name"       = "example1.org"
      "nios.forwarding_servers.0.forward_to.0.address"    = "11.10.1.2"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      forwarding_servers = [{
        name                    = "{{grid_master_hostname}}"
        forwarders_only         = false
        use_override_forwarders = false
        forward_to              = [{ name = "example22.org", address = "11.10.11.22" }]
      }]
    }
    check = {
      "nios.forwarding_servers.#"                         = "1"
      "nios.forwarding_servers.0.name"                    = "{{grid_master_hostname}}"
      "nios.forwarding_servers.0.forwarders_only"         = "false"
      "nios.forwarding_servers.0.use_override_forwarders" = "false"
      "nios.forwarding_servers.0.forward_to.0.name"       = "example22.org"
      "nios.forwarding_servers.0.forward_to.0.address"    = "11.10.11.22"
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      locked            = true
    }
    check = {
      "nios.locked" = "true"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      locked            = false
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ms_ad_integrated  = true
    }
    check = {
      "nios.ms_ad_integrated" = "true"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ms_ad_integrated  = false
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ms_ddns_mode      = "ANY"
    }
    check = {
      "nios.ms_ddns_mode" = "ANY"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ms_ddns_mode      = "NONE"
    }
    check = {
      "nios.ms_ddns_mode" = "NONE"
    }
  }

}

case "ns_group" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_nsgroup" "test_ns_group1" {
    nios = {
      name = "{{random4}}"
      grid_primary = [
        {
          name = "{{grid_master_hostname}}"
        }
      ]
    }
  }
  resource "infoblox_nsgroup" "test_ns_group2" {
    nios = {
      name = "{{random5}}"
      grid_primary = [
        {
          name = "{{grid_master_hostname}}"
        }
      ]
    }
  }
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ns_group          = "ns_group1"
    }
    depends_on = [infoblox_nsgroup.test_ns_group1, infoblox_nsgroup.test_ns_group2]
    check = {
      "nios.ns_group" = "ns_group1"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      ns_group          = "ns_group2"
    }
    depends_on = [infoblox_nsgroup.test_ns_group1, infoblox_nsgroup.test_ns_group2]
    check = {
      "nios.ns_group" = "ns_group2"
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
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      prefix            = "0-127"
    }
    check = {
      "nios.prefix" = "0-127"
    }
  }

  step {
    nios {
      fqdn              = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      prefix            = "128/26"
    }
    check = {
      "nios.prefix" = "128/26"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_dns_view" {
    nios = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.test_dns_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      fqdn              = "{{random}}.${infoblox_zone_auth.test.nios.fqdn}"
      external_ns_group = "ensg1"
      view              = infoblox_view.test_dns_view.nios.name
    }
    check = {
      "nios.view" = "{{random2}}"
    }
  }

}
