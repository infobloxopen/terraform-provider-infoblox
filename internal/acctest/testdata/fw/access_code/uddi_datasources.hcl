# TODO: The following prerequisites MUST exist on the CSP tenant before running these tests:
#   - named list : "tf-provider-test-access-code"  (type: custom_list)

case "filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.access_key", "uddi.activation", "uddi.description", "uddi.expiration", "uddi.name"]

  step {
    uddi {
      name       = "{{random}}"
      activation = "2030-01-01T00:00:00Z"
      expiration = "2031-01-01T00:00:00Z"
      rules      = [{ type = "custom_list", data = "tf-provider-test-access-code" }] # TODO: hardcoded named list prerequisite (see file header)
    }
  }
}
