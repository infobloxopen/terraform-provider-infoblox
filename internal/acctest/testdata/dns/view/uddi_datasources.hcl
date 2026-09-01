# Auto-generated datasource acceptance-test cases for View.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

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

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random}}" }
    }
  }

}
