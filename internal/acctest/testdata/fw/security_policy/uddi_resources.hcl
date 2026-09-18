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

# TODO: The following prerequisites MUST exist on the portal before running these tests:
#   - named_list : tf-test-named-list  (id: 1752596, type: custom_list)
# TODO: update prerequisites_hcl to use infoblox_named_list once PR #625 is merged
case "access_codes" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_access_code" "ac_test1" {
    uddi = {
      name       = "{{random2}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ data = "tf-test-named-list", type = "custom_list" }]
    }
  }
  resource "infoblox_access_code" "ac_test2" {
    uddi = {
      name       = "{{random3}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ data = "tf-test-named-list", type = "custom_list" }]
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
  }

  step {
    uddi {
      name         = "{{random}}"
      access_codes = [infoblox_access_code.ac_test1.uddi.access_key, infoblox_access_code.ac_test2.uddi.access_key]
    }
    check = {
      "uddi.access_codes.#" = "2"
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

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_td_custom_redirect not yet implemented
case "default_redirect_name" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_td_custom_redirect not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_td_custom_redirect_unknown" "test_a" {
    uddi = {
      name = "{{random2}}"
      data = "156.2.3.10"
    }
  }
  resource "infoblox_td_custom_redirect_unknown" "test_b" {
    uddi = {
      name = "{{random3}}"
      data = "192.2.3.10"
    }
  }
  PREREQ

  step {
    uddi {
      name           = "{{random}}"
      default_action = "action_redirect"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      default_action = "action_redirect"
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

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_td_network_list not yet implemented
case "network_lists" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_td_network_list not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_td_network_list_unknown" "nl_test1" {
    uddi = {
      name = "{{random2}}"
      items = ["{{random4}}/32"]
    }
  }
  resource "infoblox_td_network_list_unknown" "nl_test2" {
    uddi = {
      name = "{{random3}}"
      items = ["{{random5}}/32"]
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random}}"
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

# TODO: unskip once infoblox_named_list is merged (PR #625)
case "rules" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_named_list not yet registered (pending PR #625)"
  parallel    = true
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

}

case "dfps" {
  backend  = "uddi"
  parallel = true

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
