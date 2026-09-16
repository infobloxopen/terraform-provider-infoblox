# Auto-generated datasource acceptance-test cases for View.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

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

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
