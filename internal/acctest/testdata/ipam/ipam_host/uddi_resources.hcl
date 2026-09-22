# Auto-generated resource acceptance-test cases for IpamHost.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
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
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment updated"
    }
    check = {
      "uddi.comment" = "test comment updated"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "addresses" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

//   step {
//     uddi {
//         name = "{{random}}"
//       addresses = [{ next_available_id = infoblox_network.test.id }]
//     }
//     check = {
//       "uddi.addresses.#"         = "1"
//       "uddi.addresses.0.address" = "10.0.0.1"
//     }
//   }

//   step {
//     uddi {
//       addresses = [{ next_available_id = infoblox_network.test.id }, { next_available_id = infoblox_network.test1.id }, { next_available_id = infoblox_network.test2.id }]
//     }
//     check = {
//       "uddi.addresses.#"         = "3"
//       "uddi.addresses.0.address" = "10.0.0.1"
//       "uddi.addresses.1.address" = "192.168.1.1"
//       "uddi.addresses.2.address" = "10.0.0.1"
//     }
//   }

  step {
    uddi {
        name = "{{random}}"
      addresses = [{ address = "10.0.0.1", space = infoblox_network_view.test.id }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.addresses.#"         = "1"
      "uddi.addresses.0.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      addresses = [{ address = "10.0.0.2", space = infoblox_network_view.test.id }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.addresses.#"         = "1"
      "uddi.addresses.0.address" = "10.0.0.2"
    }
  }

}

case "addresses_next_available_id_count" {
  backend  = "uddi"
  parallel = true
  skip = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      name      = "host-$${count.index}"
      addresses = [{ next_available_id = infoblox_network.test.id }]
    }
  }

}
