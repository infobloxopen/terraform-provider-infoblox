# Auto-generated resource acceptance-test cases for Ipv6DhcpOptiondefinition.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.code"  = "10"
      "nios.name"  = "{{random}}"
      "nios.type"  = "string"
      "nios.space" = "{{random2}}"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
  }

}

case "code" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
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
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
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
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random3}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
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
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
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
  resource "infoblox_ipv6_dhcp_optionspace" "test1" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  resource "infoblox_ipv6_dhcp_optionspace" "test2" {
    nios = {
      name = "{{random3}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test1.nios.name
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
      space = infoblox_ipv6_dhcp_optionspace.test2.nios.name
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
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    nios = {
      name = "{{random2}}"
      enterprise_number = 10
    }
  }
  PREREQ

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "string"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "string"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "boolean"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "boolean"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "8-bit unsigned integer"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "8-bit unsigned integer"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "ip-address"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "ip-address"
    }
  }

  step {
    nios {
      code  = 10
      name  = "{{random}}"
      type  = "array of 8-bit integer"
      space = infoblox_ipv6_dhcp_optionspace.test.nios.name
    }
    check = {
      "nios.type" = "array of 8-bit integer"
    }
  }

}
