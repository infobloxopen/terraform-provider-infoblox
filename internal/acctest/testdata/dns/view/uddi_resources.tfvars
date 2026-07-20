# Auto-generated resource acceptance-test cases for View (uddi).
case "basic" {
  # basic — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name"                        = "{{random}}"
      "uddi.disabled"                    = "false"
      "uddi.ecs_enabled"                 = "false"
      "uddi.filter_aaaa_on_v4"           = "no"
      "uddi.gss_tsig_enabled"            = "false"
      "uddi.notify"                      = "false"
      "uddi.use_forwarders_for_subzones" = "true"
    }
  }

}

case "disappears" {
  # disappears — generated from terraform-provider-uddi
  backend = "uddi"
  disappears = true
  expect_non_empty_plan = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "name" {
  # name — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "add_edns_option_in_outgoing_query" {
  # add_edns_option_in_outgoing_query — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                              = "{{random}}"
      add_edns_option_in_outgoing_query = false
    }
    check = {
      "uddi.add_edns_option_in_outgoing_query" = "false"
    }
  }

  step {
    uddi {
      name                              = "{{random}}"
      add_edns_option_in_outgoing_query = true
    }
    check = {
      "uddi.add_edns_option_in_outgoing_query" = "true"
    }
  }

}

case "comment" {
  # comment — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "another test comment"
    }
    check = {
      "uddi.comment" = "another test comment"
    }
  }

}

case "custom_root_ns" {
  # custom_root_ns — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.custom_root_ns.0.address" = "192.168.10.10"
      "uddi.custom_root_ns.0.fqdn"    = "tf-example.com."
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.custom_root_ns.#"         = "2"
      "uddi.custom_root_ns.0.address" = "192.168.11.11"
      "uddi.custom_root_ns.0.fqdn"    = "tf-infoblox.com."
      "uddi.custom_root_ns.1.address" = "192.168.11.12"
      "uddi.custom_root_ns.1.fqdn"    = "tf-infoblox-acc.com."
    }
  }

}

case "custom_root_ns_enabled" {
  # custom_root_ns_enabled — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                   = "{{random}}"
      custom_root_ns_enabled = false
    }
    check = {
      "uddi.custom_root_ns_enabled" = "false"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      custom_root_ns_enabled = true
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      custom_root_ns_enabled = true
    }
    check = {
      "uddi.custom_root_ns_enabled" = "true"
    }
  }

}

case "disabled" {
  # disabled — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name     = "{{random}}"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "dnssec_enable_validation" {
  # dnssec_enable_validation — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                     = "{{random}}"
      dnssec_enable_validation = true
    }
    check = {
      "uddi.dnssec_enable_validation" = "true"
    }
  }

  step {
    uddi {
      name                     = "{{random}}"
      dnssec_enable_validation = false
    }
    check = {
      "uddi.dnssec_enable_validation" = "false"
    }
  }

}

case "dnssec_enabled" {
  # dnssec_enabled — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name           = "{{random}}"
      dnssec_enabled = true
    }
    check = {
      "uddi.dnssec_enabled" = "true"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      dnssec_enabled = false
    }
    check = {
      "uddi.dnssec_enabled" = "false"
    }
  }

}

case "dnssec_trust_anchors" {
  # dnssec_trust_anchors — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.dnssec_trust_anchors.0.algorithm" = "8"
      "uddi.dnssec_trust_anchors.0.zone"      = "tf-infoblox.com."
      "uddi.dnssec_trust_anchors.0.sep"       = "false"
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.dnssec_trust_anchors.0.algorithm" = "7"
      "uddi.dnssec_trust_anchors.0.zone"      = "tf-infoblox.com."
      "uddi.dnssec_trust_anchors.0.sep"       = "true"
    }
  }

}

case "dnssec_validate_expiry" {
  # dnssec_validate_expiry — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                   = "{{random}}"
      dnssec_validate_expiry = true
    }
    check = {
      "uddi.dnssec_validate_expiry" = "true"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      dnssec_validate_expiry = false
    }
    check = {
      "uddi.dnssec_validate_expiry" = "false"
    }
  }

}

