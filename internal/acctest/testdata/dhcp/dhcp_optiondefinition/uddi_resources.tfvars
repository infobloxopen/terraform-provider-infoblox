# Auto-generated resource acceptance-test cases for DhcpOptiondefinition.
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type         = "boolean"
    }
  }

}

case "array" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
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
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type         = "boolean"
      array        = false
    }
    check = {
      "uddi.array" = "false"
    }
  }

}

case "code" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
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
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type         = "boolean"
    }
    check = {
      "uddi.code" = "235"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
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
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type         = "boolean"
      comment      = "boolean option code type update"
    }
    check = {
      "uddi.comment" = "boolean option code type update"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
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
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type         = "boolean"
    }
    check = {
      "uddi.name" = "basic_opt_code_1"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: config helper 'testAccOptionCodeOptionSpace' could not be parsed (no resource block found)
case "option_space" {
  backend     = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccOptionCodeOptionSpace' could not be parsed (no resource block found)"
}

case "type" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_option_space_unknown" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      code         = 234
      name         = "basic_opt_code"
      option_space = infoblox_dhcp_option_space_unknown.test.id
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
      option_space = infoblox_dhcp_option_space_unknown.test.id
      type         = "int16"
    }
    check = {
      "uddi.type" = "int16"
    }
  }

}
