# Auto-generated datasource acceptance-test cases for NsgroupForwardingmember.
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
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
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
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
