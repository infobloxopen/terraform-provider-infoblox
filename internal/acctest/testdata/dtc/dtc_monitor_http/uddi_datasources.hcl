case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.name", "uddi.port", "uddi.comment", "uddi.disabled", "uddi.https", "uddi.interval", "uddi.timeout", "uddi.retry_down", "uddi.retry_up", "uddi.request"]

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      Site = "uddi.tags.Site"
    }
  }

  pair_checks = ["uddi.name", "uddi.port", "uddi.comment", "uddi.disabled", "uddi.https", "uddi.interval", "uddi.timeout", "uddi.retry_down", "uddi.retry_up", "uddi.request"]

  step {
    uddi {
      name    = "dtc-monitor-http-{{random}}"
      port    = 80
      request = "GET / HTTP/1.0\r\n\r\n"
      tags    = { Site = "{{random2}}" }
    }
  }

}
