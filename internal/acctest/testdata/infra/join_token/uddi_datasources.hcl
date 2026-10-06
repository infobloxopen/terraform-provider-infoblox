# Auto-generated datasource acceptance-test cases for JoinToken.
#
# join_token is excluded from the pair checks: the API returns it only on the
# create response, so the data source never carries it.

case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.name", "uddi.description"]

  step {
    uddi {
      name        = "{{random}}"
      description = "Join token for filter case"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      Site = "uddi.tags.Site"
    }
  }

  pair_checks = ["uddi.name", "uddi.description"]

  step {
    uddi {
      name        = "{{random}}"
      description = "Join token for tag filter case"
      tags        = { Site = "{{random2}}" }
    }
  }

}
