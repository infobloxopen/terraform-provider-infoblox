# Auto-generated resource acceptance-test cases for Address.

# TODO: "next_available_range" case needs pre-created range - drop this once infoblox_range is onboarded.

case "basic" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.address" = "10.0.0.1"
    }
    check_pair = {
      "uddi.space" = infoblox_network_view.test.id
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network.test]
  }

}

case "address" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      address = "10.0.0.5"
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.address" = "10.0.0.5"
    }
  }

}

case "comment" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      comment = "some comment"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      comment = "updated comment"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "external_keys" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address       = "10.0.0.1"
      space         = infoblox_network_view.test.id
      external_keys = { key1 = "value1" }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.external_keys.key1" = "value1"
    }
  }

  step {
    uddi {
      address       = "10.0.0.1"
      space         = infoblox_network_view.test.id
      external_keys = { key1 = "value1changed", key2 = "value2" }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.external_keys.key1" = "value1changed"
      "uddi.external_keys.key2" = "value2"
    }
  }

}

case "hwaddr" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      hwaddr  = "00:11:22:33:44:55"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.hwaddr" = "00:11:22:33:44:55"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      hwaddr  = "55:44:33:22:11:00"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.hwaddr" = "55:44:33:22:11:00"
    }
  }

}

case "interface" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address   = "10.0.0.1"
      space     = infoblox_network_view.test.id
      interface = "eth0"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.interface" = "eth0"
    }
  }

  step {
    uddi {
      address   = "10.0.0.1"
      space     = infoblox_network_view.test.id
      interface = "eth1"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.interface" = "eth1"
    }
  }

}

case "names" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      names   = [{ name = "name1", type = "user" }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.names.0.name" = "name1"
      "uddi.names.0.type" = "user"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      names   = [{ name = "name2", type = "user" }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.names.0.name" = "name2"
      "uddi.names.0.type" = "user"
    }
  }

}

case "space" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "one" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.one.id
    }
  }
  resource "infoblox_network_view" "two" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_network" "two" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.two.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.one.id
    }
    depends_on = [infoblox_network.one, infoblox_network.two]
    check_pair = {
      "uddi.space" = infoblox_network_view.one.id
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.two.id
    }
    depends_on = [infoblox_network.one, infoblox_network.two]
    check_pair = {
      "uddi.space" = infoblox_network_view.two.id
    }
  }

}

case "tags" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      tags    = { tag1 = "value1", tag2 = "value2" }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.tags.tag1"     = "value1"
      "uddi.tags.tag2"     = "value2"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      tags    = { tag2 = "value2changed", tag3 = "value3" }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.tags.tag2"     = "value2changed"
      "uddi.tags.tag3"     = "value3"
    }
  }

}

case "next_available_subnet" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "one" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  resource "infoblox_network" "two" {
    uddi = {
      address = "12.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space              = infoblox_network_view.test.id
      dynamic_allocation = { next_available_id = infoblox_network.one.id }
    }
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      space              = infoblox_network_view.test.id
      dynamic_allocation = { next_available_id = infoblox_network.two.id }
    }
    check = {
      "uddi.address" = "12.0.0.1"
    }
  }

}

case "next_available_address_block" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network_container" "one" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  resource "infoblox_network" "one" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 26
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network_container.one]
  }
  resource "infoblox_network_container" "two" {
    uddi = {
      address = "12.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  resource "infoblox_network" "two" {
    uddi = {
      address = "12.0.0.0"
      cidr    = 26
      space   = infoblox_network_view.test.id
    }
    depends_on = [infoblox_network_container.two]
  }
  PREREQ

  step {
    uddi {
      space              = infoblox_network_view.test.id
      dynamic_allocation = { next_available_id = infoblox_network_container.one.id }
    }
    depends_on = [infoblox_network.one]
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      space              = infoblox_network_view.test.id
      dynamic_allocation = { next_available_id = infoblox_network_container.two.id }
    }
    depends_on = [infoblox_network.two]
    check = {
      "uddi.address" = "12.0.0.1"
    }
  }

}

case "next_available_range" {
  backend     = "uddi"
  parallel    = true

  step {
    uddi {
      space              = "ipam/ip_space/84c53c33-a2d7-11f1-a4fc-eecab8c1578d"
      dynamic_allocation = { next_available_id = "ipam/range/8e6ec141-a2d7-11f1-829e-02fb57fee572" }
    }
    check = {
      "uddi.address" = "10.0.0.10"
    }
  }

  step {
    uddi {
      space              = "ipam/ip_space/84c53c33-a2d7-11f1-a4fc-eecab8c1578d"
      dynamic_allocation = { next_available_id = "ipam/range/8f0bfc0e-a2d7-11f1-a4fc-eecab8c1578d" }
    }
    check = {
      "uddi.address" = "10.0.0.30"
    }
  }

}

case "next_available_id_count" {
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
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  resource "infoblox_address" "bulk" {
    count = 5
    uddi = {
      space              = infoblox_network_view.test.id
      dynamic_allocation = { next_available_id = infoblox_network.test.id }
    }
  }
  PREREQ

  step {
    uddi {
      space              = infoblox_network_view.test.id
      dynamic_allocation = { next_available_id = infoblox_network.test.id }
    }
    depends_on = [infoblox_address.bulk]
    check = {
      "uddi.address" = "10.0.0.6"
    }
    check_pair = {
      "uddi.space" = infoblox_network_view.test.id
    }
  }

}
