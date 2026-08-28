# Auto-generated resource acceptance-test cases for Address.
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
  }

}

case "address" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      address = "10.0.0.5"
      space   = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.address" = "10.0.0.5"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      comment = "some comment"
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      comment = "updated comment"
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "hwaddr" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      hwaddr  = "00:11:22:33:44:55"
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.hwaddr" = "00:11:22:33:44:55"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      hwaddr  = "55:44:33:22:11:00"
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.hwaddr" = "55:44:33:22:11:00"
    }
  }

}

case "interface" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address   = "10.0.0.1"
      space     = infoblox_ip_space.test.id
      interface = "eth0"
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.interface" = "eth0"
    }
  }

  step {
    uddi {
      address   = "10.0.0.1"
      space     = infoblox_ip_space.test.id
      interface = "eth1"
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.interface" = "eth1"
    }
  }

}

case "names" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.names.0.name" = "name1"
      "uddi.names.0.type" = "user"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
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
  resource "infoblox_ip_space" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "one" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.one.id
    }
  }
  resource "infoblox_ip_space" "two" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_subnet" "two" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.two.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.one.id
    }
    depends_on = [infoblox_subnet.one, infoblox_subnet.two]
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.two.id
    }
    depends_on = [infoblox_subnet.one, infoblox_subnet.two]
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      tags    = { tag1 = "value1", tag2 = "value2" }
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      address = "10.0.0.1"
      space   = infoblox_ip_space.test.id
      tags    = { tag2 = "value2changed", tag3 = "value3" }
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "next_available_id_count" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_ip_space.test.id
    }
  }

}

case "next_available_subnet" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_ip_space.test.id
    }
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      space = infoblox_ip_space.test.id
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
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_address_block" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 26
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      space = infoblox_ip_space.test.id
    }
    depends_on = [infoblox_subnet.test]
    check = {
      "uddi.address" = "12.0.0.1"
    }
  }

}

case "next_available_range" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ip_space" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_subnet" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_ip_space.test.id
    }
  }
  resource "infoblox_range" "test" {
    uddi = {
      start = "10.0.0.10"
      end = "10.0.0.20"
      space = infoblox_ip_space.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_ip_space.test.id
    }
    check = {
      "uddi.address" = "10.0.0.10"
    }
  }

  step {
    uddi {
      space = infoblox_ip_space.test.id
    }
    check = {
      "uddi.address" = "10.0.0.16"
    }
  }

}
