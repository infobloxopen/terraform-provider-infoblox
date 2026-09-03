# Auto-generated datasource acceptance-test cases for Filteroption.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.apply_as_class", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.expression", "nios.lease_time", "nios.name", "nios.next_server", "nios.option_space", "nios.pxe_lease_time"]

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

  pair_checks = ["nios.apply_as_class", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.expression", "nios.lease_time", "nios.name", "nios.next_server", "nios.option_space", "nios.pxe_lease_time"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
