# TsigKey — uddi datasource test cases
# Mirrors the legacy BloxOne provider's TSIG data-source coverage
# (internal/service/keys/api_tsig_data_source_test.go): filters and tag_filters.
case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  step {
    uddi {
      name   = "{{random}}."
      secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
      tags   = { tag1 = "{{random}}" }
    }
  }

}
