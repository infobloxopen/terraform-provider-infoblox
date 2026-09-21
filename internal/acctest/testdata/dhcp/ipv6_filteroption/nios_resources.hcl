# Auto-generated resource acceptance-test cases for Ipv6filteroption.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"           = "{{random}}"
      "nios.apply_as_class" = "true"
      "nios.option_space"   = "DHCPv6"
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

case "apply_as_class" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      apply_as_class = true
    }
    check = {
      "nios.apply_as_class" = "true"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      apply_as_class = false
    }
    check = {
      "nios.apply_as_class" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "Comment for the object"
    }
    check = {
      "nios.comment" = "Comment for the object"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "Updated comment for the object"
    }
    check = {
      "nios.comment" = "Updated comment for the object"
    }
  }

}

case "expression" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      expression = "(option dhcp6.server-id=\"server-id\")"
    }
    check = {
      "nios.expression" = "(option dhcp6.server-id=\"server-id\")"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      expression = "(option dhcp6.server-id=\"server-id\" AND option dhcp6.vendor-class=\"DHCPv6\")"
    }
    check = {
      "nios.expression" = "(option dhcp6.server-id=\"server-id\" AND option dhcp6.vendor-class=\"DHCPv6\")"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random3}}"
      ext_attrs = { Site = "{{random}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random}}"
    }
  }

  step {
    nios {
      name      = "{{random3}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

}

case "lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      lease_time = 1000
    }
    check = {
      "nios.lease_time" = "1000"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      lease_time = 3000
    }
    check = {
      "nios.lease_time" = "3000"
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

case "option_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      option_list = [{ name = "dhcp6.subscriber-id", value = "subscriber-id" }, { num = 23, value = "fc00::,2001:db8::" }, { name = "dhcp6.remote-id", num = 37, value = "remote-id", vendor_class = "DHCPv6" }, { name = "dhcp6.fqdn", num = 39, value = "example.com", vendor_class = "DHCPv6" }]
    }
    check = {
      "nios.option_list.#"              = "4"
      "nios.option_list.0.name"         = "dhcp6.subscriber-id"
      "nios.option_list.0.value"        = "subscriber-id"
      "nios.option_list.1.num"          = "23"
      "nios.option_list.1.value"        = "fc00::,2001:db8::"
      "nios.option_list.2.vendor_class" = "DHCPv6"
      "nios.option_list.2.name"         = "dhcp6.remote-id"
      "nios.option_list.2.num"          = "37"
      "nios.option_list.2.value"        = "remote-id"
      "nios.option_list.3.vendor_class" = "DHCPv6"
      "nios.option_list.3.name"         = "dhcp6.fqdn"
      "nios.option_list.3.num"          = "39"
      "nios.option_list.3.value"        = "example.com"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      option_list = [{ name = "dhcp6.remote-id", num = 37, value = "remote-id", vendor_class = "DHCPv6" }, { name = "dhcp6.subscriber-id", value = "subscriber-id" }]
    }
    check = {
      "nios.option_list.#"              = "2"
      "nios.option_list.0.vendor_class" = "DHCPv6"
      "nios.option_list.0.name"         = "dhcp6.remote-id"
      "nios.option_list.0.num"          = "37"
      "nios.option_list.0.value"        = "remote-id"
      "nios.option_list.1.name"         = "dhcp6.subscriber-id"
      "nios.option_list.1.value"        = "subscriber-id"
    }
  }

}

case "option_space" {
  backend     = "nios"
 parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      option_space = "DHCPv6"
    }
    check = {
      "nios.option_space" = "DHCPv6"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      option_space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.option_space" = "{{random2}}"
    }
  }

}
