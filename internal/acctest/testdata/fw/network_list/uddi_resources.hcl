# Auto-generated resource acceptance-test cases for NetworkList.
case "basic" {
  backend = "uddi"

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip}}/32" }]
    }
    check = {
      "uddi.name"                 = "{{random}}"
      "uddi.addr_block.0.address" = "{{random_public_ip}}/32"
      "uddi.description"          = ""
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip}}/32" }]
    }
  }

}

case "name" {
  backend = "uddi"

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip}}/32" }]
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name       = "{{random2}}"
      addr_block = [{ address = "{{random_public_ip}}/32" }]
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "description" {
  backend = "uddi"

  step {
    uddi {
      name        = "{{random}}"
      addr_block  = [{ address = "{{random_public_ip}}/32" }]
      description = "Test Description"
    }
    check = {
      "uddi.name"        = "{{random}}"
      "uddi.description" = "Test Description"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      addr_block  = [{ address = "{{random_public_ip}}/32" }]
      description = "Updated Test Description"
    }
    check = {
      "uddi.name"        = "{{random}}"
      "uddi.description" = "Updated Test Description"
    }
  }

}

case "addr_block" {
  backend = "uddi"

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip}}/32", description = "Test Address Block 1" }]
    }
    check = {
      "uddi.name"                     = "{{random}}"
      "uddi.addr_block.0.address"     = "{{random_public_ip}}/32"
      "uddi.addr_block.0.description" = "Test Address Block 1"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      addr_block = [{ address = "{{random_public_ip2}}/32", description = "Test Address Block 2" }]
    }
    check = {
      "uddi.name"                     = "{{random}}"
      "uddi.addr_block.0.address"     = "{{random_public_ip2}}/32"
      "uddi.addr_block.0.description" = "Test Address Block 2"
    }
  }

}

case "addr_block_multi" {
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
      addr_block = [
        { address = "{{random_public_ip}}/32", description = "Test Address Block 1" },
        { address = "{{random_public_ip2}}/32", description = "Test Address Block 2" },
      ]
    }
    check = {
      "uddi.name"                     = "{{random}}"
      "uddi.addr_block.#"             = "2"
      "uddi.addr_block.0.address"     = "{{random_public_ip}}/32"
      "uddi.addr_block.0.description" = "Test Address Block 1"
      "uddi.addr_block.1.address"     = "{{random_public_ip2}}/32"
      "uddi.addr_block.1.description" = "Test Address Block 2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      addr_block = [
        { address = "{{random_public_ip}}/32", description = "Test Address Block 1" },
        { address = "{{random_public_ip2}}/32", description = "Test Address Block 2" },
        { address = "{{random_public_ip3}}/32", description = "Test Address Block 3" },
      ]
    }
    check = {
      "uddi.name"                     = "{{random}}"
      "uddi.addr_block.#"             = "3"
      "uddi.addr_block.2.address"     = "{{random_public_ip3}}/32"
      "uddi.addr_block.2.description" = "Test Address Block 3"
    }
  }

}
