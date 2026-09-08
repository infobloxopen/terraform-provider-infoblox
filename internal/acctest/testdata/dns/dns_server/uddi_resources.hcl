# TOCO : Objects to be added in the grid for testing
# ACL with IP - 10.0.0.0/24
# TSIG Key with name - tsig-key.

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name"                        = "{{random}}"
      "uddi.ecs_enabled"                 = "false"
      "uddi.filter_aaaa_on_v4"           = "no"
      "uddi.gss_tsig_enabled"            = "false"
      "uddi.notify"                      = "false"
      "uddi.use_forwarders_for_subzones" = "true"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "add_edns_option_in_outgoing_query" {
  backend  = "uddi"
  parallel = true

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

case "auto_sort_views" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name            = "{{random}}"
      auto_sort_views = true
    }
    check = {
      "uddi.auto_sort_views" = "true"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      auto_sort_views = false
    }
    check = {
      "uddi.auto_sort_views" = "false"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

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
      comment = "test updated commentE"
    }
    check = {
      "uddi.comment" = "test updated commentE"
    }
  }

}

case "custom_root_ns" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name           = "{{random}}"
      custom_root_ns = [{ address = "192.168.10.10", fqdn = "tf-example.com." }]
    }
    check = {
      "uddi.custom_root_ns.0.address" = "192.168.10.10"
      "uddi.custom_root_ns.0.fqdn"    = "tf-example.com."
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      custom_root_ns = [{ address = "192.168.11.11", fqdn = "tf-infoblox.com." }, { address = "192.168.11.12", fqdn = "tf-infoblox-acc.com." }]
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
  backend  = "uddi"
  parallel = true

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
      custom_root_ns = [
        {
            address = "192.168.10.10"
            fqdn = "tf-infoblox.com."
        }
      ]
    }
    check = {
      "uddi.custom_root_ns_enabled" = "true"
    }
  }

}

case "dnssec_enable_validation" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                 = "{{random}}"
      dnssec_trust_anchors = [
        {
            algorithm = 8,
            zone = "tf-infoblox.com.",
            sep = false,
            public_key = "AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU="
        }
    ]
    }
    check = {
      "uddi.dnssec_trust_anchors.0.algorithm" = "8"
      "uddi.dnssec_trust_anchors.0.zone"      = "tf-infoblox.com."
      "uddi.dnssec_trust_anchors.0.sep"       = "false"
    }
  }

  step {
    uddi {
      name                 = "{{random}}"
      dnssec_trust_anchors = [
          {
              algorithm = 7,
              zone = "tf-infoblox.com.",
              sep = true,
              public_key = "AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU="
          }
      ]
    }
    check = {
      "uddi.dnssec_trust_anchors.0.algorithm" = "7"
      "uddi.dnssec_trust_anchors.0.zone"      = "tf-infoblox.com."
      "uddi.dnssec_trust_anchors.0.sep"       = "true"
    }
  }

}

case "dnssec_validate_expiry" {
  backend  = "uddi"
  parallel = true

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

case "ecs_enabled" {
  backend  = "uddi"
  parallel = true

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
      ecs_zones = [
        {
            access = "allow"
            fqdn = "tf-infoblox.com."
        }
      ]
    }
    check = {
      "uddi.ecs_enabled" = "true"
    }
  }

}

case "ecs_forwarding" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name      = "{{random}}"
      ecs_zones = [{ access = "allow", fqdn = "tf-infoblox.com." }]
    }
    check = {
      "uddi.ecs_zones.0.access" = "allow"
      "uddi.ecs_zones.0.fqdn"   = "tf-infoblox.com."
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      ecs_zones = [{ access = "deny", fqdn = "tf-test-infoblox.com." }]
    }
    check = {
      "uddi.ecs_zones.0.access" = "deny"
      "uddi.ecs_zones.0.fqdn"   = "tf-test-infoblox.com."
    }
  }

}

