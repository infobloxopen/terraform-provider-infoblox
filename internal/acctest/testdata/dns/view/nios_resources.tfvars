# Auto-generated resource acceptance-test cases for View.
//
// TODO : Objects to be present in the grid before running the test cases
// Blacklist Rulesets (type BLACKLIST) - blacklist_ruleset_1, blacklist_ruleset_2
// NXDOMAIN Rulesets (type NXDOMAIN)   - nxdomain_ruleset_1, nxdomain_ruleset_2
// DNS64 Group                         - dns64_group
// DDNS Principal Cluster Group        - dynamic_update_grp_1, dynamic_update_grp_2
//
case "basic" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                                            = "{{random}}"
      "nios.blacklist_action"                                = "REDIRECT"
      "nios.blacklist_log_query"                             = "false"
      "nios.blacklist_redirect_ttl"                          = "60"
      "nios.ddns_force_creation_timestamp_update"            = "false"
      "nios.ddns_principal_tracking"                         = "false"
      "nios.ddns_restrict_patterns"                          = "false"
      "nios.ddns_restrict_protected"                         = "false"
      "nios.ddns_restrict_secure"                            = "false"
      "nios.ddns_restrict_static"                            = "false"
      "nios.disable"                                         = "false"
      "nios.dns64_enabled"                                   = "false"
      "nios.dnssec_enabled"                                  = "false"
      "nios.dnssec_expired_signatures_enabled"               = "false"
      "nios.dnssec_validation_enabled"                       = "true"
      "nios.edns_udp_size"                                   = "1220"
      "nios.enable_blacklist"                                = "false"
      "nios.enable_fixed_rrset_order_fqdns"                  = "false"
      "nios.enable_match_recursive_only"                     = "false"
      "nios.filter_aaaa"                                     = "NO"
      "nios.forward_only"                                    = "false"
      "nios.max_cache_ttl"                                   = "604800"
      "nios.max_ncache_ttl"                                  = "10800"
      "nios.max_udp_size"                                    = "1220"
      "nios.network_view"                                    = "default"
      "nios.notify_delay"                                    = "5"
      "nios.nxdomain_log_query"                              = "false"
      "nios.nxdomain_redirect"                               = "false"
      "nios.nxdomain_redirect_ttl"                           = "60"
      "nios.recursion"                                       = "false"
      "nios.response_rate_limiting.enable_rrl"               = "false"
      "nios.response_rate_limiting.log_only"                 = "false"
      "nios.response_rate_limiting.responses_per_second"     = "100"
      "nios.response_rate_limiting.window"                   = "15"
      "nios.response_rate_limiting.slip"                     = "2"
      "nios.root_name_server_type"                           = "INTERNET"
      "nios.rpz_drop_ip_rule_enabled"                        = "false"
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv4"         = "29"
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv6"         = "112"
      "nios.rpz_qname_wait_recurse"                          = "false"
      "nios.scavenging_settings.enable_auto_reclamation"     = "false"
      "nios.scavenging_settings.enable_recurrent_scavenging" = "false"
      "nios.scavenging_settings.enable_rr_last_queried"      = "false"
      "nios.scavenging_settings.enable_scavenging"           = "false"
      "nios.scavenging_settings.enable_zone_last_queried"    = "false"
      "nios.scavenging_settings.reclaim_associated_records"  = "false"
    }
  }

}

case "disappears" {
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "blacklist_action" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      blacklist_action = "REFUSE"
    }
    check = {
      "nios.blacklist_action" = "REFUSE"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      blacklist_action = "REDIRECT"
    }
    check = {
      "nios.blacklist_action" = "REDIRECT"
    }
  }

}

case "blacklist_log_query" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      blacklist_log_query = true
    }
    check = {
      "nios.blacklist_log_query" = "true"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      blacklist_log_query = false
    }
    check = {
      "nios.blacklist_log_query" = "false"
    }
  }

}

