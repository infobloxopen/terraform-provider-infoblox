# Auto-generated datasource acceptance-test cases for DtcLbdn.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.auto_consolidated_monitors", "nios.comment", "nios.disable", "nios.lb_method", "nios.name", "nios.persistence", "nios.priority", "nios.topology", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
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

  pair_checks = ["nios.auto_consolidated_monitors", "nios.comment", "nios.disable", "nios.lb_method", "nios.name", "nios.persistence", "nios.priority", "nios.topology", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
