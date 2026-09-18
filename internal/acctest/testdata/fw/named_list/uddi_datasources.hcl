# Auto-generated datasource acceptance-test cases for NamedList.
case "filters" {
  backend = "uddi"

  # `_filter` only supports "type", which isn't unique to this test's resource,
  # so only pair_check the filtered-on field itself, not the other attributes.
  filter {
    type   = "filters"
    values = {
      type = "uddi.type"
    }
  }

  pair_checks = ["uddi.type"]

  step {
    uddi {
      # trimspace(...) is a no-op; it just skips the auto pair-check on "name".
      name = trimspace("{{random}}")
      type = "custom_list"
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