case "blacklist_redirect_addresses" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                         = "{{random}}"
      blacklist_redirect_addresses = ["10.0.0.1", "10.0.0.29"]
    }
    check = {
      "nios.blacklist_redirect_addresses.0" = "10.0.0.1"
      "nios.blacklist_redirect_addresses.1" = "10.0.0.29"
    }
  }

  step {
    nios {
      name                         = "{{random}}"
      blacklist_redirect_addresses = ["10.0.0.23", "10.0.0.54"]
    }
    check = {
      "nios.blacklist_redirect_addresses.0" = "10.0.0.23"
      "nios.blacklist_redirect_addresses.1" = "10.0.0.54"
    }
  }

}

case "blacklist_redirect_ttl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                   = "{{random}}"
      blacklist_redirect_ttl = 75
    }
    check = {
      "nios.blacklist_redirect_ttl" = "75"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      blacklist_redirect_ttl = 90
    }
    check = {
      "nios.blacklist_redirect_ttl" = "90"
    }
  }

}

case "blacklist_rulesets" {
  backend = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ruleset_unknown" "test_ruleset1" {
  #   nios = {
  #     name = "{{random2}}"
  #     type = "BLACKLIST"
  #   }
  # }
  # resource "infoblox_ruleset_unknown" "test_ruleset3" {
  #   nios = {
  #     name = "{{random3}}"
  #     type = "BLACKLIST"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name               = "{{random}}"
      blacklist_rulesets = ["blacklist_ruleset_1"]
    }
    # depends_on = [infoblox_ruleset_unknown.test_ruleset1, infoblox_ruleset_unknown.test_ruleset3]
    check = {
      "nios.blacklist_rulesets.0" = "blacklist_ruleset_1"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      blacklist_rulesets = ["blacklist_ruleset_2"]
    }
    # depends_on = [infoblox_ruleset_unknown.test_ruleset1, infoblox_ruleset_unknown.test_ruleset3]
    check = {
      "nios.blacklist_rulesets.0" = "blacklist_ruleset_2"
    }
  }

}

case "comment" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "dns view comment"
    }
    check = {
      "nios.comment" = "dns view comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "updated dns view comment"
    }
    check = {
      "nios.comment" = "updated dns view comment"
    }
  }

}

case "custom_root_name_servers" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                     = "{{random}}"
      custom_root_name_servers = [{ address = "10.0.0.2", name = "external-server-1" }]
    }
    check = {
      "nios.custom_root_name_servers.0.address" = "10.0.0.2"
      "nios.custom_root_name_servers.0.name"    = "external-server-1"
    }
  }

  step {
    nios {
      name                     = "{{random}}"
      custom_root_name_servers = [{ address = "10.0.0.23", name = "external-server-2" }]
    }
    check = {
      "nios.custom_root_name_servers.0.address" = "10.0.0.23"
      "nios.custom_root_name_servers.0.name"    = "external-server-2"
    }
  }

}

case "ddns_force_creation_timestamp_update" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                                 = "{{random}}"
      ddns_force_creation_timestamp_update = true
    }
    check = {
      "nios.ddns_force_creation_timestamp_update" = "true"
    }
  }

  step {
    nios {
      name                                 = "{{random}}"
      ddns_force_creation_timestamp_update = false
    }
    check = {
      "nios.ddns_force_creation_timestamp_update" = "false"
    }
  }

}

case "ddns_principal_group" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      ddns_principal_group = "dynamic_update_grp_1"
    }
    check = {
      "nios.ddns_principal_group" = "dynamic_update_grp_1"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      ddns_principal_group = "dynamic_update_grp_2"
    }
    check = {
      "nios.ddns_principal_group" = "dynamic_update_grp_2"
    }
  }

}

case "ddns_principal_tracking" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      ddns_principal_tracking = true
    }
    check = {
      "nios.ddns_principal_tracking" = "true"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      ddns_principal_tracking = false
    }
    check = {
      "nios.ddns_principal_tracking" = "false"
    }
  }

}

case "ddns_restrict_patterns" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                   = "{{random}}"
      ddns_restrict_patterns = true
    }
    check = {
      "nios.ddns_restrict_patterns" = "true"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      ddns_restrict_patterns = false
    }
    check = {
      "nios.ddns_restrict_patterns" = "false"
    }
  }

}

