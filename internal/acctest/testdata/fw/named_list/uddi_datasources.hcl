# Auto-generated datasource acceptance-test cases for NamedList.
case "filters" {
  backend = "uddi"

  # `_filter` on /named_lists only supports the "type" field (confirmed against
  # the live API and its swagger spec). "type" isn't unique to this test's
  # resource, so results.0 may be any pre-existing named list of the same
  # type -- only pair_check what's guaranteed true for every match (the
  # filtered-on field itself), not the resource's other attributes.
  filter {
    type   = "filters"
    values = {
      type = "uddi.type"
    }
  }

  pair_checks = ["uddi.type"]

  step {
    uddi {
      name        = "{{random}}"
      description = "Example Domain"
      type        = "custom_list"
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
