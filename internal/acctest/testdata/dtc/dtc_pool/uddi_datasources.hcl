# Auto-generated datasource acceptance-test cases for DtcPool (uddi).
case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.name", "uddi.method", "uddi.comment", "uddi.disabled", "uddi.pool_availability", "uddi.ttl"]

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type = "tag_filters"
    values = {
      Site = "uddi.tags.Site"
    }
  }

  pair_checks = ["uddi.name", "uddi.method", "uddi.comment", "uddi.disabled", "uddi.pool_availability", "uddi.ttl"]

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      tags   = { Site = "{{random2}}" }
    }
  }

}