case "dtc_config" {
  # dtc_config — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name       = "{{random}}"
      dtc_config = { default_ttl = 700 }
    }
    check = {
      "uddi.dtc_config.default_ttl" = "700"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      dtc_config = { default_ttl = 500 }
    }
    check = {
      "uddi.dtc_config.default_ttl" = "500"
    }
  }

}

case "ecs_enabled" {
  # ecs_enabled — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name        = "{{random}}"
      ecs_enabled = false
    }
    check = {
      "uddi.ecs_enabled" = "false"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      ecs_enabled = true
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      ecs_enabled = true
    }
    check = {
      "uddi.ecs_enabled" = "true"
    }
  }

}

case "ecs_forwarding" {
  # ecs_forwarding — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name           = "{{random}}"
      ecs_forwarding = false
    }
    check = {
      "uddi.ecs_forwarding" = "false"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      ecs_forwarding = true
    }
    check = {
      "uddi.ecs_forwarding" = "true"
    }
  }

}

case "ecs_prefix_v4" {
  # ecs_prefix_v4 — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name          = "{{random}}"
      ecs_prefix_v4 = 20
    }
    check = {
      "uddi.ecs_prefix_v4" = "20"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      ecs_prefix_v4 = 1
    }
    check = {
      "uddi.ecs_prefix_v4" = "1"
    }
  }

}

case "ecs_prefix_v6" {
  # ecs_prefix_v6 — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name          = "{{random}}"
      ecs_prefix_v6 = 50
    }
    check = {
      "uddi.ecs_prefix_v6" = "50"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      ecs_prefix_v6 = 1
    }
    check = {
      "uddi.ecs_prefix_v6" = "1"
    }
  }

}

case "ecs_zones" {
  # ecs_zones — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.ecs_zones.0.access" = "allow"
      "uddi.ecs_zones.0.fqdn"   = "tf-infoblox.com."
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.ecs_zones.0.access" = "deny"
      "uddi.ecs_zones.0.fqdn"   = "tf-test-infoblox.com."
    }
  }

}

case "edns_udp_size" {
  # edns_udp_size — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name          = "{{random}}"
      edns_udp_size = 1200
    }
    check = {
      "uddi.edns_udp_size" = "1200"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      edns_udp_size = 1000
    }
    check = {
      "uddi.edns_udp_size" = "1000"
    }
  }

}

case "filter_aaaa_acl" {
  # filter_aaaa_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "filter_aaaa_on_v4" {
  # filter_aaaa_on_v4 — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name              = "{{random}}"
      filter_aaaa_on_v4 = "no"
    }
    check = {
      "uddi.filter_aaaa_on_v4" = "no"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      filter_aaaa_on_v4 = "break_dnssec"
    }
    check = {
      "uddi.filter_aaaa_on_v4" = "break_dnssec"
    }
  }

}

case "forwarders" {
  # forwarders — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.forwarders.0.address" = "192.168.10.10"
      "uddi.forwarders.0.fqdn"    = "tf-example.com."
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.forwarders.0.address" = "192.168.11.11"
      "uddi.forwarders.0.fqdn"    = "tf-infoblox.com."
    }
  }

}

case "forwarders_only" {
  # forwarders_only — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      forwarders_only = false
    }
    check = {
      "uddi.forwarders_only" = "false"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      forwarders_only = true
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      forwarders_only = true
    }
    check = {
      "uddi.forwarders_only" = "true"
    }
  }

}

case "gss_tsig_enabled" {
  # gss_tsig_enabled — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name             = "{{random}}"
      gss_tsig_enabled = false
    }
    check = {
      "uddi.gss_tsig_enabled" = "false"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      gss_tsig_enabled = true
    }
    check = {
      "uddi.gss_tsig_enabled" = "true"
    }
  }

}

