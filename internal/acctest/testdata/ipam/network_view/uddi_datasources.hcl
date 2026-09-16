# Auto-generated datasource acceptance-test cases for Networkview.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.ddns_client_update", "uddi.ddns_conflict_resolution_mode", "uddi.ddns_domain", "uddi.ddns_generate_name", "uddi.ddns_generated_prefix", "uddi.ddns_send_updates", "uddi.ddns_ttl_percent", "uddi.ddns_update_on_renew", "uddi.ddns_use_conflict_resolution", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname_rewrite_char", "uddi.hostname_rewrite_enabled", "uddi.hostname_rewrite_regex", "uddi.name", "uddi.vendor_specific_option_option_space"]

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.ddns_client_update", "uddi.ddns_conflict_resolution_mode", "uddi.ddns_domain", "uddi.ddns_generate_name", "uddi.ddns_generated_prefix", "uddi.ddns_send_updates", "uddi.ddns_ttl_percent", "uddi.ddns_update_on_renew", "uddi.ddns_use_conflict_resolution", "uddi.header_option_filename", "uddi.header_option_server_address", "uddi.header_option_server_name", "uddi.hostname_rewrite_char", "uddi.hostname_rewrite_enabled", "uddi.hostname_rewrite_regex", "uddi.name", "uddi.vendor_specific_option_option_space"]

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random}}" }
    }
  }

}
