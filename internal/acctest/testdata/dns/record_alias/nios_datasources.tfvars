# Auto-generated datasource acceptance-test cases for RecordAlias.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.disable", "nios.name", "nios.target_name", "nios.target_type", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
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

  pair_checks = ["nios.comment", "nios.creator", "nios.disable", "nios.name", "nios.target_name", "nios.target_type", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      ext_attrs   = { Site = "{{random2}}" }
    }
  }

}
