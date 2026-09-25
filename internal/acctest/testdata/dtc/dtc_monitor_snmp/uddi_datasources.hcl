case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = { name = "uddi.name" }
  }

  pair_checks = ["uddi.name", "uddi.version", "uddi.comment", "uddi.community", "uddi.disabled", "uddi.interval", "uddi.port", "uddi.retry_down", "uddi.retry_up", "uddi.timeout"]

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
    }
  }
}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = { Site = "uddi.tags.Site" }
  }

  pair_checks = ["uddi.name", "uddi.version", "uddi.comment", "uddi.community", "uddi.disabled", "uddi.interval", "uddi.port", "uddi.retry_down", "uddi.retry_up", "uddi.timeout"]

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
      tags    = { Site = "{{random2}}" }
    }
  }
}
