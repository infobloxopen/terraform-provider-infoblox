# Hand-authored datasource acceptance-test cases for RecordRpzNaptr.

case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.flags", "nios.name", "nios.order", "nios.preference", "nios.regexp", "nios.replacement", "nios.rp_zone", "nios.services", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.tf-acc-rpz-naptr.example.com"
      rp_zone     = "tf-acc-rpz-naptr.example.com"
      order       = 10
      preference  = 10
      replacement = "."
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

  pair_checks = ["nios.comment", "nios.disable", "nios.flags", "nios.name", "nios.order", "nios.preference", "nios.regexp", "nios.replacement", "nios.rp_zone", "nios.services", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name        = "{{random2}}.tf-acc-rpz-naptr.example.com"
      rp_zone     = "tf-acc-rpz-naptr.example.com"
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random}}" }
    }
  }

}
