# Auto-generated resource acceptance-test cases for Filteroption.
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
      "nios.option_space"   = "DHCP"
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

case "bootfile" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      bootfile = "pxelinux.0"
    }
    check = {
      "nios.bootfile" = "pxelinux.0"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      bootfile = "pxelinux.efi"
    }
    check = {
      "nios.bootfile" = "pxelinux.efi"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      bootserver = "boot.example.com"
    }
    check = {
      "nios.bootserver" = "boot.example.com"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      bootserver = "boot-updated.example.com"
    }
    check = {
      "nios.bootserver" = "boot-updated.example.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "test filter option"
    }
    check = {
      "nios.comment" = "test filter option"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "updated filter option"
    }
    check = {
      "nios.comment" = "updated filter option"
    }
  }

}

case "expression" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      expression = "option domain-name=\"example.com\""
    }
    check = {
      "nios.expression" = "option domain-name=\"example.com\""
    }
  }

  step {
    nios {
      name       = "{{random}}"
      expression = "option ntp-servers=\"2.2.2.2\""
    }
    check = {
      "nios.expression" = "option ntp-servers=\"2.2.2.2\""
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
      lease_time = 600
    }
    check = {
      "nios.lease_time" = "600"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      lease_time = 1200
    }
    check = {
      "nios.lease_time" = "1200"
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

case "next_server" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      next_server = "192.168.1.1"
    }
    check = {
      "nios.next_server" = "192.168.1.1"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      next_server = "192.168.1.2"
    }
    check = {
      "nios.next_server" = "192.168.1.2"
    }
  }

}

case "option_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      option_list = [{ name = "domain-name", value = "example.com" }]
    }
    check = {
      "nios.option_list.0.name"  = "domain-name"
      "nios.option_list.0.value" = "example.com"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      option_list = [{ name = "time-offset", num = 2, value = "1200" }]
    }
    check = {
      "nios.option_list.0.name"  = "time-offset"
      "nios.option_list.0.value" = "1200"
    }
  }

}

case "option_space" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      option_space = "DHCP"
    }
    check = {
      "nios.option_space" = "DHCP"
    }
  }

}

case "pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      pxe_lease_time = 1200
    }
    check = {
      "nios.pxe_lease_time" = "1200"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      pxe_lease_time = 1800
    }
    check = {
      "nios.pxe_lease_time" = "1800"
    }
  }

}
