# Auto-generated datasource acceptance-test cases for DhcpOptionspace.
# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: data source config helper 'testAccOptionSpaceDataSourceConfigFilters' could not be parsed (no resource/data block found)
case "filters" {
  backend     = "uddi"
  skip        = true
  skip_reason = "data source config helper 'testAccOptionSpaceDataSourceConfigFilters' could not be parsed (no resource/data block found)"
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
