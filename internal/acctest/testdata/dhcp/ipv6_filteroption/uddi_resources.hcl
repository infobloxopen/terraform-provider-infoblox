# Ipv6filteroption — uddi resource cases
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.name"       = "{{random}}"
      "uddi.role"       = "values"
      "uddi.protocol"   = "ip6"
      "uddi.lease_time" = "0"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment update"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.comment" = "test comment update"
    }
  }

}

case "dhcp_options" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      dhcp_options = [{
        type         = "option"
        option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
        option_value = "true"
      }]
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      dhcp_options = [{
        type         = "option"
        option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
        option_value = "false"
      }]
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "false"
    }
  }

}

case "lease_time" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name       = "{{random}}"
      lease_time = 600
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.lease_time" = "600"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      lease_time = 1200
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.lease_time" = "1200"
    }
  }

}

case "role" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      role = "values"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.role" = "values"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      role = "selection"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.role" = "selection"
    }
  }

}

case "rules" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.rules.match"                = "any"
      "uddi.rules.rules.0.compare"      = "equals"
      "uddi.rules.rules.0.option_value" = "true"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "all"
        rules = [{
          compare      = "not_equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "false"
        }]
      }
    }
    check = {
      "uddi.rules.match"                = "all"
      "uddi.rules.rules.0.compare"      = "not_equals"
      "uddi.rules.rules.0.option_value" = "false"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
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
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "vendor_specific_option_option_space" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optionspace" "test2" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name                                = "{{random}}"
      vendor_specific_option_option_space = infoblox_ipv6_dhcp_optionspace.test.id
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

  step {
    uddi {
      name                                = "{{random}}"
      vendor_specific_option_option_space = infoblox_ipv6_dhcp_optionspace.test2.id
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
    }
  }

}
