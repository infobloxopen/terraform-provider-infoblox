# Auto-generated datasource acceptance-test cases for Namedacl.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.name"]

  step {
    uddi {
      name = "{{random}}"
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

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.name"]

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random}}" }
    }
  }

}