case "filter_aaaa_acl" {
  backend  = "uddi"
  parallel = true
  skip = true
  skip_reason = "Requires ACL and TSIG Support"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_acl" "test" {
    uddi = {
      name = "\"acl-\"+name"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_keys_tsig" "test" {
    uddi = {
      name = "\"tsig-\"+name+\".\""
    }
  }
  PREREQ

  step {
    uddi {
      name            = "{{random}}"
      filter_aaaa_acl = [{ access = "allow", element = "ip", address = "192.168.10.10" }]
    }
    check = {
      "uddi.filter_aaaa_acl.0.access"  = "allow"
      "uddi.filter_aaaa_acl.0.element" = "ip"
      "uddi.filter_aaaa_acl.0.address" = "192.168.10.10"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      filter_aaaa_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.filter_aaaa_acl.0.access"  = "deny"
      "uddi.filter_aaaa_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      filter_aaaa_acl = [{ element = "acl", acl = infoblox_dns_acl.test.id }]
    }
    check = {
      "uddi.filter_aaaa_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      filter_aaaa_acl = [{ element = "tsig_key", access = "deny" }]
    }
    check = {
      "uddi.filter_aaaa_acl.0.access"  = "deny"
      "uddi.filter_aaaa_acl.0.element" = "tsig_key"
    }
  }

}

case "filter_aaaa_on_v4" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      forwarders = [{ address = "192.168.10.10", fqdn = "tf-example.com." }]
    }
    check = {
      "uddi.forwarders.0.address" = "192.168.10.10"
      "uddi.forwarders.0.fqdn"    = "tf-example.com."
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      forwarders = [{ address = "192.168.11.11", fqdn = "tf-infoblox.com." }]
    }
    check = {
      "uddi.forwarders.0.address" = "192.168.11.11"
      "uddi.forwarders.0.fqdn"    = "tf-infoblox.com."
    }
  }

}

case "forwarders_only" {
  backend  = "uddi"
  parallel = true

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
      forwarders = [
      		{
      			address = "192.168.11.11"
      			fqdn = "tf-infoblox.com."
      		}
      ]
    }
    check = {
      "uddi.forwarders_only" = "true"
    }
  }
}