case "inheritance_sources" {
  # inheritance_sources — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                = "{{random}}"
      inheritance_sources = { add_edns_option_in_outgoing_query = { action = "inherit" }, custom_root_ns_block = { action = "inherit" }, dnssec_validation_block = { action = "inherit" }, ecs_block = { action = "inherit" }, edns_udp_size = { action = "inherit" }, filter_aaaa_on_v4 = { action = "inherit" }, forwarders_block = { action = "inherit" }, gss_tsig_enabled = { action = "inherit" }, kerberos_keys = { action = "inherit" }, lame_ttl = { action = "inherit" }, match_recursive_only = { action = "inherit" }, max_cache_ttl = { action = "inherit" }, max_negative_ttl = { action = "inherit" }, minimal_responses = { action = "inherit" }, notify = { action = "inherit" }, recursion_enabled = { action = "inherit" }, sort_list = { action = "inherit" }, synthesize_address_records_from_https = { action = "inherit" }, transfer_acl = { action = "inherit" }, use_forwarders_for_subzones = { action = "inherit" }, zone_authority = { default_ttl = { action = "inherit" }, expire = { action = "inherit" }, mname_block = { action = "inherit" }, negative_ttl = { action = "inherit" }, refresh = { action = "inherit" }, retry = { action = "inherit" }, rname = { action = "inherit" } } }
    }
    check = {
      "uddi.inheritance_sources.add_edns_option_in_outgoing_query.action"     = "inherit"
      "uddi.inheritance_sources.custom_root_ns_block.action"                  = "inherit"
      "uddi.inheritance_sources.dnssec_validation_block.action"               = "inherit"
      "uddi.inheritance_sources.ecs_block.action"                             = "inherit"
      "uddi.inheritance_sources.filter_aaaa_on_v4.action"                     = "inherit"
      "uddi.inheritance_sources.forwarders_block.action"                      = "inherit"
      "uddi.inheritance_sources.gss_tsig_enabled.action"                      = "inherit"
      "uddi.inheritance_sources.lame_ttl.action"                              = "inherit"
      "uddi.inheritance_sources.edns_udp_size.action"                         = "inherit"
      "uddi.inheritance_sources.match_recursive_only.action"                  = "inherit"
      "uddi.inheritance_sources.max_cache_ttl.action"                         = "inherit"
      "uddi.inheritance_sources.max_negative_ttl.action"                      = "inherit"
      "uddi.inheritance_sources.minimal_responses.action"                     = "inherit"
      "uddi.inheritance_sources.notify.action"                                = "inherit"
      "uddi.inheritance_sources.recursion_enabled.action"                     = "inherit"
      "uddi.inheritance_sources.sort_list.action"                             = "inherit"
      "uddi.inheritance_sources.synthesize_address_records_from_https.action" = "inherit"
      "uddi.inheritance_sources.transfer_acl.action"                          = "inherit"
      "uddi.inheritance_sources.use_forwarders_for_subzones.action"           = "inherit"
      "uddi.inheritance_sources.zone_authority.default_ttl.action"            = "inherit"
      "uddi.inheritance_sources.zone_authority.expire.action"                 = "inherit"
      "uddi.inheritance_sources.zone_authority.mname_block.action"            = "inherit"
      "uddi.inheritance_sources.zone_authority.negative_ttl.action"           = "inherit"
      "uddi.inheritance_sources.zone_authority.refresh.action"                = "inherit"
      "uddi.inheritance_sources.zone_authority.retry.action"                  = "inherit"
      "uddi.inheritance_sources.zone_authority.rname.action"                  = "inherit"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      inheritance_sources = { add_edns_option_in_outgoing_query = { action = "override" }, custom_root_ns_block = { action = "override" }, dnssec_validation_block = { action = "override" }, ecs_block = { action = "override" }, edns_udp_size = { action = "override" }, filter_aaaa_on_v4 = { action = "override" }, forwarders_block = { action = "override" }, gss_tsig_enabled = { action = "override" }, kerberos_keys = { action = "override" }, lame_ttl = { action = "override" }, match_recursive_only = { action = "override" }, max_cache_ttl = { action = "override" }, max_negative_ttl = { action = "override" }, minimal_responses = { action = "override" }, notify = { action = "override" }, recursion_enabled = { action = "override" }, sort_list = { action = "override" }, synthesize_address_records_from_https = { action = "override" }, transfer_acl = { action = "override" }, use_forwarders_for_subzones = { action = "override" }, zone_authority = { default_ttl = { action = "override" }, expire = { action = "override" }, mname_block = { action = "override" }, negative_ttl = { action = "override" }, refresh = { action = "override" }, retry = { action = "override" }, rname = { action = "override" } } }
    }
    check = {
      "uddi.inheritance_sources.add_edns_option_in_outgoing_query.action"     = "override"
      "uddi.inheritance_sources.custom_root_ns_block.action"                  = "override"
      "uddi.inheritance_sources.dnssec_validation_block.action"               = "override"
      "uddi.inheritance_sources.ecs_block.action"                             = "override"
      "uddi.inheritance_sources.filter_aaaa_on_v4.action"                     = "override"
      "uddi.inheritance_sources.forwarders_block.action"                      = "override"
      "uddi.inheritance_sources.gss_tsig_enabled.action"                      = "override"
      "uddi.inheritance_sources.lame_ttl.action"                              = "override"
      "uddi.inheritance_sources.edns_udp_size.action"                         = "override"
      "uddi.inheritance_sources.match_recursive_only.action"                  = "override"
      "uddi.inheritance_sources.max_cache_ttl.action"                         = "override"
      "uddi.inheritance_sources.max_negative_ttl.action"                      = "override"
      "uddi.inheritance_sources.minimal_responses.action"                     = "override"
      "uddi.inheritance_sources.notify.action"                                = "override"
      "uddi.inheritance_sources.recursion_enabled.action"                     = "override"
      "uddi.inheritance_sources.sort_list.action"                             = "override"
      "uddi.inheritance_sources.synthesize_address_records_from_https.action" = "override"
      "uddi.inheritance_sources.transfer_acl.action"                          = "override"
      "uddi.inheritance_sources.use_forwarders_for_subzones.action"           = "override"
      "uddi.inheritance_sources.zone_authority.default_ttl.action"            = "override"
      "uddi.inheritance_sources.zone_authority.expire.action"                 = "override"
      "uddi.inheritance_sources.zone_authority.mname_block.action"            = "override"
      "uddi.inheritance_sources.zone_authority.negative_ttl.action"           = "override"
      "uddi.inheritance_sources.zone_authority.refresh.action"                = "override"
      "uddi.inheritance_sources.zone_authority.retry.action"                  = "override"
      "uddi.inheritance_sources.zone_authority.rname.action"                  = "override"
    }
  }

}

