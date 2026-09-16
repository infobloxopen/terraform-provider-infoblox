# Auto-generated datasource acceptance-test cases for ZoneAuth.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      fqdn = "uddi.fqdn"
    }
  }

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
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
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
      tags         = { tag1 = "{{random}}" }
    }
  }

}
