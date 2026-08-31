# Auto-generated datasource acceptance-test cases for DnsServer.
case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.add_edns_option_in_outgoing_query", "uddi.auto_sort_views", "uddi.comment", "uddi.custom_root_ns_enabled", "uddi.dnssec_enable_validation", "uddi.dnssec_enabled", "uddi.dnssec_validate_expiry", "uddi.ecs_enabled", "uddi.ecs_forwarding", "uddi.ecs_prefix_v4", "uddi.ecs_prefix_v6", "uddi.filter_aaaa_on_v4", "uddi.forwarders_only", "uddi.gss_tsig_enabled", "uddi.lame_ttl", "uddi.log_query_response", "uddi.match_recursive_only", "uddi.max_cache_ttl", "uddi.max_negative_ttl", "uddi.minimal_responses", "uddi.name", "uddi.notify", "uddi.query_port", "uddi.recursion_enabled", "uddi.recursive_clients", "uddi.resolver_query_timeout", "uddi.secondary_axfr_query_limit", "uddi.secondary_soa_query_limit", "uddi.synthesize_address_records_from_https", "uddi.use_forwarders_for_subzones", "uddi.use_root_forwarders_for_local_resolution_with_b1td"]

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

  pair_checks = ["uddi.add_edns_option_in_outgoing_query", "uddi.auto_sort_views", "uddi.comment", "uddi.custom_root_ns_enabled", "uddi.dnssec_enable_validation", "uddi.dnssec_enabled", "uddi.dnssec_validate_expiry", "uddi.ecs_enabled", "uddi.ecs_forwarding", "uddi.ecs_prefix_v4", "uddi.ecs_prefix_v6", "uddi.filter_aaaa_on_v4", "uddi.forwarders_only", "uddi.gss_tsig_enabled", "uddi.lame_ttl", "uddi.log_query_response", "uddi.match_recursive_only", "uddi.max_cache_ttl", "uddi.max_negative_ttl", "uddi.minimal_responses", "uddi.name", "uddi.notify", "uddi.query_port", "uddi.recursion_enabled", "uddi.recursive_clients", "uddi.resolver_query_timeout", "uddi.secondary_axfr_query_limit", "uddi.secondary_soa_query_limit", "uddi.synthesize_address_records_from_https", "uddi.use_forwarders_for_subzones", "uddi.use_root_forwarders_for_local_resolution_with_b1td"]

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random}}" }
    }
  }

}
