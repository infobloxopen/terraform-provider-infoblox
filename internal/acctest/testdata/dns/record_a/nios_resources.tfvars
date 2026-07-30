# Auto-generated resource acceptance-test cases for RecordA.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
    check = {
      "nios.ipv4addr"           = "10.0.0.20"
      "nios.name"               = "{{random}}.example.com"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
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
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      comment  = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      comment  = "This is an updated record"
    }
    check = {
      "nios.comment" = "This is an updated record"
    }
  }

}

case "creator" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      creator  = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}.example.com"
      ipv4addr       = "10.0.0.20"
      view           = "default"
      creator        = "DYNAMIC"
      ddns_principal = "DDNS_PRINCIPAL_1"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_1"
    }
  }

  step {
    nios {
      name           = "{{random}}.example.com"
      ipv4addr       = "10.0.0.20"
      view           = "default"
      creator        = "DYNAMIC"
      ddns_principal = "DDNS_PRINCIPAL_2"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_2"
    }
  }

}

case "ddns_protected" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}.example.com"
      ipv4addr       = "10.0.0.20"
      view           = "default"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random}}.example.com"
      ipv4addr       = "10.0.0.20"
      view           = "default"
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "extattrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}.example.com"
      ipv4addr  = "10.0.0.20"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}.example.com"
      ipv4addr  = "10.0.0.20"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "forbid_reclamation" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}.example.com"
      ipv4addr           = "10.0.0.20"
      view               = "default"
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "{{random}}.example.com"
      ipv4addr           = "10.0.0.20"
      view               = "default"
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "func_call" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_network" "test_func_call" {
  #   nios = {
  #     network = "85.85.0.0/16"
  #     network_view = "default"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name = "{{random}}.example.com"
      view = "default"
      # dynamic_allocation = { network = infoblox_network.test.nios.network, network_view = "default" }
      dynamic_allocation = { network = "12.0.0.0/24", network_view = "default" }
      comment            = "Original Function Call"
    }
    # depends_on = [infoblox_network.test_func_call]
  }

  step {
    nios {
      name = "{{random}}.example.com"
      view = "default"
      # dynamic_allocation = { network = infoblox_network.test.nios.network, network_view = "default" }
      dynamic_allocation = { network = "12.0.0.0/24", network_view = "default" }
      comment            = "Function Call with Update"
    }
    # depends_on = [infoblox_network.test_func_call]
  }

}

case "ipv4addr" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "10.0.0.20"
    }
  }

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.1.0.20"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "10.1.0.20"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
    check = {
      "nios.name" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      name     = "{{random2}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
    check = {
      "nios.name" = "{{random2}}.example.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
      ttl      = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}
