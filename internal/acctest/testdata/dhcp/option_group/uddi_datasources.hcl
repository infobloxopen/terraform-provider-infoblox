# Auto-generated datasource acceptance-test cases for OptionGroup.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.name", "uddi.protocol"]

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip4"
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

  pair_checks = ["uddi.comment", "uddi.name", "uddi.protocol"]

  step {
    uddi {
      name     = "{{random}}"
      protocol = "ip6"
      tags     = { tag1 = "{{random}}" }
    }
  }

}
