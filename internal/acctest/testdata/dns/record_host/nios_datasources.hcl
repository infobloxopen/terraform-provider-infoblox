# Auto-generated datasource acceptance-test cases for RecordHost.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.allow_telnet", "nios.comment", "nios.configure_for_dns", "nios.ddns_protected", "nios.device_description", "nios.device_location", "nios.device_type", "nios.device_vendor", "nios.disable", "nios.disable_discovery", "nios.name", "nios.network_view", "nios.rrset_order", "nios.ttl", "nios.use_dns_ea_inheritance", "nios.view"]

  step {
    nios {
      name = "{{random}}.example.com"
      view = "default"
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

  pair_checks = ["nios.allow_telnet", "nios.comment", "nios.configure_for_dns", "nios.ddns_protected", "nios.device_description", "nios.device_location", "nios.device_type", "nios.device_vendor", "nios.disable", "nios.disable_discovery", "nios.name", "nios.network_view", "nios.rrset_order", "nios.ttl", "nios.use_dns_ea_inheritance", "nios.view"]

  step {
    nios {
      name      = "{{random}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
  }

}
