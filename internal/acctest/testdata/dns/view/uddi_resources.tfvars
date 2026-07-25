# Auto-generated resource acceptance-test cases for View.
case "basic" {
  backend = "uddi"
  parallel = true

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
  backend = "uddi"
  disappears = true
  expect_non_empty_plan = true
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "name" {
  backend = "uddi"
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

case "add_edns_option_in_outgoing_query" {
  backend = "uddi"
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

case "compartment_id" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name           = "{{random}}"
      compartment_id = "c4695."
    }
    check = {
      "uddi.compartment_id" = "c4695."
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      compartment_id = ""
    }
    check = {
      "uddi.compartment_id" = ""
    }
  }

}

case "comment" {
  backend = "uddi"
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
      comment = "another test comment"
    }
    check = {
      "uddi.comment" = "another test comment"
    }
  }

}

case "custom_root_ns" {
  backend = "uddi"
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
  backend = "uddi"
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
      custom_root_ns = [{ address = "192.168.10.10", fqdn = "tf-example.com." }]
      custom_root_ns_enabled = true
    }
    check = {
      "uddi.custom_root_ns_enabled" = "true"
    }
  }

}

case "disabled" {
  backend = "uddi"
  parallel = true

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
  backend = "uddi"
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
  backend = "uddi"
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
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      dnssec_trust_anchors = [
        {
          algorithm = 8
          public_key = "AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU="
          zone      = "tf-infoblox.com."
          sep       = false
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
      name = "{{random}}"
      dnssec_trust_anchors = [
        {
          algorithm = 7
          public_key = "AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU="
          zone      = "tf-infoblox.com."
          sep       = true
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
  backend = "uddi"
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

case "dtc_config" {
  backend = "uddi"
  parallel = true

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
  backend = "uddi"
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
      ecs_zones = [
        {
          access = "allow"
          fqdn = "tf-infoblox.com."
        }
      ]
      ecs_enabled = true
    }
    check = {
      "uddi.ecs_enabled" = "true"
    }
  }

}

case "ecs_forwarding" {
  backend = "uddi"
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
  backend = "uddi"
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
  backend = "uddi"
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
  backend = "uddi"
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

case "edns_udp_size" {
  backend = "uddi"
  parallel = true

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
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

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
      filter_aaaa_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.filter_aaaa_acl.0.element" = "acl"
      "uddi.filter_aaaa_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      filter_aaaa_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.filter_aaaa_acl.0.access"       = "deny"
      "uddi.filter_aaaa_acl.0.element"      = "tsig_key"
      "uddi.filter_aaaa_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "filter_aaaa_on_v4" {
  backend = "uddi"
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
  backend = "uddi"
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
  backend = "uddi"
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
      forwarders = [
          {
            address = "192.168.11.11"
            fqdn = "tf-infoblox.com."
          }
      ]
      forwarders_only = true
    }
    check = {
      "uddi.forwarders_only" = "true"
    }
  }

}

case "gss_tsig_enabled" {
  backend = "uddi"
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
  backend = "uddi"
  parallel = true

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
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test_space" {
  #   uddi = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      name      = "{{random}}"
      ip_spaces = ["ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"]
    }
    check = {
      "uddi.ip_spaces.#" = "1"
      "uddi.ip_spaces.0" = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      ip_spaces = ["ipam/ip_space/1fcd4065-8847-11f1-b283-5eecb1762ec1"]
    }
    check = {
      "uddi.ip_spaces.#" = "1"
      "uddi.ip_spaces.0" = "ipam/ip_space/1fcd4065-8847-11f1-b283-5eecb1762ec1"
    }
  }

}

