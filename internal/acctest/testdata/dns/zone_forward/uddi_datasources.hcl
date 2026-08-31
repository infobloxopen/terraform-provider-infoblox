# Auto-generated datasource acceptance-test cases for ZoneForward.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      fqdn = "uddi.fqdn"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.disabled", "uddi.forward_only", "uddi.fqdn", "uddi.parent", "uddi.view"]

  step {
    uddi {
      fqdn = "{{random}}.com."
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.disabled", "uddi.forward_only", "uddi.fqdn", "uddi.parent", "uddi.view"]

  step {
    uddi {
      fqdn = "{{random}}.com."
      tags = { tag1 = "{{random}}" }
    }
  }

}
