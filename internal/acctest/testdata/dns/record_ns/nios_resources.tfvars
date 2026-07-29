# Auto-generated resource acceptance-test cases for RecordNs.
case "basic" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = "addressesHCL"
      view       = "default"
    }
    check = {
      "nios.name"                        = "example.com"
      "nios.nameserver"                  = "{{random}}.example.com"
      "nios.addresses.0.address"         = "20.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "false"
      "nios.view"                        = "default"
      "nios.ms_delegation_name"          = ""
    }
  }

}

case "disappears" {
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true
  parallel = true

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = "addressesHCL"
      view       = "default"
    }
  }

}

case "addresses" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = "addressesHCL1"
      view       = "default"
    }
    check = {
      "nios.addresses.0.address"         = "20.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "false"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = "addressesHCL2"
      view       = "default"
    }
    check = {
      "nios.addresses.0.address"         = "40.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "false"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "ns1.example.com"
      addresses  = "addressesHCL3"
      view       = "default"
    }
    check = {
      "nios.addresses.0.address"         = "40.0.0.0"
      "nios.addresses.0.auto_create_ptr" = "true"
    }
  }

}

case "nameserver" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random}}.example.com"
      addresses  = "addressesHCL"
      view       = "default"
    }
    check = {
      "nios.nameserver" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      name       = "example.com"
      nameserver = "{{random2}}.example.com"
      addresses  = "addressesHCL"
      view       = "default"
    }
    check = {
      "nios.nameserver" = "{{random2}}.example.com"
    }
  }

}
