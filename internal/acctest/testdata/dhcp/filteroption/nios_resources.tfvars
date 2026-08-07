# Auto-generated resource acceptance-test cases for Filteroption.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    # comment and expression are omitted by NIOS when unset, so they flatten to
    # null rather than "" and cannot be asserted here.
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
      bootfile = "BOOTFILE_REPLACE_ME"
    }
    check = {
      "nios.bootfile" = "BOOTFILE_REPLACE_ME"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      bootfile = "BOOTFILE_UPDATE_REPLACE_ME"
    }
    check = {
      "nios.bootfile" = "BOOTFILE_UPDATE_REPLACE_ME"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      bootserver = "BOOTSERVER_REPLACE_ME"
    }
    check = {
      "nios.bootserver" = "BOOTSERVER_REPLACE_ME"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      bootserver = "BOOTSERVER_UPDATE_REPLACE_ME"
    }
    check = {
      "nios.bootserver" = "BOOTSERVER_UPDATE_REPLACE_ME"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "COMMENT_REPLACE_ME"
    }
    check = {
      "nios.comment" = "COMMENT_REPLACE_ME"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "COMMENT_UPDATE_REPLACE_ME"
    }
    check = {
      "nios.comment" = "COMMENT_UPDATE_REPLACE_ME"
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
      next_server = "NEXT_SERVER_REPLACE_ME"
    }
    check = {
      "nios.next_server" = "NEXT_SERVER_REPLACE_ME"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      next_server = "NEXT_SERVER_UPDATE_REPLACE_ME"
    }
    check = {
      "nios.next_server" = "NEXT_SERVER_UPDATE_REPLACE_ME"
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

# The generator emitted a prerequisite for a DHCP option space resource, but the
# provider exposes no such resource and "DHCP" is the only space defined on the
# grid, so this case covers the one value that can be asserted.
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
