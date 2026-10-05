# Auto-generated resource acceptance-test cases for RirOrganization.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.ext_attrs.RIPE Admin Contact"     = "ib-contact"
      "nios.ext_attrs.RIPE Country"           = "United Kingdom (GB)"
      "nios.ext_attrs.RIPE Technical Contact" = "TEST123-IB"
      "nios.ext_attrs.RIPE Email"             = "support@infoblox.com"
      "nios.id"                               = "ORG-CB{{random_int}}-IBTEST"
      "nios.maintainer"                       = "infoblox"
      "nios.name"                             = "{{random}}"
      "nios.sender_email"                     = "support@infoblox.com"
      "nios.rir"                              = "RIPE"
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
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
  }

}

case "import" {
  backend       = "nios"
  parallel      = true
  import        = true
  import_ignore = ["nios.password", "nios.ext_attrs", "nios.ext_attrs_all"]

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
        "RIPE Remarks"           = "Example RIR Organization"
        "RIPE Organization Type" = "IANA"
        "RIPE Notify"            = "support@infoblox.com"
      }
    }
    check = {
      "nios.ext_attrs.RIPE Admin Contact"     = "ib-contact"
      "nios.ext_attrs.RIPE Country"           = "United Kingdom (GB)"
      "nios.ext_attrs.RIPE Technical Contact" = "TEST123-IB"
      "nios.ext_attrs.RIPE Email"             = "support@infoblox.com"
      "nios.ext_attrs.RIPE Remarks"           = "Example RIR Organization"
      "nios.ext_attrs.RIPE Organization Type" = "IANA"
      "nios.ext_attrs.RIPE Notify"            = "support@infoblox.com"
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
        "RIPE Remarks"           = "Example Updated RIR Organization"
        "RIPE Organization Type" = "OTHER"
      }
    }
    check = {
      "nios.ext_attrs.RIPE Admin Contact"     = "ib-contact"
      "nios.ext_attrs.RIPE Country"           = "United Kingdom (GB)"
      "nios.ext_attrs.RIPE Technical Contact" = "TEST123-IB"
      "nios.ext_attrs.RIPE Email"             = "support@infoblox.com"
      "nios.ext_attrs.RIPE Remarks"           = "Example Updated RIR Organization"
      "nios.ext_attrs.RIPE Organization Type" = "OTHER"
    }
  }

}

case "id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.id" = "ORG-CB{{random_int}}-IBTEST"
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int2}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.id" = "ORG-CB{{random_int2}}-IBTEST"
    }
  }

}

case "maintainer" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.maintainer" = "infoblox"
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "nios-support"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.maintainer" = "nios-support"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random2}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "password" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass2"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
  }

}

case "rir" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.rir" = "RIPE"
    }
  }

}

case "sender_email" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.sender_email" = "support@infoblox.com"
    }
  }

  step {
    nios {
      id           = "ORG-CB{{random_int}}-IBTEST"
      maintainer   = "infoblox"
      name         = "{{random}}"
      password     = "test-pass"
      rir          = "RIPE"
      sender_email = "support2@infoblox.com"
      ext_attrs = {
        "RIPE Admin Contact"     = "ib-contact"
        "RIPE Country"           = "United Kingdom (GB)"
        "RIPE Technical Contact" = "TEST123-IB"
        "RIPE Email"             = "support@infoblox.com"
      }
    }
    check = {
      "nios.sender_email" = "support2@infoblox.com"
    }
  }

}
