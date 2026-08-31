# Auto-generated datasource acceptance-test cases for ServicerestartGroup.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.mode", "nios.name", "nios.service"]

  step {
    nios {
      name    = "{{random}}"
      service = "DNS"
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

  pair_checks = ["nios.comment", "nios.mode", "nios.name", "nios.service"]

  step {
    nios {
      name      = "{{random}}"
      service   = "DNS"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
