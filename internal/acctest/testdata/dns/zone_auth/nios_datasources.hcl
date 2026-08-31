# Auto-generated datasource acceptance-test cases for ZoneAuth.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      fqdn = "nios.fqdn"
      view = "nios.view"
    }
  }

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
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
      fqdn      = "{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