case "lame_ttl" {
  backend = "uddi"
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

case "match_clients_acl" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

  step {
    uddi {
      name              = "{{random}}"
      match_clients_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.match_clients_acl.0.access"  = "allow"
      "uddi.match_clients_acl.0.element" = "ip"
      "uddi.match_clients_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      match_clients_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.match_clients_acl.0.access"  = "deny"
      "uddi.match_clients_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      match_clients_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.match_clients_acl.0.element" = "acl"
      "uddi.match_clients_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      match_clients_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.match_clients_acl.0.access"       = "deny"
      "uddi.match_clients_acl.0.element"      = "tsig_key"
      "uddi.match_clients_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "match_destinations_acl" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

  step {
    uddi {
      name                   = "{{random}}"
      match_destinations_acl = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.match_destinations_acl.0.access"  = "allow"
      "uddi.match_destinations_acl.0.element" = "ip"
      "uddi.match_destinations_acl.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      match_destinations_acl = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.match_destinations_acl.0.access"  = "deny"
      "uddi.match_destinations_acl.0.element" = "any"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      match_destinations_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.match_destinations_acl.0.element" = "acl"
      "uddi.match_destinations_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      match_destinations_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.match_destinations_acl.0.access"       = "deny"
      "uddi.match_destinations_acl.0.element"      = "tsig_key"
      "uddi.match_destinations_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "match_recursive_only" {
  backend = "uddi"
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
  backend = "uddi"
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
  backend = "uddi"
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

case "max_udp_size" {
  backend = "uddi"
  parallel = true

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
  backend = "uddi"
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

case "notify" {
  backend = "uddi"
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
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

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
      query_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.query_acl.0.element" = "acl"
      "uddi.query_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      query_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.query_acl.0.access"       = "deny"
      "uddi.query_acl.0.element"      = "tsig_key"
      "uddi.query_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "recursion_acl" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

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
      recursion_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.recursion_acl.0.element" = "acl"
      "uddi.recursion_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name          = "{{random}}"
      recursion_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.recursion_acl.0.access"       = "deny"
      "uddi.recursion_acl.0.element"      = "tsig_key"
      "uddi.recursion_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "recursion_enabled" {
  backend = "uddi"
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

case "sort_list" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      sort_list = [
        {
          element = "ip"
          source  = "192.168.11.11"
          prioritized_networks = ["192.168.12.12"]
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
      name = "{{random}}"
      sort_list = [
        {
          element = "any"
          prioritized_networks = ["192.168.13.13"]
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
  backend = "uddi"
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
  backend = "uddi"
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
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

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
      transfer_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.transfer_acl.0.element" = "acl"
      "uddi.transfer_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
      transfer_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.transfer_acl.0.access"       = "deny"
      "uddi.transfer_acl.0.element"      = "tsig_key"
      "uddi.transfer_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "update_acl" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_acl_unknown" "test" {
  #   uddi = {
  #     name = "\"acl-\"+name"
  #   }
  # }
  # resource "infoblox_tsig_key_unknown" "test" {
  #   uddi = {
  #   }
  # }
  # PREREQ

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
      update_acl = [{ element = "acl", acl = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4" }]
    }
    check = {
      "uddi.update_acl.0.element" = "acl"
      "uddi.update_acl.0.acl"     = "dns/acl/86db2788-6e9d-40ad-ab18-79f8ada358b4"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      update_acl = [{ access = "deny", element = "tsig_key", tsig_key = { key = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240" } }]
    }
    check = {
      "uddi.update_acl.0.access"       = "deny"
      "uddi.update_acl.0.element"      = "tsig_key"
      "uddi.update_acl.0.tsig_key.key" = "keys/tsig/4832d039-dad9-4e82-813c-ecc56385b240"
    }
  }

}

case "use_root_forwarders_for_local_resolution_with_b1td" {
  backend = "uddi"
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
      forwarders = [
          {
            address = "192.168.11.11"
            fqdn = "tf-infoblox.com."
          }
      ]
      use_root_forwarders_for_local_resolution_with_b1td = true
    }
    check = {
      "uddi.use_root_forwarders_for_local_resolution_with_b1td" = "true"
    }
  }

}

case "zone_authority" {
  backend = "uddi"
  parallel = true

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
