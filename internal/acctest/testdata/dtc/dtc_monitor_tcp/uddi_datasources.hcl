# Auto-generated datasource acceptance-test cases for DtcMonitorTcp (uddi).
case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.name", "uddi.port", "uddi.comment", "uddi.disabled", "uddi.interval", "uddi.retry_down", "uddi.retry_up", "uddi.timeout"]

  step {
    uddi {
      name = "{{random}}"
      port = 49152
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

  pair_checks = ["uddi.name", "uddi.port", "uddi.comment", "uddi.disabled", "uddi.interval", "uddi.retry_down", "uddi.retry_up", "uddi.timeout"]

  step {
    uddi {
      name = "{{random}}"
      port = 49152
      tags = { Site = "{{random2}}" }
    }
  }

}
