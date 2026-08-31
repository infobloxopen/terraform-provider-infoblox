# Auto-generated datasource acceptance-test cases for DtcPool.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.auto_consolidated_monitors", "nios.availability", "nios.comment", "nios.disable", "nios.lb_alternate_method", "nios.lb_alternate_topology", "nios.lb_preferred_method", "nios.lb_preferred_topology", "nios.name", "nios.quorum", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
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

  pair_checks = ["nios.auto_consolidated_monitors", "nios.availability", "nios.comment", "nios.disable", "nios.lb_alternate_method", "nios.lb_alternate_topology", "nios.lb_preferred_method", "nios.lb_preferred_topology", "nios.name", "nios.quorum", "nios.ttl", "nios.use_ttl"]

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      ext_attrs           = { Site = "{{random2}}" }
    }
  }

}
