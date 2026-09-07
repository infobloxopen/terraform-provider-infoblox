# Auto-generated datasource acceptance-test cases for NamedList.
case "filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      display_name = "uddi.tags.display_name"
    }
  }

  pair_checks = ["uddi.confidence_level", "uddi.description", "uddi.name", "uddi.threat_level", "uddi.type"]

  step {
    uddi {
      name        = "{{random}}"
      description = "Exaample Domain"
      type        = "custom_list"
      tags        = { display_name = "Terraform Example Named List" }
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      display_name = "uddi.tags.display_name"
    }
  }

  pair_checks = ["uddi.confidence_level", "uddi.description", "uddi.name", "uddi.threat_level", "uddi.type"]

  step {
    uddi {
      name        = "{{random}}"
      description = "Exaample Domain"
      type        = "custom_list"
      tags        = { display_name = "Terraform Example Named List" }
    }
  }

}
