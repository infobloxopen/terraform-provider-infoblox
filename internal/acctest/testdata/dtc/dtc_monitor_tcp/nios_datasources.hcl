# Auto-generated datasource acceptance-test cases for DtcMonitorTcp.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.interval", "nios.name", "nios.port", "nios.retry_down", "nios.retry_up", "nios.timeout"]

  step {
    nios {
      name = "{{random}}"
      port = 49152
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.interval", "nios.name", "nios.port", "nios.retry_down", "nios.retry_up", "nios.timeout"]

  step {
    nios {
      name      = "{{random}}"
      port      = 49152
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
