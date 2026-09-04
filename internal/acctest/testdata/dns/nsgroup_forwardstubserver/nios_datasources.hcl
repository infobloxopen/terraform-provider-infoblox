# Auto-generated datasource acceptance-test cases for NsgroupForwardstubserver.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.name"]

  step {
    nios {
      name = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
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

  pair_checks = ["nios.comment", "nios.name"]

  step {
    nios {
      name      = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
