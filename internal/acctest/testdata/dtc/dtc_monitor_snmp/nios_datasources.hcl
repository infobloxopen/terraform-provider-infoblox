# Auto-generated datasource acceptance-test cases for DtcMonitorSnmp.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.community", "nios.context", "nios.engine_id", "nios.interval", "nios.name", "nios.port", "nios.retry_down", "nios.retry_up", "nios.timeout", "nios.user", "nios.version"]

  step {
    nios {
      name = "{{random}}"
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

  pair_checks = ["nios.comment", "nios.community", "nios.context", "nios.engine_id", "nios.interval", "nios.name", "nios.port", "nios.retry_down", "nios.retry_up", "nios.timeout", "nios.user", "nios.version"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
