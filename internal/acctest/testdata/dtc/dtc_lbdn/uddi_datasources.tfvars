# Auto-generated datasource acceptance-test cases for DtcLbdn (UDDI backend).
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = { name = "uddi.name" }
  }

  pair_checks = ["uddi.name", "uddi.view", "uddi.comment", "uddi.disabled", "uddi.ttl"]

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = { Site = "uddi.tags.Site" }
  }

  pair_checks = ["uddi.name", "uddi.view", "uddi.comment", "uddi.disabled", "uddi.ttl"]

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      tags = { Site = "{{random2}}" }
    }
  }

}
