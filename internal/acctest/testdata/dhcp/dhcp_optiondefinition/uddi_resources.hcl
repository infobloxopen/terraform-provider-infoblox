# Auto-generated resource acceptance-test cases for DhcpOptiondefinition.
case "basic" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
    check = {
      "uddi.code"  = "234"
      "uddi.name"  = "basic_opt_code"
      "uddi.type"  = "boolean"
      "uddi.array" = "false"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }

}

case "array" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
      array        = true
    }
    check = {
      "uddi.array" = "true"
    }
  }

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
      array        = false
    }
    check = {
      "uddi.array" = "false"
    }
  }

}

case "code" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
    check = {
      "uddi.code" = "234"
    }
  }

  step {
    uddi {
      code         = 235
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
    check = {
      "uddi.code" = "235"
    }
  }

}

case "comment" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
      comment      = "boolean option code type"
    }
    check = {
      "uddi.comment" = "boolean option code type"
    }
  }

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
      comment      = "boolean option code type update"
    }
    check = {
      "uddi.comment" = "boolean option code type update"
    }
  }

}

case "name" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
    check = {
      "uddi.name" = "basic_opt_code"
    }
  }

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code_1"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
    check = {
      "uddi.name" = "basic_opt_code_1"
    }
  }

}

case "option_space" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test1" {
    uddi = {
      name = "{{random}}"
    }
  }

  resource "infoblox_dhcp_optionspace" "test2" {
    uddi = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code_1"
      option_space = infoblox_dhcp_optionspace.test1.id
      type         = "boolean"
    }
    check = {
      "uddi.name"         = "basic_opt_code_1"
    }
    check_pair = {
      "uddi.option_space" = infoblox_dhcp_optionspace.test1.id
    }
  }

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code_1"
      option_space = infoblox_dhcp_optionspace.test2.id
      type         = "boolean"
    }
    check = {
      "uddi.name"         = "basic_opt_code_1"
    }
    check_pair = {
      "uddi.option_space" = infoblox_dhcp_optionspace.test2.id
    }
  }

}

case "type" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "boolean"
    }
    check = {
      "uddi.type" = "boolean"
    }
  }

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type         = "int16"
    }
    check = {
      "uddi.type" = "int16"
    }
  }

}