case "ddns_restrict_patterns_list" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                        = "{{random}}"
      ddns_restrict_patterns_list = ["pattern1.example.com"]
    }
    check = {
      "nios.ddns_restrict_patterns_list.0" = "pattern1.example.com"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      ddns_restrict_patterns_list = ["pattern2.example.com", "pattern3.example.com"]
    }
    check = {
      "nios.ddns_restrict_patterns_list.0" = "pattern2.example.com"
      "nios.ddns_restrict_patterns_list.1" = "pattern3.example.com"
    }
  }

}

case "ddns_restrict_protected" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      ddns_restrict_protected = true
    }
    check = {
      "nios.ddns_restrict_protected" = "true"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      ddns_restrict_protected = false
    }
    check = {
      "nios.ddns_restrict_protected" = "false"
    }
  }

}

case "ddns_restrict_secure" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      ddns_restrict_secure = true
    }
    check = {
      "nios.ddns_restrict_secure" = "true"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      ddns_restrict_secure = false
    }
    check = {
      "nios.ddns_restrict_secure" = "false"
    }
  }

}

case "ddns_restrict_static" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      ddns_restrict_static = true
    }
    check = {
      "nios.ddns_restrict_static" = "true"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      ddns_restrict_static = false
    }
    check = {
      "nios.ddns_restrict_static" = "false"
    }
  }

}

case "disable" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      disable = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "dns64_enabled" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      dns64_enabled = true
      dns64_groups  = ["default"]
    }
    check = {
      "nios.dns64_enabled" = "true"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      dns64_enabled = false
      dns64_groups  = ["default"]
    }
    check = {
      "nios.dns64_enabled" = "false"
    }
  }

}

case "dns64_groups" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      dns64_groups = ["default"]
    }
    check = {
      "nios.dns64_groups.0" = "default"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      dns64_groups = ["dns64_group"]
    }
    check = {
      "nios.dns64_groups.0" = "dns64_group"
    }
  }

}

case "dnssec_enabled" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      dnssec_enabled = true
    }
    check = {
      "nios.dnssec_enabled" = "true"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      dnssec_enabled = false
    }
    check = {
      "nios.dnssec_enabled" = "false"
    }
  }

}

case "dnssec_expired_signatures_enabled" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                              = "{{random}}"
      dnssec_expired_signatures_enabled = true
    }
    check = {
      "nios.dnssec_expired_signatures_enabled" = "true"
    }
  }

  step {
    nios {
      name                              = "{{random}}"
      dnssec_expired_signatures_enabled = false
    }
    check = {
      "nios.dnssec_expired_signatures_enabled" = "false"
    }
  }

}

case "dnssec_negative_trust_anchors" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                          = "{{random}}"
      dnssec_negative_trust_anchors = ["examplezone.com"]
    }
    check = {
      "nios.dnssec_negative_trust_anchors.0" = "examplezone.com"
    }
  }

  step {
    nios {
      name                          = "{{random}}"
      dnssec_negative_trust_anchors = ["examplezone2.com", "examplezone3.com"]
    }
    check = {
      "nios.dnssec_negative_trust_anchors.0" = "examplezone2.com"
      "nios.dnssec_negative_trust_anchors.1" = "examplezone3.com"
    }
  }

}