case "gss_tsig_enabled" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      inheritance_sources = { add_edns_option_in_outgoing_query = { action = "inherit" }, custom_root_ns_block = { action = "inherit" }, dnssec_validation_block = { action = "inherit" }, ecs_block = { action = "inherit" }, filter_aaaa_on_v4 = { action = "inherit" }, forwarders_block = { action = "inherit" }, gss_tsig_enabled = { action = "inherit" }, kerberos_keys = { action = "inherit" }, lame_ttl = { action = "inherit" }, log_query_response = { action = "inherit" }, match_recursive_only = { action = "inherit" }, max_cache_ttl = { action = "inherit" }, max_negative_ttl = { action = "inherit" }, minimal_responses = { action = "inherit" }, notify = { action = "inherit" }, query_port = { action = "inherit" }, recursion_enabled = { action = "inherit" }, recursive_clients = { action = "inherit" }, resolver_query_timeout = { action = "inherit" }, secondary_axfr_query_limit = { action = "inherit" }, secondary_soa_query_limit = { action = "inherit" }, sort_list = { action = "inherit" }, synthesize_address_records_from_https = { action = "inherit" }, transfer_acl = { action = "inherit" }, use_forwarders_for_subzones = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.add_edns_option_in_outgoing_query.action"     = "inherit"
      "uddi.inheritance_sources.custom_root_ns_block.action"                  = "inherit"
      "uddi.inheritance_sources.dnssec_validation_block.action"               = "inherit"
      "uddi.inheritance_sources.ecs_block.action"                             = "inherit"
      "uddi.inheritance_sources.filter_aaaa_on_v4.action"                     = "inherit"
      "uddi.inheritance_sources.forwarders_block.action"                      = "inherit"
      "uddi.inheritance_sources.gss_tsig_enabled.action"                      = "inherit"
      "uddi.inheritance_sources.kerberos_keys.action"                         = "inherit"
      "uddi.inheritance_sources.lame_ttl.action"                              = "inherit"
      "uddi.inheritance_sources.log_query_response.action"                    = "inherit"
      "uddi.inheritance_sources.match_recursive_only.action"                  = "inherit"
      "uddi.inheritance_sources.max_cache_ttl.action"                         = "inherit"
      "uddi.inheritance_sources.max_negative_ttl.action"                      = "inherit"
      "uddi.inheritance_sources.minimal_responses.action"                     = "inherit"
      "uddi.inheritance_sources.notify.action"                                = "inherit"
      "uddi.inheritance_sources.query_port.action"                            = "inherit"
      "uddi.inheritance_sources.recursion_enabled.action"                     = "inherit"
      "uddi.inheritance_sources.recursive_clients.action"                     = "inherit"
      "uddi.inheritance_sources.resolver_query_timeout.action"                = "inherit"
      "uddi.inheritance_sources.secondary_axfr_query_limit.action"            = "inherit"
      "uddi.inheritance_sources.secondary_soa_query_limit.action"             = "inherit"
      "uddi.inheritance_sources.sort_list.action"                             = "inherit"
      "uddi.inheritance_sources.synthesize_address_records_from_https.action" = "inherit"
      "uddi.inheritance_sources.transfer_acl.action"                          = "inherit"
      "uddi.inheritance_sources.use_forwarders_for_subzones.action"           = "inherit"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      inheritance_sources = { add_edns_option_in_outgoing_query = { action = "override" }, custom_root_ns_block = { action = "override" }, dnssec_validation_block = { action = "override" }, ecs_block = { action = "override" }, filter_aaaa_on_v4 = { action = "override" }, forwarders_block = { action = "override" }, gss_tsig_enabled = { action = "override" }, kerberos_keys = { action = "override" }, lame_ttl = { action = "override" }, log_query_response = { action = "override" }, match_recursive_only = { action = "override" }, max_cache_ttl = { action = "override" }, max_negative_ttl = { action = "override" }, minimal_responses = { action = "override" }, notify = { action = "override" }, query_port = { action = "override" }, recursion_enabled = { action = "override" }, recursive_clients = { action = "override" }, resolver_query_timeout = { action = "override" }, secondary_axfr_query_limit = { action = "override" }, secondary_soa_query_limit = { action = "override" }, sort_list = { action = "override" }, synthesize_address_records_from_https = { action = "override" }, transfer_acl = { action = "override" }, use_forwarders_for_subzones = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.add_edns_option_in_outgoing_query.action"     = "override"
      "uddi.inheritance_sources.custom_root_ns_block.action"                  = "override"
      "uddi.inheritance_sources.dnssec_validation_block.action"               = "override"
      "uddi.inheritance_sources.ecs_block.action"                             = "override"
      "uddi.inheritance_sources.filter_aaaa_on_v4.action"                     = "override"
      "uddi.inheritance_sources.forwarders_block.action"                      = "override"
      "uddi.inheritance_sources.gss_tsig_enabled.action"                      = "override"
      "uddi.inheritance_sources.kerberos_keys.action"                         = "override"
      "uddi.inheritance_sources.lame_ttl.action"                              = "override"
      "uddi.inheritance_sources.log_query_response.action"                    = "override"
      "uddi.inheritance_sources.match_recursive_only.action"                  = "override"
      "uddi.inheritance_sources.max_cache_ttl.action"                         = "override"
      "uddi.inheritance_sources.max_negative_ttl.action"                      = "override"
      "uddi.inheritance_sources.minimal_responses.action"                     = "override"
      "uddi.inheritance_sources.notify.action"                                = "override"
      "uddi.inheritance_sources.query_port.action"                            = "override"
      "uddi.inheritance_sources.recursion_enabled.action"                     = "override"
      "uddi.inheritance_sources.recursive_clients.action"                     = "override"
      "uddi.inheritance_sources.resolver_query_timeout.action"                = "override"
      "uddi.inheritance_sources.secondary_axfr_query_limit.action"            = "override"
      "uddi.inheritance_sources.secondary_soa_query_limit.action"             = "override"
      "uddi.inheritance_sources.sort_list.action"                             = "override"
      "uddi.inheritance_sources.synthesize_address_records_from_https.action" = "override"
      "uddi.inheritance_sources.transfer_acl.action"                          = "override"
      "uddi.inheritance_sources.use_forwarders_for_subzones.action"           = "override"
    }
  }

}

