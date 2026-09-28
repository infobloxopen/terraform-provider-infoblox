# Auto-generated datasource acceptance-test cases for InternalDomainList.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.description", "uddi.name"]

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
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

  pair_checks = ["uddi.description", "uddi.name"]

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
      tags             = { tag1 = "example.somedomain.com" }
    }
  }

}
