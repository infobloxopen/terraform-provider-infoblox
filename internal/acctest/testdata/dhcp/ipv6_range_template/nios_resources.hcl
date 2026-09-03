# Auto-generated resource acceptance-test cases for Ipv6rangetemplate.
// Objects to be present on grid to run the tests
// - logic_filter_rules(ipv6_option_filter, ipv6_option_filter1), option_filter_rules (ipv6_option_filter)
// cloud_api_compatible is made true to pass all acceptance tests

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.name"                    = "{{random}}"
      "nios.cloud_api_compatible"    = "true"
      "nios.recycle_leases"          = "true"
      "nios.server_association_type" = "NONE"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
  }

}

// The testcase will fail, as this is a known issue
// If the user is a cloud-user, then they need Terraform internal ID with cloud permission and enable cloud delegation for the user to create a range template.
// if the user is a non cloud-user, they need to have  Terraform internal ID without cloud permission.
case "cloud_api_compatible" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping this test as it is a known issue."
  parallel    = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.cloud_api_compatible" = "true"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = false
    }
    check = {
      "nios.cloud_api_compatible" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      comment              = "example comment"
    }
    check = {
      "nios.comment" = "example comment"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      comment              = "example comment updated"
    }
    check = {
      "nios.comment" = "example comment updated"
    }
  }

}

case "delegated_member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      delegated_member     = { name = "{{grid_master_hostname}}" }
    }
    check = {
      "nios.delegated_member.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      delegated_member     = { name = "{{grid_member_hostname}}" }
    }
    check = {
      "nios.delegated_member.name" = "{{grid_member_hostname}}"
    }
  }

}

case "exclude" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      exclude              = [{ number_of_addresses = 10, offset = 20 }]
    }
    check = {
      "nios.exclude.#"                     = "1"
      "nios.exclude.0.number_of_addresses" = "10"
      "nios.exclude.0.offset"              = "20"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      exclude              = [{ number_of_addresses = 15, offset = 25, comment = "exclude for range template" }]
    }
    check = {
      "nios.exclude.#"                     = "1"
      "nios.exclude.0.number_of_addresses" = "15"
      "nios.exclude.0.offset"              = "25"
      "nios.exclude.0.comment"             = "exclude for range template"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      logic_filter_rules   = [{ filter = "ipv6_option_filter", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      logic_filter_rules   = [{ filter = "ipv6_option_filter1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
      member                  = { name = "{{grid_master_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.member.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
      member                  = { name = "{{grid_member_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.member.name" = "{{grid_member_hostname}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                 = "{{random2}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "number_of_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.number_of_addresses" = "100"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 150
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.number_of_addresses" = "150"
    }
  }

}

case "offset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 200
      cloud_api_compatible = true
    }
    check = {
      "nios.offset" = "200"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 250
      cloud_api_compatible = true
    }
    check = {
      "nios.offset" = "250"
    }
  }

}

case "option_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      option_filter_rules  = [{ filter = "ipv6_option_filter", permission = "Allow" }]
    }
    check = {
      "nios.option_filter_rules.#"            = "1"
      "nios.option_filter_rules.0.filter"     = "ipv6_option_filter"
      "nios.option_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      option_filter_rules  = [{ filter = "ipv6_option_filter", permission = "Deny" }]
    }
    check = {
      "nios.option_filter_rules.#"            = "1"
      "nios.option_filter_rules.0.filter"     = "ipv6_option_filter"
      "nios.option_filter_rules.0.permission" = "Deny"
    }
  }

}

case "recycle_leases" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      recycle_leases       = false
    }
    check = {
      "nios.recycle_leases" = "false"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      recycle_leases       = true
    }
    check = {
      "nios.recycle_leases" = "true"
    }
  }

}

case "server_association_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
      server_association_type = "MEMBER"
      member                  = { name = "{{grid_master_hostname}}" }
    }
    check = {
      "nios.server_association_type" = "MEMBER"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
      server_association_type = "NONE"
    }
    check = {
      "nios.server_association_type" = "NONE"
    }
  }

}