case "ip_spaces" {
  # ip_spaces — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "helper declares prerequisite resource 'bloxone_ipam_ip_space' which has no buildable infoblox equivalent (not in prereq_type_map.json)"
}

case "lame_ttl" {
  # lame_ttl — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name     = "{{random}}"
      lame_ttl = 3000
    }
    check = {
      "uddi.lame_ttl" = "3000"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      lame_ttl = 3600
    }
    check = {
      "uddi.lame_ttl" = "3600"
    }
  }

}

case "match_clients_acl" {
  # match_clients_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "match_destinations_acl" {
  # match_destinations_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "match_recursive_only" {
  # match_recursive_only — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                 = "{{random}}"
      match_recursive_only = false
    }
    check = {
      "uddi.match_recursive_only" = "false"
    }
  }

  step {
    uddi {
      name                 = "{{random}}"
      match_recursive_only = true
    }
    check = {
      "uddi.match_recursive_only" = "true"
    }
  }

}

case "max_cache_ttl" {
  # max_cache_ttl — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name          = "{{random}}"
      max_cache_ttl = 600000
    }
    check = {
      "uddi.max_cache_ttl" = "600000"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      max_cache_ttl = 1
    }
    check = {
      "uddi.max_cache_ttl" = "1"
    }
  }

}

case "max_negative_ttl" {
  # max_negative_ttl — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name             = "{{random}}"
      max_negative_ttl = 10000
    }
    check = {
      "uddi.max_negative_ttl" = "10000"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      max_negative_ttl = 1
    }
    check = {
      "uddi.max_negative_ttl" = "1"
    }
  }

}

case "max_udp_size" {
  # max_udp_size — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name         = "{{random}}"
      max_udp_size = 1232
    }
    check = {
      "uddi.max_udp_size" = "1232"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      max_udp_size = 512
    }
    check = {
      "uddi.max_udp_size" = "512"
    }
  }

}

