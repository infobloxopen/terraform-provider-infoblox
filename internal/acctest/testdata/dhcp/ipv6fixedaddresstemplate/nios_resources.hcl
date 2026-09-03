# Auto-generated resource acceptance-test cases for Ipv6fixedaddresstemplate.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"           = "{{random}}"
      "nios.comment"        = ""
      "nios.valid_lifetime" = "43200"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "domain_name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      domain_name = "example.com"
    }
    check = {
      "nios.domain_name" = "example.com"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      domain_name = "example.org"
    }
    check = {
      "nios.domain_name" = "example.org"
    }
  }

}

case "domain_name_servers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      domain_name_servers = ["2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"]
    }
    check = {
      "nios.domain_name_servers.#" = "3"
      "nios.domain_name_servers.0" = "2001:4860:4860::8888"
      "nios.domain_name_servers.1" = "2001:4860:4860::9999"
      "nios.domain_name_servers.2" = "2001:4860:4860::8899"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      domain_name_servers = ["2001:4860:4860::8888", "2001:4860:4860::8844"]
    }
    check = {
      "nios.domain_name_servers.#" = "2"
      "nios.domain_name_servers.0" = "2001:4860:4860::8888"
      "nios.domain_name_servers.1" = "2001:4860:4860::8844"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      logic_filter_rules = [{ filter = "ipv6_option_filter", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      logic_filter_rules = [{ filter = "ipv6_option_filter1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "number_of_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      number_of_addresses = 10
      offset              = 10
    }
    check = {
      "nios.number_of_addresses" = "10"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      number_of_addresses = 20
      offset              = 10
    }
    check = {
      "nios.number_of_addresses" = "20"
    }
  }

}

case "offset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      offset              = 10
      number_of_addresses = 20
    }
    check = {
      "nios.offset" = "10"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      offset              = 15
      number_of_addresses = 20
    }
    check = {
      "nios.offset" = "15"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      options = [{ name = "domain-name", num = "15", value = "example.com" }, { num = "37", value = "remote-id", vendor_class = "DHCPv6" }, { name = "dhcp6.subscriber-id", value = "subscriber-id", vendor_class = "DHCPv6" }]
    }
    check = {
      "nios.options.#"              = "3"
      "nios.options.0.name"         = "domain-name"
      "nios.options.0.num"          = "15"
      "nios.options.0.value"        = "example.com"
      "nios.options.1.num"          = "37"
      "nios.options.1.value"        = "remote-id"
      "nios.options.1.vendor_class" = "DHCPv6"
      "nios.options.2.name"         = "dhcp6.subscriber-id"
      "nios.options.2.value"        = "subscriber-id"
      "nios.options.2.vendor_class" = "DHCPv6"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      options = [{ name = "domain-name", num = "15", value = "example.org" }, { num = "37", value = "remote-id-updated", vendor_class = "DHCPv6" }]
    }
    check = {
      "nios.options.#"              = "2"
      "nios.options.0.name"         = "domain-name"
      "nios.options.0.num"          = "15"
      "nios.options.0.value"        = "example.org"
      "nios.options.1.num"          = "37"
      "nios.options.1.value"        = "remote-id-updated"
      "nios.options.1.vendor_class" = "DHCPv6"
    }
  }

}

case "preferred_lifetime" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      preferred_lifetime = 200
      valid_lifetime     = 43200
    }
    check = {
      "nios.preferred_lifetime" = "200"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      preferred_lifetime = 600
      valid_lifetime     = 43200
    }
    check = {
      "nios.preferred_lifetime" = "600"
    }
  }

}

case "valid_lifetime" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      valid_lifetime = 200
    }
    check = {
      "nios.valid_lifetime" = "200"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      valid_lifetime = 400
    }
    check = {
      "nios.valid_lifetime" = "400"
    }
  }

}
