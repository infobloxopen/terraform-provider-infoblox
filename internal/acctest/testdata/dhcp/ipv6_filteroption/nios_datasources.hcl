# Auto-generated datasource acceptance-test cases for Ipv6filteroption.
case "filters" {
  backend = "nios"
  parallel = true
  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.apply_as_class", "nios.comment", "nios.expression", "nios.lease_time", "nios.name", "nios.option_list", "nios.option_space"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  parallel = true
  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.apply_as_class", "nios.comment", "nios.expression", "nios.lease_time", "nios.name", "nios.option_list", "nios.option_space"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