case "dnssec_trusted_keys" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      dnssec_trusted_keys = [{ algorithm = "14", dnssec_must_be_secure = false, fqdn = "test.com", key = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAweW4MAnsKGjk4dt6a42CrIA/BV9YEKThzXZVlBSdUfn0D2YDOMkWvlMxUPVd5iEc2DXulrpBSNbxL1y7Ude11fs1+cOgvcgmQX1Yvu9e14OzeMfk3ZJB8Ldnmb5xrNR9y4ASqh771PZA6xK3qVS+k7YLGp3xnRrd1+zMLcUMI5J+8ZBOIn/6K37DkirKhBv5hKfttTNQbPiwDXwS/vduUv0vUN/xLUKg6099abOn05nefWg+BoxuMySVtqhB6pgW+1BrGrSISOTZDTKojguftya3vqFhb5m/G3F39BdIAlNWP/P2lP8ksuER/pczE6muS8CS2ArCbaN+Z7iddg5P6wIDAQAB", secure_entry_point = true }]
    }
    check = {
      "nios.dnssec_trusted_keys.0.algorithm"             = "14"
      "nios.dnssec_trusted_keys.0.dnssec_must_be_secure" = "false"
      "nios.dnssec_trusted_keys.0.fqdn"                  = "test.com"
      "nios.dnssec_trusted_keys.0.key"                   = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAweW4MAnsKGjk4dt6a42CrIA/BV9YEKThzXZVlBSdUfn0D2YDOMkWvlMxUPVd5iEc2DXulrpBSNbxL1y7Ude11fs1+cOgvcgmQX1Yvu9e14OzeMfk3ZJB8Ldnmb5xrNR9y4ASqh771PZA6xK3qVS+k7YLGp3xnRrd1+zMLcUMI5J+8ZBOIn/6K37DkirKhBv5hKfttTNQbPiwDXwS/vduUv0vUN/xLUKg6099abOn05nefWg+BoxuMySVtqhB6pgW+1BrGrSISOTZDTKojguftya3vqFhb5m/G3F39BdIAlNWP/P2lP8ksuER/pczE6muS8CS2ArCbaN+Z7iddg5P6wIDAQAB"
      "nios.dnssec_trusted_keys.0.secure_entry_point"    = "true"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      dnssec_trusted_keys = [{ algorithm = "14", dnssec_must_be_secure = false, fqdn = "test2.com", key = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAweW4MAnsKGjk4dt6a42CrIA/BV9YEKThzXZVlBSdUfn0D2YDOMkWvlMxUPVd5iEc2DXulrpBSNbxL1y7Ude11fs1+cOgvcgmQX1Yvu9e14OzeMfk3ZJB8Ldnmb5xrNR9y4ASqh771PZA6xK3qVS+k7YLGp3xnRrd1+zMLcUMI5J+8ZBOIn/6K37DkirKhBv5hKfttTNQbPiwDXwS/vduUv0vUN/xLUKg6099abOn05nefWg+BoxuMySVtqhB6pgW+1BrGrSISOTZDTKojguftya3vqFhb5m/G3F39BdIAlNWP/P2lP8ksuER/pczE6muS8CS2ArCbaN+Z7iddg5P6wIDAQAB", secure_entry_point = true }]
    }
    check = {
      "nios.dnssec_trusted_keys.0.algorithm"             = "14"
      "nios.dnssec_trusted_keys.0.dnssec_must_be_secure" = "false"
      "nios.dnssec_trusted_keys.0.fqdn"                  = "test2.com"
      "nios.dnssec_trusted_keys.0.key"                   = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAweW4MAnsKGjk4dt6a42CrIA/BV9YEKThzXZVlBSdUfn0D2YDOMkWvlMxUPVd5iEc2DXulrpBSNbxL1y7Ude11fs1+cOgvcgmQX1Yvu9e14OzeMfk3ZJB8Ldnmb5xrNR9y4ASqh771PZA6xK3qVS+k7YLGp3xnRrd1+zMLcUMI5J+8ZBOIn/6K37DkirKhBv5hKfttTNQbPiwDXwS/vduUv0vUN/xLUKg6099abOn05nefWg+BoxuMySVtqhB6pgW+1BrGrSISOTZDTKojguftya3vqFhb5m/G3F39BdIAlNWP/P2lP8ksuER/pczE6muS8CS2ArCbaN+Z7iddg5P6wIDAQAB"
      "nios.dnssec_trusted_keys.0.secure_entry_point"    = "true"
    }
  }

}

case "dnssec_validation_enabled" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                      = "{{random}}"
      dnssec_validation_enabled = false
    }
    check = {
      "nios.dnssec_validation_enabled" = "false"
    }
  }

  step {
    nios {
      name                      = "{{random}}"
      dnssec_validation_enabled = true
    }
    check = {
      "nios.dnssec_validation_enabled" = "true"
    }
  }

}

