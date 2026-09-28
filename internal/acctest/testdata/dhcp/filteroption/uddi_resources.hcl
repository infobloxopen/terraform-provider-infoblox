# Filteroption — uddi resource cases
// An option code and a non-default option space have to be created before running the test cases.

case "basic" {
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_OPTION_CODE_1_ID"]
  skip_reason       = "UDDI_OPTION_CODE_1_ID environment variable must be set for this test to run"

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.name"       = "{{random}}"
      "uddi.role"       = "values"
      "uddi.lease_time" = "0"
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
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.comment" = "test comment update"
    }
  }

}

case "dhcp_options" {
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_OPTION_GROUP_1_ID"]
  skip_reason       = "UDDI_OPTION_GROUP_1_ID environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code = 234
      name = "{{random4}}"
      option_space = infoblox_dhcp_optionspace.test.id
      type = "boolean"
    }
  }

//   resource "infoblox_dhcp_option_group_unknown" "test" {
//       uddi = {
//         name = "\"og-\"+optionSpace"
//         protocol = "ip4"
//       }
//   }
  PREREQ

  step {
    uddi {
      name         = "{{random}}"
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_optiondefinition.test.id, option_value = "true" }]
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      dhcp_options = [{ type = "group", group = "{{uddi_option_group_1_id}}" }]
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.dhcp_options.#"       = "1"
      "uddi.dhcp_options.0.type"  = "group"
      "uddi.dhcp_options.0.group" = "{{uddi_option_group_1_id}}"
    }
  }

}

case "header_option_filename" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = "pxeboot.img"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.header_option_filename" = "pxeboot.img"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = "pxeboot-update.img"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.header_option_filename" = "pxeboot-update.img"
    }
  }

}

case "header_option_server_address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = "192.168.10.10"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.header_option_server_address" = "192.168.10.10"
    }
  }

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = "192.168.11.11"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.header_option_server_address" = "192.168.11.11"
    }
  }

}

case "header_option_server_name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = "tf-infoblox-test.com."
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.header_option_server_name" = "tf-infoblox-test.com."
    }
  }

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = "tf-infoblox.com."
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.header_option_server_name" = "tf-infoblox.com."
    }
  }

}

case "lease_time" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      lease_time = 600
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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

  step {
    uddi {
      name = "{{random}}"
      role = "values"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.rules.match"                = "any"
      "uddi.rules.rules.0.compare"      = "equals"
      "uddi.rules.rules.0.option_value" = "value1"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "all"
        rules = [{
          compare      = "not_equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value2"
        }]
      }
    }
    check = {
      "uddi.rules.match"                = "all"
      "uddi.rules.rules.0.compare"      = "not_equals"
      "uddi.rules.rules.0.option_value" = "value2"
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
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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
          option_code  = "{{uddi_option_code_1_id}}"
          option_value = "value1"
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
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code = 234
      name = "{{random4}}"
      option_space = infoblox_dhcp_optionspace.test.id
      type = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name                                = "{{random}}"
      vendor_specific_option_option_space = infoblox_dhcp_optionspace.test.id
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_dhcp_optiondefinition.test.id
          option_value = "value1"
        }]
      }
    }
    check = {
      "uddi.vendor_specific_option_option_space" = infoblox_dhcp_optionspace.test.id
    }
  }

}
