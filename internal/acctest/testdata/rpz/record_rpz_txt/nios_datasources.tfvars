# Auto-generated datasource acceptance-test cases for RecordRpzTxt.
# TODO: The following prerequisites MUST exist on the grid before running these tests:
#   - RPZ zone : test-rpz.com  (view: default)

case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.text", "nios.ttl", "nios.view"]

  step {
    nios {
      name    = "{{random2}}.test-rpz.com"
      text    = "Record Text"
      rp_zone = "test-rpz.com"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.text", "nios.ttl", "nios.view"]

  step {
    nios {
      name      = "{{random2}}.test-rpz.com"
      text      = "Record Text"
      rp_zone   = "test-rpz.com"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