case "edns_udp_size" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      edns_udp_size = 1232
    }
    check = {
      "nios.edns_udp_size" = "1232"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      edns_udp_size = 4096
    }
    check = {
      "nios.edns_udp_size" = "4096"
    }
  }

}

case "enable_blacklist" {
  backend = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ruleset_unknown" "test_ruleset1" {
  #   nios = {
  #     name = "{{random2}}"
  #     type = "BLACKLIST"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                         = "{{random}}"
      enable_blacklist             = true
      blacklist_redirect_addresses = ["10.0.0.2"]
      blacklist_rulesets           = ["blacklist_ruleset_1"]
    }
    check = {
      "nios.enable_blacklist" = "true"
    }
  }

  step {
    nios {
      name                         = "{{random}}"
      enable_blacklist             = false
      blacklist_redirect_addresses = ["10.0.0.2"]
      blacklist_rulesets           = ["blacklist_ruleset_1"]
    }
    check = {
      "nios.enable_blacklist" = "false"
    }
  }

}

case "enable_fixed_rrset_order_fqdns" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                           = "{{random}}"
      enable_fixed_rrset_order_fqdns = true
    }
    check = {
      "nios.enable_fixed_rrset_order_fqdns" = "true"
    }
  }

  step {
    nios {
      name                           = "{{random}}"
      enable_fixed_rrset_order_fqdns = false
    }
    check = {
      "nios.enable_fixed_rrset_order_fqdns" = "false"
    }
  }

}

case "enable_match_recursive_only" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                        = "{{random}}"
      enable_match_recursive_only = true
    }
    check = {
      "nios.enable_match_recursive_only" = "true"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      enable_match_recursive_only = false
    }
    check = {
      "nios.enable_match_recursive_only" = "false"
    }
  }

}

case "ext_attrs" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "filter_aaaa" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      filter_aaaa = "BREAK_DNSSEC"
    }
    check = {
      "nios.filter_aaaa" = "BREAK_DNSSEC"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      filter_aaaa = "YES"
    }
    check = {
      "nios.filter_aaaa" = "YES"
    }
  }

}

case "filter_aaaa_list" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      filter_aaaa_list = [{ address = "10.0.0.23", permission = "DENY" }]
    }
    check = {
      "nios.filter_aaaa_list.0.address"    = "10.0.0.23"
      "nios.filter_aaaa_list.0.permission" = "DENY"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      filter_aaaa_list = [{ address = "10.0.0.12", permission = "ALLOW" }]
    }
    check = {
      "nios.filter_aaaa_list.0.address"    = "10.0.0.12"
      "nios.filter_aaaa_list.0.permission" = "ALLOW"
    }
  }

}

case "fixed_rrset_order_fqdns" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      fixed_rrset_order_fqdns = [{ fqdn = "example.com", record_type = "AAAA" }]
    }
    check = {
      "nios.fixed_rrset_order_fqdns.0.fqdn"        = "example.com"
      "nios.fixed_rrset_order_fqdns.0.record_type" = "AAAA"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      fixed_rrset_order_fqdns = [{ fqdn = "example.org", record_type = "BOTH" }]
    }
    check = {
      "nios.fixed_rrset_order_fqdns.0.fqdn"        = "example.org"
      "nios.fixed_rrset_order_fqdns.0.record_type" = "BOTH"
    }
  }

}

case "forward_only" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      forward_only = true
      forwarders   = ["10.192.81.23"]
    }
    check = {
      "nios.forward_only" = "true"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      forward_only = false
      forwarders   = ["10.192.81.23"]
    }
    check = {
      "nios.forward_only" = "false"
    }
  }

}

case "forwarders" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      forwarders = ["10.123.86.42"]
    }
    check = {
      "nios.forwarders.0" = "10.123.86.42"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      forwarders = ["10.252.23.44"]
    }
    check = {
      "nios.forwarders.0" = "10.252.23.44"
    }
  }

}

