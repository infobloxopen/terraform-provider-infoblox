// Objects to be present on the grid for testing
// dns host
// internal forwarders

# Auto-generated resource acceptance-test cases for ZoneForward.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn = "{{random}}.com."
    }
    check = {
      "uddi.fqdn" = "{{random}}.com."
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

  step {
    uddi {
      fqdn = "{{random}}.com."
    }
  }

}

case "fqdn" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn = "{{random}}.com."
    }
    check = {
      "uddi.fqdn" = "{{random}}.com."
    }
  }

  step {
    uddi {
      fqdn = "{{random2}}.com."
    }
    check = {
      "uddi.fqdn" = "{{random2}}.com."
    }
  }

}

case "compartment_id" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn           = "{{random}}.com."
      compartment_id = "c4695."
    }
    check = {
      "uddi.compartment_id" = "c4695."
    }
  }

  step {
    uddi {
      fqdn           = "{{random}}.com."
      compartment_id = ""
    }
    check = {
      "uddi.compartment_id" = ""
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn    = "{{random}}.com."
      comment = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      fqdn    = "{{random}}.com."
      comment = "test comment update"
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
      fqdn     = "{{random}}.com."
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      fqdn     = "{{random}}.com."
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "external_forwarders_address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn                = "{{random}}.com."
      external_forwarders = [{ address = "192.168.10.10", fqdn = "abc.com." }]
    }
    check = {
      "uddi.external_forwarders.0.address" = "192.168.10.10"
      "uddi.external_forwarders.0.fqdn"    = "abc.com."
    }
  }

  step {
    uddi {
      fqdn                = "{{random}}.com."
      external_forwarders = [{ address = "192.168.11.11", fqdn = "def.com." }]
    }
    check = {
      "uddi.external_forwarders.0.address" = "192.168.11.11"
      "uddi.external_forwarders.0.fqdn"    = "def.com."
    }
  }

}

case "external_forwarders_fqdn" {
  backend = "uddi"

  step {
    uddi {
      fqdn                = "{{random}}.com."
      external_forwarders = [{ address = "192.168.10.10", fqdn = "tf-infoblox-test.com." }]
    }
    check = {
      "uddi.external_forwarders.0.address" = "192.168.10.10"
      "uddi.external_forwarders.0.fqdn"    = "tf-infoblox-test.com."
    }
  }

  step {
    uddi {
      fqdn                = "{{random}}.com."
      external_forwarders = [{ address = "192.168.11.11", fqdn = "tf-infoblox.com." }]
    }
    check = {
      "uddi.external_forwarders.0.address" = "192.168.11.11"
      "uddi.external_forwarders.0.fqdn"    = "tf-infoblox.com."
    }
  }

}

case "forward_only" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn         = "{{random}}.com."
      forward_only = false
    }
    check = {
      "uddi.forward_only" = "false"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      forward_only = true
    }
    check = {
      "uddi.forward_only" = "true"
    }
  }

}

case "hosts" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn = "{{random}}.com."
      hosts = ["dns/host/470522"]
    }
    check = {
      "uddi.hosts.#" = "1"
    }
  }

  step {
    uddi {
      fqdn = "{{random}}.com."
      hosts = ["dns/host/470521"]
    }
    check = {
      "uddi.hosts.#" = "1"
    }
  }

}

case "internal_forwarders" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn = "{{random}}.com."
      internal_forwarders = ["dns/host/470521"]
    }
    check = {
      "uddi.internal_forwarders.#" = "1"
    }
  }

  step {
    uddi {
      fqdn = "{{random}}.com."
      internal_forwarders = ["dns/host/470522"]
    }
    check = {
      "uddi.internal_forwarders.#" = "1"
    }
  }

}

case "nsgs" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_forward_nsg" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_forward_nsg" "two" {
    uddi = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn = "{{random}}.com."
      nsgs = [infoblox_forward_nsg.one.id]
    }
  }

  step {
    uddi {
      fqdn = "{{random}}.com."
      nsgs = [infoblox_forward_nsg.two.id]
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      fqdn = "{{random}}.com."
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      fqdn = "{{random}}.com."
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
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
  resource "infoblox_view" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_view" "two" {
    uddi = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    uddi {
      fqdn = "{{random}}.com."
      view = infoblox_view.one.id
    }
  }

  step {
    uddi {
      fqdn = "{{random2}}.com."
      view = infoblox_view.two.id
    }
  }

}
