# Auto-generated datasource acceptance-test cases for RecordPtr.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      ipv4addr = "nios.ipv4addr"
      ptrdname = "nios.ptrdname"
      view     = "nios.view"
    }
  }

  step {
    nios {
      ipv4addr = "192.168.104.122"
      ptrdname = "{{random}}.example.com"
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
      ipv4addr  = "192.168.104.123"
      ptrdname  = "{{random}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
