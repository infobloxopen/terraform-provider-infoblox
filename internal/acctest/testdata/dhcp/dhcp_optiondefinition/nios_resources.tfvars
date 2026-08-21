# Auto-generated resource acceptance-test cases for DhcpOptiondefinition.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.code" = "10"
      "nios.name" = "{{random}}"
      "nios.type" = "string"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
  }

}

case "code" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.code" = "10"
    }
  }

  step {
    nios {
      code  = 20
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.code" = "20"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random2}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "space" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test1" {
    nios = {
      name = "{{random2}}"
    }
  }

  resource "infoblox_dhcp_optionspace" "test2" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test1.nios.name
    }
    check = {
      "nios.space" = "{{random2}}"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test2.nios.name
    }
    check = {
      "nios.space" = "{{random3}}"
    }
  }

}

case "type" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "string"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "ip-address"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "ip-address"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "8-bit unsigned integer"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "8-bit unsigned integer"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "domain-name"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "domain-name"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "array of ip-address"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "array of ip-address"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "boolean"
      space = infoblox_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "boolean"
    }
  }

}
