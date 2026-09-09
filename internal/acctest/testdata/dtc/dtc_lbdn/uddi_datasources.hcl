# Auto-generated datasource acceptance-test cases for DtcLbdn (UDDI backend).
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = { name = "uddi.name" }
  }

  pair_checks = ["uddi.name", "uddi.view", "uddi.comment", "uddi.disabled", "uddi.ttl", "uddi.precedence"]

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
    }
  }

}

case "tag_filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  filter {
    type   = "tag_filters"
    values = { Site = "uddi.tags.Site" }
  }

  pair_checks = ["uddi.name", "uddi.view", "uddi.comment", "uddi.disabled", "uddi.ttl", "uddi.precedence"]

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
      tags = { Site = "{{random2}}" }
    }
  }

}
