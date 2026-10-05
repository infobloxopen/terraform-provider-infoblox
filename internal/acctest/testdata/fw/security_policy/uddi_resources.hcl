# Auto-generated resource acceptance-test cases for SecurityPolicy.
case "basic" {
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name"                  = "{{random}}"
      "uddi.description"           = ""
      "uddi.default_action"        = "action_allow"
      "uddi.default_redirect_name" = ""
      "uddi.ecs"                   = "false"
      "uddi.onprem_resolve"        = "false"
      "uddi.safe_search"           = "false"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  skip                  = true
  skip_reason           = "t.Skip: Test Skipped due to inconsistent error codes returned by the API [TDDFW-397]"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      description = "TEST_DESCRIPTION"
    }
    check = {
      "uddi.description" = "TEST_DESCRIPTION"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      description = "TEST_DESCRIPTION_UPDATE"
    }
    check = {
      "uddi.description" = "TEST_DESCRIPTION_UPDATE"
    }
  }

}

case "access_codes" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_named_list" "nl_test" {
    uddi = {
      name            = "{{random4}}"
      type            = "custom_list"
      items_described = [{ item = "tf-domain.com", description = "Example Domain" }]
    }
  }
  resource "infoblox_access_code" "ac_test1" {
    uddi = {
      name       = "{{random2}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ data = infoblox_named_list.nl_test.uddi.name, type = infoblox_named_list.nl_test.uddi.type }]
    }
  }
  resource "infoblox_access_code" "ac_test2" {
    uddi = {
      name       = "{{random3}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ data = infoblox_named_list.nl_test.uddi.name, type = infoblox_named_list.nl_test.uddi.type }]
    }
  }
  PREREQ

  step {
    uddi {
      name         = "{{random}}"
      access_codes = [infoblox_access_code.ac_test1.uddi.access_key]
    }
    check = {
      "uddi.access_codes.#" = "1"
    }
    check_pair = {
      "uddi.access_codes.0" = infoblox_access_code.ac_test1.uddi.access_key
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      access_codes = [infoblox_access_code.ac_test1.uddi.access_key, infoblox_access_code.ac_test2.uddi.access_key]
    }
    check = {
      "uddi.access_codes.#" = "2"
    }
    check_pair = {
      "uddi.access_codes.0" = infoblox_access_code.ac_test1.uddi.access_key
      "uddi.access_codes.1" = infoblox_access_code.ac_test2.uddi.access_key
    }
  }

}

case "default_action" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name           = "{{random}}"
      default_action = "action_allow"
    }
    check = {
      "uddi.default_action" = "action_allow"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      default_action = "action_redirect"
    }
    check = {
      "uddi.default_action" = "action_redirect"
    }
  }

}

case "default_redirect_name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_custom_redirect" "test_a" {
    uddi = {
      name = "{{random2}}"
      data = "156.2.3.10"
    }
  }
  resource "infoblox_custom_redirect" "test_b" {
    uddi = {
      name = "{{random3}}"
      data = "192.2.3.10"
    }
  }
  PREREQ

  step {
    uddi {
      name                  = "{{random}}"
      default_action        = "action_redirect"
      default_redirect_name = infoblox_custom_redirect.test_a.uddi.name
    }
    check = {
      "uddi.default_redirect_name" = "{{random2}}"
    }
  }

  step {
    uddi {
      name                  = "{{random}}"
      default_action        = "action_redirect"
      default_redirect_name = infoblox_custom_redirect.test_b.uddi.name
    }
    check = {
      "uddi.default_redirect_name" = "{{random3}}"
    }
  }

}

case "ecs" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      ecs  = true
    }
    check = {
      "uddi.ecs" = "true"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      ecs  = false
    }
    check = {
      "uddi.ecs" = "false"
    }
  }

}

