# TODO: The following prerequisites MUST exist on the CSP tenant before running these tests:
#   - named list : "tf-provider-test-access-code"   (type: custom_list)
#   - named list : "tf-provider-test-access-code-2" (type: custom_list)

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.name"         = "{{random}}"
      "uddi.activation"   = "2030-01-01T00:00:00Z"
      "uddi.expiration"   = "2031-01-01T00:00:00Z"
      "uddi.rules.0.type" = "custom_list"
      "uddi.rules.0.data" = "tf-provider-test-access-code"
    }
  }
}

case "disappears" {
  backend               = "uddi"
  skip                  = true
  skip_reason           = "t.Skip: Test Skipped due to inconsistent error codes returned by the API [TDDFW-397]"
  disappears            = true
  expect_non_empty_plan = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
  }
}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name       = "{{random2}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
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
      activation  = "2030-01-01T00:00:00Z"
      expiration  = "2031-01-01T00:00:00Z"
      description = "First description"
      rules       = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.description" = "First description"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      activation  = "2030-01-01T00:00:00Z"
      expiration  = "2031-01-01T00:00:00Z"
      description = "Updated description"
      rules       = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.description" = "Updated description"
    }
  }
}

case "activation" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-06-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.activation" = "2030-06-01T00:00:00Z"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-09-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.activation" = "2030-09-01T00:00:00Z"
    }
  }
}

case "expiration" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-06-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.expiration" = "2031-06-01T00:00:00Z"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2032-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.expiration" = "2032-01-01T00:00:00Z"
    }
  }
}

case "rules" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }]
    }
    check = {
      "uddi.rules.0.type" = "custom_list"
      "uddi.rules.0.data" = "tf-provider-test-access-code"
    }
  }
}

case "multiple_rules" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules = [
        { type = "custom_list", data = "tf-provider-test-access-code" },
        { type = "custom_list", data = "tf-provider-test-access-code-2" },
      ]
    }
    check = {
      "uddi.rules.0.type" = "custom_list"
      "uddi.rules.0.data" = "tf-provider-test-access-code"
      "uddi.rules.1.type" = "custom_list"
      "uddi.rules.1.data" = "tf-provider-test-access-code-2"
    }
  }
}

case "rules_order" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules = [
        { type = "custom_list", data = "tf-provider-test-access-code" },
        { type = "custom_list", data = "tf-provider-test-access-code-2" },
      ]
    }
    check = {
      "uddi.rules.0.data" = "tf-provider-test-access-code"
      "uddi.rules.1.data" = "tf-provider-test-access-code-2"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules = [
        { type = "custom_list", data = "tf-provider-test-access-code-2" },
        { type = "custom_list", data = "tf-provider-test-access-code" },
      ]
    }
    check = {
      "uddi.rules.0.data" = "tf-provider-test-access-code-2"
      "uddi.rules.1.data" = "tf-provider-test-access-code"
    }
  }
}
