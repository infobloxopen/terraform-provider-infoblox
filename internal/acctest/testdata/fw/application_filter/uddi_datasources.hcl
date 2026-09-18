case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.name", "uddi.description", "uddi.readonly"]

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type = "tag_filters"
    values = {
      Site = "uddi.tags.Site"
    }
  }

  pair_checks = ["uddi.name", "uddi.description"]

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
      tags     = { Site = "{{random2}}" }
    }
  }

}
