# Auto-generated resource acceptance-test cases for ForwardNsg.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name"            = "{{random}}"
      "uddi.forwarders_only" = "false"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      comment = "This Forward NSG is created through Terraform"
    }
    check = {
      "uddi.comment" = "This Forward NSG is created through Terraform"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "This Forward NSG was created through Terraform"
    }
    check = {
      "uddi.comment" = "This Forward NSG was created through Terraform"
    }
  }

}

case "external_forwarders_address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      external_forwarders = [{ address = "192.168.1.0" }]
    }
    check = {
      "uddi.external_forwarders.#"         = "1"
      "uddi.external_forwarders.0.address" = "192.168.1.0"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      external_forwarders = [{ address = "192.168.1.1" }]
    }
    check = {
      "uddi.external_forwarders.#"         = "1"
      "uddi.external_forwarders.0.address" = "192.168.1.1"
    }
  }

}

case "external_forwarders" {
  backend = "uddi"

  step {
    uddi {
      name                = "{{random}}"
      external_forwarders = [{ address = "192.168.1.0", fqdn = "terraform-acc-forward-ext." }]
    }
    check = {
      "uddi.external_forwarders.#"         = "1"
      "uddi.external_forwarders.0.fqdn"    = "terraform-acc-forward-ext."
      "uddi.external_forwarders.0.address" = "192.168.1.0"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      external_forwarders = [{ address = "192.168.1.0", fqdn = "terraform-acc-forward-ext-1." }]
    }
    check = {
      "uddi.external_forwarders.#"         = "1"
      "uddi.external_forwarders.0.fqdn"    = "terraform-acc-forward-ext-1."
      "uddi.external_forwarders.0.address" = "192.168.1.0"
    }
  }

}

case "forwarders_only" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name            = "{{random}}"
      forwarders_only = true
    }
    check = {
      "uddi.forwarders_only" = "true"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      forwarders_only = false
    }
    check = {
      "uddi.forwarders_only" = "false"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "hosts" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name  = "{{random}}"
      hosts = ["dns/host/1008608"]
    }
    check = {
      "uddi.hosts.#" = "1"
      "uddi.hosts.0" = "dns/host/1008608"
    }
  }

  step {
    uddi {
      name  = "{{random}}"
      hosts = ["dns/host/1390921"]
    }
    check = {
      "uddi.hosts.#" = "1"
      "uddi.hosts.0" = "dns/host/1390921"
    }
  }

}

case "internal_forwarders" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      internal_forwarders = ["dns/host/1008608"]
    }
    check = {
      "uddi.internal_forwarders.#" = "1"
      "uddi.internal_forwarders.0" = "dns/host/1008608"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      internal_forwarders = ["dns/host/1390921"]
    }
    check = {
      "uddi.internal_forwarders.#" = "1"
      "uddi.internal_forwarders.0" = "dns/host/1390921"
    }
  }

}

case "nsgs" {
  backend     = "uddi"
  skip        = true
  skip_reason = "helper declares 2 'bloxone_dns_forward_nsg' resource blocks with no single func_call target (ambiguous which is the resource under test)"
}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
