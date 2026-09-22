# Auto-generated datasource acceptance-test cases for DtcMonitorHttp.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.ciphers", "nios.client_cert", "nios.comment", "nios.content_check", "nios.content_check_input", "nios.content_check_op", "nios.content_check_regex", "nios.content_extract_group", "nios.content_extract_type", "nios.content_extract_value", "nios.enable_sni", "nios.interval", "nios.name", "nios.port", "nios.request", "nios.result", "nios.result_code", "nios.retry_down", "nios.retry_up", "nios.secure", "nios.timeout", "nios.validate_cert"]

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

  pair_checks = ["nios.ciphers", "nios.client_cert", "nios.comment", "nios.content_check", "nios.content_check_input", "nios.content_check_op", "nios.content_check_regex", "nios.content_extract_group", "nios.content_extract_type", "nios.content_extract_value", "nios.enable_sni", "nios.interval", "nios.name", "nios.port", "nios.request", "nios.result", "nios.result_code", "nios.retry_down", "nios.retry_up", "nios.secure", "nios.timeout", "nios.validate_cert"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
