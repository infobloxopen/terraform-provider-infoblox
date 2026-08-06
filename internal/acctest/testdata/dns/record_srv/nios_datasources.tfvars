# Auto-generated datasource acceptance-test cases for RecordSrv.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.port", "nios.priority", "nios.target", "nios.ttl", "nios.use_ttl", "nios.view", "nios.weight"]

  step {
    nios {
      name     = "{{random}}.example.com"
      target   = "{{random2}}.target.com"
      port     = 80
      priority = 10
      weight   = 360
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

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.name", "nios.port", "nios.priority", "nios.target", "nios.ttl", "nios.use_ttl", "nios.view", "nios.weight"]

  step {
    nios {
      name      = "{{random}}.example.com"
      target    = "{{random2}}.target.com"
      port      = 80
      priority  = 10
      weight    = 360
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
