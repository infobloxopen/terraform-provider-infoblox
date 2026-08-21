# DtcServer — uddi resource cases

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
    }
    check = {
      "uddi.name"          = "{{random}}"
      "uddi.address"       = "{{random_ip}}"
      "uddi.endpoint_type" = "address"
      "uddi.disabled"      = "false"
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
      name    = "{{random}}"
      address = "{{random_ip}}"
    }
  }

}

case "address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
    }
    check = {
      "uddi.address" = "{{random_ip}}"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip2}}"
       endpoint_type = "address"
    }
    check = {
      "uddi.address" = "{{random_ip2}}"
    }
  }

}

case "auto_create_response_records" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                         = "{{random}}"
      address                      = "{{random_ip}}"
      auto_create_response_records = true
    }
    check = {
      "uddi.auto_create_response_records" = "true"
      "uddi.records.#"                    = "1"
      "uddi.records.0.type"               = "A"
    }
  }

  step {
    uddi {
      name                         = "{{random}}"
      address                      = "{{random_ip}}"
      auto_create_response_records = false
    }
    check = {
      "uddi.auto_create_response_records" = "false"
      "uddi.records.#"                    = "0"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
      comment = "initial comment"
    }
    check = {
      "uddi.comment" = "initial comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      address  = "{{random_ip}}"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      address  = "{{random_ip}}"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "endpoint_type" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name          = "{{random}}"
      address       = "{{random_ip}}"
      endpoint_type = "address"
    }
    check = {
      "uddi.endpoint_type" = "address"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      fqdn          = "{{random2}}.example.com."
      endpoint_type = "fqdn"
    }
    check = {
      "uddi.endpoint_type" = "fqdn"
      "uddi.fqdn"          = "{{random2}}.example.com."
    }
  }

}

case "fqdn" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name          = "{{random}}"
      fqdn          = "{{random2}}.example.com."
      endpoint_type = "fqdn"
    }
    check = {
      "uddi.fqdn"          = "{{random2}}.example.com."
      "uddi.endpoint_type" = "fqdn"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      fqdn          = "{{random3}}.example.com."
      endpoint_type = "fqdn"
    }
    check = {
      "uddi.fqdn" = "{{random3}}.example.com."
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name    = "{{random2}}"
      address = "{{random_ip}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "records" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
      records = [
        { type = "A", rdata = { address = "192.168.1.1" } },
      ]
    }
    check = {
      "uddi.records.#"      = "1"
      "uddi.records.0.type" = "A"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
      records = [
        { type = "A",    rdata = { address = "192.168.1.1" } },
        { type = "AAAA", rdata = { address = "2001:db8::1" } },
      ]
    }
    check = {
      "uddi.records.#"      = "2"
      "uddi.records.0.type" = "A"
      "uddi.records.1.type" = "AAAA"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
      tags    = { env = "{{random2}}" }
    }
    check = {
      "uddi.tags.env" = "{{random2}}"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      address = "{{random_ip}}"
      tags    = { env = "{{random3}}" }
    }
    check = {
      "uddi.tags.env" = "{{random3}}"
    }
  }

}
