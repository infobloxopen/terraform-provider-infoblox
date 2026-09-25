# Auto-generated resource acceptance-test cases for OptionGroup.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
    }
    check = {
      "uddi.name"     = "{{random}}"
      "uddi.protocol" = "ip4"
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
      name     = "{{random}}"
      protocol = "ip4"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
      comment  = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
      comment  = "test comment update"
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
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code = 234
      name = "test_dhcp_option_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type = "boolean"
    }
  }
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_option_group" "test_group" {
    uddi = {
      name     = "{{random3}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      name         = "{{random2}}"
      protocol     = "ip4"
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_optiondefinition.test.id, option_value = true }]
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      name         = "{{random2}}"
      protocol     = "ip4"
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_optiondefinition.test.id, option_value = false }]
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.option_value" = "false"
    }
  }

  step {
    uddi {
      name         = "{{random2}}"
      protocol     = "ip4"
      dhcp_options = [{ type = "group", group = infoblox_option_group.test_group.id }]
    }
    check = {
      "uddi.dhcp_options.#"          = "1"
      "uddi.dhcp_options.0.type"     = "group"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name     = "option_group_test_1"
      protocol = "ip4"
    }
    check = {
      "uddi.name" = "option_group_test_1"
    }
  }

}

case "protocol" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
    }
    check = {
      "uddi.protocol" = "ip4"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
      tags     = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
      tags     = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
