# Auto-generated resource acceptance-test cases for RecordA (nios).
case "basic" {
  # basic — generated from terraform-provider-nios
  backend = "nios"

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
  # disappears — generated from terraform-provider-nios
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true

  step {
    nios {
      name     = "{{random}}.example.com"
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
  }

}

case "comment" {
  # comment — generated from terraform-provider-nios
  backend = "nios"

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
  # creator — generated from terraform-provider-nios
  backend = "nios"

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
  # ddns_principal — generated from terraform-provider-nios
  backend = "nios"

  step {
    nios {
      name           = "{{random}}.example.com"
      ipv4addr       = "10.0.0.20"
      view           = "default"
      creator        = "DYNAMIC"
      ddns_principal = "DDNS_PRINCIPAL_REPLACE_ME"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_REPLACE_ME"
    }
  }

  step {
    nios {
      name           = "{{random}}.example.com"
      ipv4addr       = "10.0.0.20"
      view           = "default"
      creator        = "DYNAMIC"
      ddns_principal = "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"
    }
  }

}

case "ddns_protected" {
  # ddns_protected — generated from terraform-provider-nios
  backend = "nios"

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
  # disable — generated from terraform-provider-nios
  backend = "nios"

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
  # extattrs — generated from terraform-provider-nios
  backend = "nios"

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
  # forbid_reclamation — generated from terraform-provider-nios
  backend = "nios"

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
  # func_call — generated from terraform-provider-nios
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network" "test_func_call" {
    nios = {
      network = "85.85.0.0/16"
      network_view = "default"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random}}.example.com"
      view               = "default"
      dynamic_allocation = { network = infoblox_network.test.nios.network, network_view = "default" }
      comment            = "Original Function Call"
    }
    depends_on = [infoblox_network.test_func_call]
  }

  step {
    nios {
      name               = "{{random}}.example.com"
      view               = "default"
      dynamic_allocation = { network = infoblox_network.test.nios.network, network_view = "default" }
      comment            = "Function Call with Update"
    }
    depends_on = [infoblox_network.test_func_call]
  }

}

case "ipv4addr" {
  # ipv4addr — generated from terraform-provider-nios
  backend = "nios"

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
  # name — generated from terraform-provider-nios
  backend = "nios"

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
  # ttl — generated from terraform-provider-nios
  backend = "nios"

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