case "minimal_responses" {
  # minimal_responses — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name              = "{{random}}"
      minimal_responses = false
    }
    check = {
      "uddi.minimal_responses" = "false"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      minimal_responses = true
    }
    check = {
      "uddi.minimal_responses" = "true"
    }
  }

}

case "notify" {
  # notify — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name   = "{{random}}"
      notify = false
    }
    check = {
      "uddi.notify" = "false"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      notify = true
    }
    check = {
      "uddi.notify" = "true"
    }
  }

}

case "query_acl" {
  # query_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "recursion_acl" {
  # recursion_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "recursion_enabled" {
  # recursion_enabled — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name              = "{{random}}"
      recursion_enabled = true
    }
    check = {
      "uddi.recursion_enabled" = "true"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      recursion_enabled = false
    }
    check = {
      "uddi.recursion_enabled" = "false"
    }
  }

}

case "sort_list" {
  # sort_list — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.sort_list.0.element"                = "ip"
      "uddi.sort_list.0.source"                 = "192.168.11.11"
      "uddi.sort_list.0.prioritized_networks.0" = "192.168.12.12"
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.sort_list.0.element"                = "any"
      "uddi.sort_list.0.prioritized_networks.0" = "192.168.13.13"
    }
  }

}

case "synthesize_address_records_from_https" {
  # synthesize_address_records_from_https — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                                  = "{{random}}"
      synthesize_address_records_from_https = false
    }
    check = {
      "uddi.synthesize_address_records_from_https" = "false"
    }
  }

  step {
    uddi {
      name                                  = "{{random}}"
      synthesize_address_records_from_https = true
    }
    check = {
      "uddi.synthesize_address_records_from_https" = "true"
    }
  }

}

case "tags" {
  # tags — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}

case "transfer_acl" {
  # transfer_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "update_acl" {
  # update_acl — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "config helper 'testAccAclIP' could not be parsed (no resource block found)"
}

case "use_root_forwarders_for_local_resolution_with_b1td" {
  # use_root_forwarders_for_local_resolution_with_b1td — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name                                               = "{{random}}"
      use_root_forwarders_for_local_resolution_with_b1td = false
    }
    check = {
      "uddi.use_root_forwarders_for_local_resolution_with_b1td" = "false"
    }
  }

  step {
    uddi {
      name                                               = "{{random}}"
      use_root_forwarders_for_local_resolution_with_b1td = true
    }
  }

  step {
    uddi {
      name                                               = "{{random}}"
      use_root_forwarders_for_local_resolution_with_b1td = true
    }
    check = {
      "uddi.use_root_forwarders_for_local_resolution_with_b1td" = "true"
    }
  }

}

case "zone_authority" {
  # zone_authority — generated from terraform-provider-uddi
  backend = "uddi"

  step {
    uddi {
      name           = "{{random}}"
      zone_authority = { default_ttl = 28600, expire = 2519200, mname = "test.b1ddi", negative_ttl = 700, refresh = 10500, retry = 3500, rname = "host", use_default_mname = false }
    }
    check = {
      "uddi.zone_authority.default_ttl"       = "28600"
      "uddi.zone_authority.expire"            = "2519200"
      "uddi.zone_authority.mname"             = "test.b1ddi"
      "uddi.zone_authority.negative_ttl"      = "700"
      "uddi.zone_authority.refresh"           = "10500"
      "uddi.zone_authority.retry"             = "3500"
      "uddi.zone_authority.rname"             = "host"
      "uddi.zone_authority.use_default_mname" = "false"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      zone_authority = { default_ttl = 30000, expire = 2519200, mname = "test-infoblox.b1ddi", negative_ttl = 800, refresh = 11800, retry = 3700, rname = "host-test", use_default_mname = false }
    }
    check = {
      "uddi.zone_authority.default_ttl"       = "30000"
      "uddi.zone_authority.expire"            = "2519200"
      "uddi.zone_authority.mname"             = "test-infoblox.b1ddi"
      "uddi.zone_authority.negative_ttl"      = "800"
      "uddi.zone_authority.refresh"           = "11800"
      "uddi.zone_authority.retry"             = "3700"
      "uddi.zone_authority.rname"             = "host-test"
      "uddi.zone_authority.use_default_mname" = "false"
    }
  }

}