case "last_queried_acl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      last_queried_acl = [{ address = "10.0.0.23", permission = "DENY" }]
    }
    check = {
      "nios.last_queried_acl.0.address"    = "10.0.0.23"
      "nios.last_queried_acl.0.permission" = "DENY"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      last_queried_acl = [{ address = "10.0.0.12", permission = "ALLOW" }]
    }
    check = {
      "nios.last_queried_acl.0.address"    = "10.0.0.12"
      "nios.last_queried_acl.0.permission" = "ALLOW"
    }
  }

}

case "match_clients" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      match_clients = [{ struct = "addressac", address = "10.0.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.match_clients.0.address"    = "10.0.0.0"
      "nios.match_clients.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      match_clients = [{ struct = "addressac", address = "192.168.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.match_clients.0.address"    = "192.168.0.0"
      "nios.match_clients.0.permission" = "ALLOW"
    }
  }

}

case "match_destinations" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      match_destinations = [{ struct = "addressac", address = "10.0.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.match_destinations.0.address"    = "10.0.0.0"
      "nios.match_destinations.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      match_destinations = [{ struct = "addressac", address = "192.168.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.match_destinations.0.address"    = "192.168.0.0"
      "nios.match_destinations.0.permission" = "ALLOW"
    }
  }

}

case "max_cache_ttl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      max_cache_ttl = 3600
    }
    check = {
      "nios.max_cache_ttl" = "3600"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      max_cache_ttl = 7200
    }
    check = {
      "nios.max_cache_ttl" = "7200"
    }
  }

}

case "max_ncache_ttl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      max_ncache_ttl = 300
    }
    check = {
      "nios.max_ncache_ttl" = "300"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      max_ncache_ttl = 600
    }
    check = {
      "nios.max_ncache_ttl" = "600"
    }
  }

}

case "max_udp_size" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      max_udp_size = 512
    }
    check = {
      "nios.max_udp_size" = "512"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      max_udp_size = 1024
    }
    check = {
      "nios.max_udp_size" = "1024"
    }
  }

}

case "name" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "network_view" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      network_view = "default"
    }
    check = {
      "nios.network_view" = "default"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      network_view = "test_network_view"
    }
    check = {
      "nios.network_view" = "test_network_view"
    }
  }

}

case "notify_delay" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      notify_delay = 78
    }
    check = {
      "nios.notify_delay" = "78"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      notify_delay = 10
    }
    check = {
      "nios.notify_delay" = "10"
    }
  }

}

case "nxdomain_log_query" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      nxdomain_log_query = true
    }
    check = {
      "nios.nxdomain_log_query" = "true"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      nxdomain_log_query = false
    }
    check = {
      "nios.nxdomain_log_query" = "false"
    }
  }

}

case "nxdomain_redirect" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                        = "{{random}}"
      nxdomain_redirect           = true
      nxdomain_redirect_addresses = ["10.45.3.2"]
    }
    check = {
      "nios.nxdomain_redirect" = "true"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      nxdomain_redirect           = false
      nxdomain_redirect_addresses = ["10.45.3.2"]
    }
    check = {
      "nios.nxdomain_redirect" = "false"
    }
  }

}

case "nxdomain_redirect_addresses" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                        = "{{random}}"
      nxdomain_redirect_addresses = ["10.87.9.7"]
    }
    check = {
      "nios.nxdomain_redirect_addresses.0" = "10.87.9.7"
    }
  }

  step {
    nios {
      name                        = "{{random}}"
      nxdomain_redirect_addresses = ["10.3.23.56", "5.4.3.5"]
    }
    check = {
      "nios.nxdomain_redirect_addresses.0" = "10.3.23.56"
      "nios.nxdomain_redirect_addresses.1" = "5.4.3.5"
    }
  }

}

case "nxdomain_redirect_addresses_v6" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                           = "{{random}}"
      nxdomain_redirect_addresses_v6 = ["2001:db8::1", "2001:db8::2"]
    }
    check = {
      "nios.nxdomain_redirect_addresses_v6.0" = "2001:db8::1"
      "nios.nxdomain_redirect_addresses_v6.1" = "2001:db8::2"
    }
  }

  step {
    nios {
      name                           = "{{random}}"
      nxdomain_redirect_addresses_v6 = ["2001:db8::3", "2001:db8::4"]
    }
    check = {
      "nios.nxdomain_redirect_addresses_v6.0" = "2001:db8::3"
      "nios.nxdomain_redirect_addresses_v6.1" = "2001:db8::4"
    }
  }

}

