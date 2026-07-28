# Auto-generated datasource acceptance-test cases for RecordA (nios).
case "filters" {
  # filters — generated from terraform-provider-nios
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
      ipv4addr = "10.0.0.20"
      view     = "default"
    }
  }

}

case "ext_attr_filters" {
  # ext_attr_filters — generated from terraform-provider-nios
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
      ipv4addr  = "10.0.0.20"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