case "network_lists" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_list" "nl_test1" {
    uddi = {
      name       = "{{random2}}"
      addr_block = [{ address = "{{random_cidr_network4}}" }]
    }
  }
  resource "infoblox_network_list" "nl_test2" {
    uddi = {
      name       = "{{random3}}"
      addr_block = [{ address = "{{random_cidr_network5}}" }]
    }
  }
  PREREQ

  step {
    uddi {
      name          = "{{random}}"
      network_lists = [infoblox_network_list.nl_test1.id]
    }
    check = {
      "uddi.network_lists.#" = "1"
    }
    check_pair = {
      "uddi.network_lists.0" = infoblox_network_list.nl_test1.id
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      network_lists = [infoblox_network_list.nl_test1.id, infoblox_network_list.nl_test2.id]
    }
    check = {
      "uddi.network_lists.#" = "2"
    }
    check_pair = {
      "uddi.network_lists.0" = infoblox_network_list.nl_test1.id
      "uddi.network_lists.1" = infoblox_network_list.nl_test2.id
    }
  }

}

case "onprem_resolve" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name           = "{{random}}"
      onprem_resolve = true
    }
    check = {
      "uddi.onprem_resolve" = "true"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      onprem_resolve = false
    }
    check = {
      "uddi.onprem_resolve" = "false"
    }
  }

}

case "precedence" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      precedence = 1
    }
    check = {
      "uddi.precedence" = "1"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      precedence = 2
    }
    check = {
      "uddi.precedence" = "2"
    }
  }

}

case "rules" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_named_list" "nl_test1" {
    uddi = {
      name            = "{{random2}}"
      type            = "custom_list"
      items_described = [{ item = "tf1-domain.com", description = "Example Domain" }]
    }
  }
  resource "infoblox_named_list" "nl_test2" {
    uddi = {
      name            = "{{random3}}"
      type            = "custom_list"
      items_described = [{ item = "tf2-domain.com", description = "Example Domain" }]
    }
  }
  PREREQ
  parallel = true

  step {
    uddi {
      name  = "{{random}}"
      rules = [{ action = "action_allow", data = infoblox_named_list.nl_test1.uddi.name, type = infoblox_named_list.nl_test1.uddi.type }]
    }
    check = {
      "uddi.rules.0.action" = "action_allow"
      "uddi.rules.0.type"   = "custom_list"
    }
  }

  step {
    uddi {
      name  = "{{random}}"
      rules = [{ action = "action_block", data = infoblox_named_list.nl_test2.uddi.name, type = infoblox_named_list.nl_test2.uddi.type }]
    }
    check = {
      "uddi.rules.0.action" = "action_block"
      "uddi.rules.0.type"   = "custom_list"
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.rules.#" = "0"
    }
  }

}

case "rules_all_attributes" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_named_list" "nl_test" {
    uddi = {
      name            = "{{random2}}"
      type            = "custom_list"
      items_described = [{ item = "tf-domain.com", description = "Example Domain" }]
    }
  }
  resource "infoblox_custom_redirect" "redirect_test" {
    uddi = {
      name = "{{random3}}"
      data = "192.0.2.10"
    }
  }
  PREREQ
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      rules = [{
        action        = "action_redirect"
        data          = infoblox_named_list.nl_test.uddi.name
        type          = infoblox_named_list.nl_test.uddi.type
        redirect_name = infoblox_custom_redirect.redirect_test.uddi.name
      }]
    }
    check = {
      "uddi.rules.0.action"        = "action_redirect"
      "uddi.rules.0.data"          = "{{random2}}"
      "uddi.rules.0.type"          = "custom_list"
      "uddi.rules.0.redirect_name" = "{{random3}}"
    }
    # policy_id is server-assigned to this policy's id; list_id is server-assigned from the named list
    check_pair = {
      "uddi.rules.0.policy_id" = infoblox_security_policy.test.id
      "uddi.rules.0.list_id"   = infoblox_named_list.nl_test.uddi.id
    }
  }

}

# TODO: add prerequisite to dynamically create a DFP resource once DFP service support is added.
case "dfps" {
  backend     = "uddi"
  skip        = true
  skip_reason = "hardcoded DFP ID — requires prerequisite once DFP service support is added"
  parallel    = true

  step {
    uddi {
      name = "{{random}}"
      dfps = [530499]
    }
    check = {
      "uddi.dfps.0" = "530499"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      dfps = []
    }
    check = {
      "uddi.dfps.#" = "0"
    }
  }

}

case "safe_search" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      safe_search = true
    }
    check = {
      "uddi.safe_search" = "true"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      safe_search = false
    }
    check = {
      "uddi.safe_search" = "false"
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
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