case "nxdomain_redirect_ttl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      nxdomain_redirect_ttl = 3600
    }
    check = {
      "nios.nxdomain_redirect_ttl" = "3600"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      nxdomain_redirect_ttl = 7200
    }
    check = {
      "nios.nxdomain_redirect_ttl" = "7200"
    }
  }

}

case "nxdomain_rulesets" {
  backend = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ruleset_unknown" "test_ruleset1" {
  #   nios = {
  #     name = "nxdomain_ruleset"
  #     type = "NXDOMAIN"
  #   }
  # }
  # resource "infoblox_ruleset_unknown" "test_ruleset2" {
  #   nios = {
  #     name = "nxdomain_ruleset2"
  #     type = "NXDOMAIN"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name              = "{{random}}"
      nxdomain_rulesets = ["nxdomain_ruleset_1"]
    }
    # depends_on = [infoblox_ruleset_unknown.test_ruleset1, infoblox_ruleset_unknown.test_ruleset2]
    check = {
      "nios.nxdomain_rulesets.0" = "nxdomain_ruleset_1"
    }
  }

  step {
    nios {
      name              = "{{random}}"
      nxdomain_rulesets = ["nxdomain_ruleset_2"]
    }
    # depends_on = [infoblox_ruleset_unknown.test_ruleset1, infoblox_ruleset_unknown.test_ruleset2]
    check = {
      "nios.nxdomain_rulesets.0" = "nxdomain_ruleset_2"
    }
  }

}

case "recursion" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      recursion = true
    }
    check = {
      "nios.recursion" = "true"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      recursion = false
    }
    check = {
      "nios.recursion" = "false"
    }
  }

}

case "response_rate_limiting" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                   = "{{random}}"
      response_rate_limiting = {
        enable_rrl           = false,
        log_only             = false,
        responses_per_second = 100,
        slip                 = 2,
        window               = 15,
      }
    }
    check = {
      "nios.response_rate_limiting.enable_rrl"           = "false"
      "nios.response_rate_limiting.log_only"             = "false"
      "nios.response_rate_limiting.responses_per_second" = "100"
      "nios.response_rate_limiting.slip"                 = "2"
      "nios.response_rate_limiting.window"               = "15"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      response_rate_limiting = {
        enable_rrl           = true,
        log_only             = true,
        responses_per_second = 200,
        slip                 = 3,
        window               = 30,
      }
    }
    check = {
      "nios.response_rate_limiting.enable_rrl"           = "true"
      "nios.response_rate_limiting.log_only"             = "true"
      "nios.response_rate_limiting.responses_per_second" = "200"
      "nios.response_rate_limiting.slip"                 = "3"
      "nios.response_rate_limiting.window"               = "30"
    }
  }

}

case "root_name_server_type" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                     = "{{random}}"
      root_name_server_type    = "CUSTOM"
      custom_root_name_servers = [{ address = "10.0.0.2", name = "external-server-1" }]
    }
    check = {
      "nios.root_name_server_type" = "CUSTOM"
    }
  }

  step {
    nios {
      name                     = "{{random}}"
      root_name_server_type    = "INTERNET"
      custom_root_name_servers = [{ address = "10.0.0.2", name = "external-server-1" }]
    }
    check = {
      "nios.root_name_server_type" = "INTERNET"
    }
  }

}

case "rpz_drop_ip_rule_enabled" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                     = "{{random}}"
      rpz_drop_ip_rule_enabled = true
    }
    check = {
      "nios.rpz_drop_ip_rule_enabled" = "true"
    }
  }

  step {
    nios {
      name                     = "{{random}}"
      rpz_drop_ip_rule_enabled = false
    }
    check = {
      "nios.rpz_drop_ip_rule_enabled" = "false"
    }
  }

}