case "lame_ttl" {
  backend  = "uddi"
  parallel = true

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

case "log_query_response" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      log_query_response = true
    }
    check = {
      "uddi.log_query_response" = "true"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      log_query_response = false
    }
    check = {
      "uddi.log_query_response" = "false"
    }
  }

}

case "match_recursive_only" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

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

case "minimal_responses" {
  backend  = "uddi"
  parallel = true

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

case "name" {
  backend  = "uddi"
  parallel = true

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

case "notify" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true
  skip = true
  skip_reason = "Requires ACL and TSIG Support"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_acl" "test" {
    uddi = {
      name = "\"acl-\"+name"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_keys_tsig" "test" {
    uddi = {
      name = "\"tsig-\"+name+\".\""
    }
  }
  PREREQ

  step {
    uddi {
      name      = "{{random}}"
      query_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.query_acl.0.access"  = "allow"
      "uddi.query_acl.0.element" = "ip"
      "uddi.query_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      query_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.query_acl.0.access"  = "deny"
      "uddi.query_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      query_acl = [{ element = "acl", acl = infoblox_dns_acl.test.id }]
    }
    check = {
      "uddi.query_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      query_acl = [{ element = "tsig_key", access = "deny" }]
    }
    check = {
      "uddi.query_acl.0.access"  = "deny"
      "uddi.query_acl.0.element" = "tsig_key"
    }
  }

}

case "query_port" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      query_port = 2
    }
    check = {
      "uddi.query_port" = "2"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      query_port = 10
    }
    check = {
      "uddi.query_port" = "10"
    }
  }

}

case "recursion_acl" {
  backend  = "uddi"
  parallel = true
  skip = true
  skip_reason = "Requires ACL and TSIG Support"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_acl" "test" {
    uddi = {
      name = "\"acl-\"+name"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_keys_tsig" "test" {
    uddi = {
      name = "\"tsig-\"+name+\".\""
    }
  }
  PREREQ

  step {
    uddi {
      name          = "{{random}}"
      recursion_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.recursion_acl.0.access"  = "allow"
      "uddi.recursion_acl.0.element" = "ip"
      "uddi.recursion_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      recursion_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.recursion_acl.0.access"  = "deny"
      "uddi.recursion_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      recursion_acl = [{ element = "acl", acl = infoblox_dns_acl.test.id }]
    }
    check = {
      "uddi.recursion_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      recursion_acl = [{ element = "tsig_key", access = "deny" }]
    }
    check = {
      "uddi.recursion_acl.0.access"  = "deny"
      "uddi.recursion_acl.0.element" = "tsig_key"
    }
  }

}

case "recursion_enabled" {
  backend  = "uddi"
  parallel = true

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

case "recursive_clients" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name              = "{{random}}"
      recursive_clients = 100
    }
    check = {
      "uddi.recursive_clients" = "100"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      recursive_clients = 200
    }
    check = {
      "uddi.recursive_clients" = "200"
    }
  }

}

case "resolver_query_timeout" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                   = "{{random}}"
      resolver_query_timeout = 15
    }
    check = {
      "uddi.resolver_query_timeout" = "15"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      resolver_query_timeout = 20
    }
    check = {
      "uddi.resolver_query_timeout" = "20"
    }
  }

}

case "secondary_axfr_query_limit" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                       = "{{random}}"
      secondary_axfr_query_limit = 2
    }
    check = {
      "uddi.secondary_axfr_query_limit" = "2"
    }
  }

  step {
    uddi {
      name                       = "{{random}}"
      secondary_axfr_query_limit = 3
    }
    check = {
      "uddi.secondary_axfr_query_limit" = "3"
    }
  }

}

