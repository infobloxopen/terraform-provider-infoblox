# Auto-generated datasource acceptance-test cases for InfraService.
#  TODO: Objects to be present in the grid for testing
#   - infra/pool : infra/pool/ch4sbtxrcat2wkppegumcuizdyll5nur

case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.description", "uddi.desired_state", "uddi.name", "uddi.pool_id", "uddi.service_type"]

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = "infra/pool/ch4sbtxrcat2wkppegumcuizdyll5nur"
      service_type = "dns"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      env = "uddi.tags_all.env"
    }
  }

  pair_checks = ["uddi.description", "uddi.desired_state", "uddi.name", "uddi.pool_id", "uddi.service_type"]

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = "infra/pool/ch4sbtxrcat2wkppegumcuizdyll5nur"
      service_type = "dns"
      tags         = { env = "{{random2}}" }
    }
  }

}
