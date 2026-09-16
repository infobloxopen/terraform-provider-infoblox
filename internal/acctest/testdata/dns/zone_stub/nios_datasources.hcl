# Auto-generated datasource acceptance-test cases for ZoneStub.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      fqdn = "nios.fqdn"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.disable_forwarding", "nios.external_ns_group", "nios.fqdn", "nios.locked", "nios.ms_ad_integrated", "nios.ms_ddns_mode", "nios.ns_group", "nios.prefix", "nios.view", "nios.zone_format"]

  step {
    nios {
      fqdn      = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
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

  pair_checks = ["nios.comment", "nios.disable", "nios.disable_forwarding", "nios.external_ns_group", "nios.fqdn", "nios.locked", "nios.ms_ad_integrated", "nios.ms_ddns_mode", "nios.ns_group", "nios.prefix", "nios.view", "nios.zone_format"]

  step {
    nios {
      fqdn = "{{random}}.com"
      stub_from = [{ address = "1.1.1.1", name = "{{random2}}" }]
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
