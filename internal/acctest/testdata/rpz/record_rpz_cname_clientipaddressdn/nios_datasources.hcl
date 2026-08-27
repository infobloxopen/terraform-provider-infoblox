case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.view"]

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
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

  pair_checks = ["nios.canonical", "nios.comment", "nios.disable", "nios.name", "nios.rp_zone", "nios.view"]

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