case "secondary_soa_query_limit" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                      = "{{random}}"
      secondary_soa_query_limit = 2
    }
    check = {
      "uddi.secondary_soa_query_limit" = "2"
    }
  }

  step {
    uddi {
      name                      = "{{random}}"
      secondary_soa_query_limit = 3
    }
    check = {
      "uddi.secondary_soa_query_limit" = "3"
    }
  }

}

case "sort_list" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name      = "{{random}}"
      sort_list = [
          {
              element = "ip"
              source = "192.168.11.11"
              prioritized_networks = [ "192.168.12.12" ]
           }
       ]
    }
    check = {
      "uddi.sort_list.0.element"                = "ip"
      "uddi.sort_list.0.source"                 = "192.168.11.11"
      "uddi.sort_list.0.prioritized_networks.0" = "192.168.12.12"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      sort_list = [
        {
            element = "any"
            prioritized_networks = [ "192.168.13.13" ]
         }
     ]
    }
    check = {
      "uddi.sort_list.0.element"                = "any"
      "uddi.sort_list.0.prioritized_networks.0" = "192.168.13.13"
    }
  }

}

case "synthesize_address_records_from_https" {
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true

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
  backend  = "uddi"
  parallel = true
  skip = true
  skip_reason = "Requires ACL and TSIG Support"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_acl" "test" {
    uddi = {
      name = "\"acl-\"+name"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_keys_tsig" "test" {
    uddi = {
      name = "\"tsig-\"+name+\".\""
    }
  }
  PREREQ

  step {
    uddi {
      name         = "{{random}}"
      transfer_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.transfer_acl.0.access"  = "allow"
      "uddi.transfer_acl.0.element" = "ip"
      "uddi.transfer_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      transfer_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.transfer_acl.0.access"  = "deny"
      "uddi.transfer_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      transfer_acl = [{ element = "acl", acl = infoblox_dns_acl.test.id }]
    }
    check = {
      "uddi.transfer_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      transfer_acl = [{ element = "tsig_key", access = "deny" }]
    }
    check = {
      "uddi.transfer_acl.0.access"  = "deny"
      "uddi.transfer_acl.0.element" = "tsig_key"
    }
  }

}

case "update_acl" {
  backend  = "uddi"
  parallel = true
  skip = true
  skip_reason = "Requires ACL and TSIG Support"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_acl" "test" {
    uddi = {
      name = "\"acl-\"+name"
      list = [{ access = "allow", element = "ip", address = "10.0.0.0/24" }]
    }
  }
  resource "infoblox_keys_tsig" "test" {
    uddi = {
      name = "\"tsig-\"+name+\".\""
    }
  }
  PREREQ

  step {
    uddi {
      name       = "{{random}}"
      update_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.update_acl.0.access"  = "allow"
      "uddi.update_acl.0.element" = "ip"
      "uddi.update_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      update_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.update_acl.0.access"  = "deny"
      "uddi.update_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      update_acl = [{ element = "acl", acl = infoblox_dns_acl.test.id }]
    }
    check = {
      "uddi.update_acl.0.element" = "acl"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      update_acl = [{ element = "tsig_key", access = "deny" }]
    }
    check = {
      "uddi.update_acl.0.access"  = "deny"
      "uddi.update_acl.0.element" = "tsig_key"
    }
  }

}

case "use_forwarders_for_subzones" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                        = "{{random}}"
      use_forwarders_for_subzones = true
    }
    check = {
      "uddi.use_forwarders_for_subzones" = "true"
    }
  }

  step {
    uddi {
      name                        = "{{random}}"
      use_forwarders_for_subzones = false
    }
    check = {
      "uddi.use_forwarders_for_subzones" = "false"
    }
  }

}

case "use_root_forwarders_for_local_resolution_with_b1td" {
  backend  = "uddi"
  parallel = true

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
      forwarders = [
      		{
      			address = "192.168.11.11"
      			fqdn = "tf-infoblox.com."
      		}
      ]
    }
    check = {
      "uddi.use_root_forwarders_for_local_resolution_with_b1td" = "true"
    }
  }

}
