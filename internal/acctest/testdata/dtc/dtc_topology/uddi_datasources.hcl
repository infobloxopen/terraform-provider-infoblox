# Auto-generated datasource acceptance-test cases for DtcTopology (UDDI backend).
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = { name = "uddi.name" }
  }

  pair_checks = [
    "uddi.name",
    "uddi.comment",
    "uddi.disabled",
    "uddi.sources.#",
  ]

  step {
    uddi {
      name    = "topology-{{random}}"
      comment = "datasource-filter-test"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = { Site = "uddi.tags.Site" }
  }

  pair_checks = [
    "uddi.name",
    "uddi.sources.#",
  ]

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
      tags    = { Site = "{{random2}}" }
    }
  }

}
