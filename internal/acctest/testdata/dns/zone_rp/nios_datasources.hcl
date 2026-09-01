# Auto-generated datasource acceptance-test cases for ZoneRp.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      fqdn = "nios.fqdn"
      view = "nios.view"
    }
  }

  pair_checks = ["nios.comment", "nios.disable", "nios.fqdn", "nios.locked", "nios.log_rpz", "nios.ns_group", "nios.prefix", "nios.record_name_policy", "nios.rpz_drop_ip_rule_enabled", "nios.rpz_drop_ip_rule_min_prefix_length_ipv4", "nios.rpz_drop_ip_rule_min_prefix_length_ipv6", "nios.rpz_policy", "nios.rpz_severity", "nios.rpz_type", "nios.soa_default_ttl", "nios.soa_email", "nios.soa_expire", "nios.soa_negative_ttl", "nios.soa_refresh", "nios.soa_retry", "nios.soa_serial_number", "nios.substitute_name", "nios.use_external_primary", "nios.use_grid_zone_timer", "nios.use_log_rpz", "nios.use_record_name_policy", "nios.use_rpz_drop_ip_rule", "nios.use_soa_email", "nios.view"]

  step {
    nios {
      fqdn = "{{random}}.com"
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

  pair_checks = ["nios.comment", "nios.disable", "nios.fqdn", "nios.locked", "nios.log_rpz", "nios.ns_group", "nios.prefix", "nios.record_name_policy", "nios.rpz_drop_ip_rule_enabled", "nios.rpz_drop_ip_rule_min_prefix_length_ipv4", "nios.rpz_drop_ip_rule_min_prefix_length_ipv6", "nios.rpz_policy", "nios.rpz_severity", "nios.rpz_type", "nios.soa_default_ttl", "nios.soa_email", "nios.soa_expire", "nios.soa_negative_ttl", "nios.soa_refresh", "nios.soa_retry", "nios.soa_serial_number", "nios.substitute_name", "nios.use_external_primary", "nios.use_grid_zone_timer", "nios.use_log_rpz", "nios.use_record_name_policy", "nios.use_rpz_drop_ip_rule", "nios.use_soa_email", "nios.view"]

  step {
    nios {
      fqdn      = "{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
