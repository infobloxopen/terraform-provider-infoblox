# DtcMonitorIcmp — uddi datasource cases

case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.disabled", "uddi.interval", "uddi.name", "uddi.retry_down", "uddi.retry_up", "uddi.timeout"]

  step {
    uddi {
      name = "dtc-monitor-icmp-{{random}}"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      Site = "uddi.tags_all.Site"
    }
  }

  pair_checks = ["uddi.comment", "uddi.disabled", "uddi.interval", "uddi.name", "uddi.retry_down", "uddi.retry_up", "uddi.timeout"]

  step {
    uddi {
      name = "dtc-monitor-icmp-{{random}}"
      tags = { Site = "{{random2}}" }
    }
  }

}
