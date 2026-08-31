# Auto-generated datasource acceptance-test cases for DtcServer.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.auto_create_host_record", "nios.comment", "nios.disable", "nios.host", "nios.name", "nios.sni_hostname", "nios.use_sni_hostname"]

  step {
    nios {
      name = "{{random}}"
      host = "{{random_ip}}"
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

  pair_checks = ["nios.auto_create_host_record", "nios.comment", "nios.disable", "nios.host", "nios.name", "nios.sni_hostname", "nios.use_sni_hostname"]

  step {
    nios {
      name      = "{{random}}"
      host      = "{{random_ip}}"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