case "rpz_drop_ip_rule_min_prefix_length_ipv4" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                                    = "{{random}}"
      rpz_drop_ip_rule_min_prefix_length_ipv4 = 30
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv4" = "30"
    }
  }

  step {
    nios {
      name                                    = "{{random}}"
      rpz_drop_ip_rule_min_prefix_length_ipv4 = 25
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv4" = "25"
    }
  }

}

case "rpz_drop_ip_rule_min_prefix_length_ipv6" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                                    = "{{random}}"
      rpz_drop_ip_rule_min_prefix_length_ipv6 = 64
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv6" = "64"
    }
  }

  step {
    nios {
      name                                    = "{{random}}"
      rpz_drop_ip_rule_min_prefix_length_ipv6 = 48
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv6" = "48"
    }
  }

}

case "rpz_qname_wait_recurse" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                   = "{{random}}"
      rpz_qname_wait_recurse = true
    }
    check = {
      "nios.rpz_qname_wait_recurse" = "true"
    }
  }

  step {
    nios {
      name                   = "{{random}}"
      rpz_qname_wait_recurse = false
    }
    check = {
      "nios.rpz_qname_wait_recurse" = "false"
    }
  }

}

case "scavenging_settings" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      scavenging_settings = { enable_scavenging = true, expression_list = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "rtype", op1_type = "FIELD", op2 = "A", op2_type = "STRING" }, { op = "ENDLIST" }] }
    }
    check = {
      "nios.scavenging_settings.enable_auto_reclamation"     = "false"
      "nios.scavenging_settings.enable_recurrent_scavenging" = "false"
      "nios.scavenging_settings.enable_rr_last_queried"      = "false"
      "nios.scavenging_settings.enable_scavenging"           = "true"
      "nios.scavenging_settings.enable_zone_last_queried"    = "false"
      "nios.scavenging_settings.reclaim_associated_records"  = "false"
      "nios.scavenging_settings.expression_list.0.op"        = "AND"
      "nios.scavenging_settings.expression_list.0.op1_type"  = "LIST"
      "nios.scavenging_settings.expression_list.1.op"        = "EQ"
      "nios.scavenging_settings.expression_list.1.op1"       = "rtype"
      "nios.scavenging_settings.expression_list.1.op1_type"  = "FIELD"
      "nios.scavenging_settings.expression_list.1.op2"       = "A"
      "nios.scavenging_settings.expression_list.1.op2_type"  = "STRING"
      "nios.scavenging_settings.expression_list.2.op"        = "ENDLIST"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      scavenging_settings = { enable_scavenging = true, expression_list = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "rtype", op1_type = "FIELD", op2 = "AAAA", op2_type = "STRING" }, { op = "ENDLIST" }] }
    }
    check = {
      "nios.scavenging_settings.enable_auto_reclamation"     = "false"
      "nios.scavenging_settings.enable_recurrent_scavenging" = "false"
      "nios.scavenging_settings.enable_rr_last_queried"      = "false"
      "nios.scavenging_settings.enable_scavenging"           = "true"
      "nios.scavenging_settings.enable_zone_last_queried"    = "false"
      "nios.scavenging_settings.reclaim_associated_records"  = "false"
      "nios.scavenging_settings.expression_list.0.op"        = "AND"
      "nios.scavenging_settings.expression_list.0.op1_type"  = "LIST"
      "nios.scavenging_settings.expression_list.1.op"        = "EQ"
      "nios.scavenging_settings.expression_list.1.op1"       = "rtype"
      "nios.scavenging_settings.expression_list.1.op1_type"  = "FIELD"
      "nios.scavenging_settings.expression_list.1.op2"       = "AAAA"
      "nios.scavenging_settings.expression_list.1.op2_type"  = "STRING"
      "nios.scavenging_settings.expression_list.2.op"        = "ENDLIST"
    }
  }

}

case "sortlist" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      sortlist = [{
        address = "13.0.0.0/24"
      }]
    }
    check = {
      "nios.sortlist.0.address" = "13.0.0.0/24"
    }
  }

  step {
    nios {
      name = "{{random}}"
      sortlist = [{
        address = "10.0.0.0/24"
      }]
    }
    check = {
      "nios.sortlist.0.address" = "10.0.0.0/24"
    }
  }

}
