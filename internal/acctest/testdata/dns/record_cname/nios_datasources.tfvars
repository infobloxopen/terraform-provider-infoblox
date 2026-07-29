# Auto-generated datasource acceptance-test cases for RecordCname.
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
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
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
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
