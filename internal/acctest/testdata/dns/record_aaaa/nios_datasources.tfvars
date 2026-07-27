# Auto-generated datasource acceptance-test cases for RecordAaaa.
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
      name     = "{{random}}.example.com"
      ipv6addr = "2002:1111::1401"
      view     = "default"
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
      name      = "{{random}}.example.com"
      ipv6addr  = "2002:1111::1401"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
