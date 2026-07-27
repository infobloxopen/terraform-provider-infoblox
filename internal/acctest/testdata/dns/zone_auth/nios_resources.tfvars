# Auto-generated resource acceptance-test cases for ZoneAuth.
// TODO : Objects to be present in the grid for testing
// GSS TSIG Key has to be configured in the grid
// Microsoft Servers 10.10.10.10, 10.0.0.0, example_server
// -NS Group - example-ns-group, updated-example-ns-group
// Shared Record Group - example_shared_record_group, updated_example_shared_record_group
// DDNS Principal Cluster Group - dynamic_update_grp_1, dynamic_update_grp_2

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
    check = {
      "nios.fqdn"                                 = "{{random}}.com"
      "nios.view"                                 = "default"
      "nios.allow_fixed_rrset_order"              = "false"
      "nios.allow_gss_tsig_for_underscore_zone"   = "false"
      "nios.allow_gss_tsig_zone_updates"          = "false"
      "nios.copy_xfer_to_notify"                  = "false"
      "nios.create_underscore_zones"              = "false"
      "nios.ddns_force_creation_timestamp_update" = "false"
      "nios.ddns_principal_tracking"              = "false"
      "nios.ddns_restrict_patterns"               = "false"
      "nios.ddns_restrict_protected"              = "false"
      "nios.ddns_restrict_secure"                 = "false"
      "nios.ddns_restrict_static"                 = "false"
      "nios.disable"                              = "false"
      "nios.disable_forwarding"                   = "false"
      "nios.dns_integrity_enable"                 = "false"
      "nios.dns_integrity_frequency"              = "3600"
      "nios.dns_integrity_verbose_logging"        = "false"
      "nios.effective_check_names_policy"         = "WARN"
      "nios.locked"                               = "false"
      "nios.ms_ad_integrated"                     = "false"
      "nios.ms_allow_transfer_mode"               = "NONE"
      "nios.ms_ddns_mode"                         = "NONE"
      "nios.ms_sync_disabled"                     = "false"
      "nios.notify_delay"                         = "5"
      "nios.use_check_names_policy"               = "false"
      "nios.use_external_primary"                 = "false"
      "nios.use_import_from"                      = "false"
      "nios.zone_format"                          = "FORWARD"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
  }

}

case "allow_active_dir" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      allow_active_dir = [{ address = "10.0.0.1", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_active_dir.#"            = "1"
      "nios.allow_active_dir.0.address"    = "10.0.0.1"
      "nios.allow_active_dir.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      allow_active_dir = [{ address = "10.0.0.2", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_active_dir.#"            = "1"
      "nios.allow_active_dir.0.address"    = "10.0.0.2"
      "nios.allow_active_dir.0.permission" = "ALLOW"
    }
  }

}

case "allow_fixed_rrset_order" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      allow_fixed_rrset_order = false
    }
    check = {
      "nios.allow_fixed_rrset_order" = "false"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      allow_fixed_rrset_order = true
    }
    check = {
      "nios.allow_fixed_rrset_order" = "true"
    }
  }

}

case "allow_gss_tsig_for_underscore_zone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                               = "{{random}}.com"
      view                               = "default"
      allow_gss_tsig_for_underscore_zone = false
    }
    check = {
      "nios.allow_gss_tsig_for_underscore_zone" = "false"
    }
  }

  step {
    nios {
      fqdn                               = "{{random}}.com"
      view                               = "default"
      allow_gss_tsig_for_underscore_zone = true
    }
    check = {
      "nios.allow_gss_tsig_for_underscore_zone" = "true"
    }
  }

}

case "allow_gss_tsig_zone_updates" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                        = "{{random}}.com"
      view                        = "default"
      allow_gss_tsig_zone_updates = true
      allow_fixed_rrset_order     = false
    }
    check = {
      "nios.allow_gss_tsig_zone_updates" = "true"
    }
  }

  step {
    nios {
      fqdn                        = "{{random}}.com"
      view                        = "default"
      allow_gss_tsig_zone_updates = false
      allow_fixed_rrset_order     = false
    }
    check = {
      "nios.allow_gss_tsig_zone_updates" = "false"
    }
  }

}

case "allow_query" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn        = "{{random}}.com"
      view        = "default"
      allow_query = [{ struct = "addressac", address = "10.0.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_query.0.address"    = "10.0.0.0"
      "nios.allow_query.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      fqdn        = "{{random}}.com"
      view        = "default"
      allow_query = [{ struct = "addressac", address = "192.168.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_query.0.address"    = "192.168.0.0"
      "nios.allow_query.0.permission" = "ALLOW"
    }
  }

}

case "allow_transfer" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn           = "{{random}}.com"
      view           = "default"
      allow_transfer = [{ struct = "addressac", address = "10.0.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_transfer.0.address"    = "10.0.0.0"
      "nios.allow_transfer.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      fqdn           = "{{random}}.com"
      view           = "default"
      allow_transfer = [{ struct = "addressac", address = "192.168.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_transfer.0.address"    = "192.168.0.0"
      "nios.allow_transfer.0.permission" = "ALLOW"
    }
  }

}

case "allow_update" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      allow_update = [{ struct = "addressac", address = "10.0.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_update.#"            = "1"
      "nios.allow_update.0.struct"     = "addressac"
      "nios.allow_update.0.address"    = "10.0.0.0"
      "nios.allow_update.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      allow_update = [{ struct = "addressac", address = "192.168.0.0", permission = "ALLOW" }]
    }
    check = {
      "nios.allow_update.#"            = "1"
      "nios.allow_update.0.struct"     = "addressac"
      "nios.allow_update.0.address"    = "192.168.0.0"
      "nios.allow_update.0.permission" = "ALLOW"
    }
  }

}

case "allow_update_forwarding" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      allow_update_forwarding = true
    }
    check = {
      "nios.allow_update_forwarding" = "true"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      allow_update_forwarding = false
    }
    check = {
      "nios.allow_update_forwarding" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      comment = "initial comment"
    }
    check = {
      "nios.comment" = "initial comment"
    }
  }

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      comment = "updated comment"
    }
    check = {
      "nios.comment" = "updated comment"
    }
  }

}

case "copy_xfer_to_notify" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                = "{{random}}.com"
      view                = "default"
      copy_xfer_to_notify = true
    }
    check = {
      "nios.copy_xfer_to_notify" = "true"
    }
  }

  step {
    nios {
      fqdn                = "{{random}}.com"
      view                = "default"
      copy_xfer_to_notify = false
    }
    check = {
      "nios.copy_xfer_to_notify" = "false"
    }
  }

}

case "create_underscore_zones" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      create_underscore_zones = true
    }
    check = {
      "nios.create_underscore_zones" = "true"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      create_underscore_zones = false
    }
    check = {
      "nios.create_underscore_zones" = "false"
    }
  }

}

case "ddns_force_creation_timestamp_update" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                                 = "{{random}}.com"
      view                                 = "default"
      ddns_force_creation_timestamp_update = true
    }
    check = {
      "nios.ddns_force_creation_timestamp_update" = "true"
    }
  }

  step {
    nios {
      fqdn                                 = "{{random}}.com"
      view                                 = "default"
      ddns_force_creation_timestamp_update = false
    }
    check = {
      "nios.ddns_force_creation_timestamp_update" = "false"
    }
  }

}

case "ddns_principal_group" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      ddns_principal_group = "dynamic_update_grp_1"
    }
    check = {
      "nios.ddns_principal_group" = "dynamic_update_grp_1"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      ddns_principal_group = "dynamic_update_grp_2"
    }
    check = {
      "nios.ddns_principal_group" = "dynamic_update_grp_2"
    }
  }

}

case "ddns_principal_tracking" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      ddns_principal_tracking = true
    }
    check = {
      "nios.ddns_principal_tracking" = "true"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      ddns_principal_tracking = false
    }
    check = {
      "nios.ddns_principal_tracking" = "false"
    }
  }

}

case "ddns_restrict_patterns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                   = "{{random}}.com"
      view                   = "default"
      ddns_restrict_patterns = true
    }
    check = {
      "nios.ddns_restrict_patterns" = "true"
    }
  }

  step {
    nios {
      fqdn                   = "{{random}}.com"
      view                   = "default"
      ddns_restrict_patterns = false
    }
    check = {
      "nios.ddns_restrict_patterns" = "false"
    }
  }

}

case "ddns_restrict_patterns_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                        = "{{random}}.com"
      view                        = "default"
      ddns_restrict_patterns_list = ["pattern1", "pattern2"]
    }
    check = {
      "nios.ddns_restrict_patterns_list.#" = "2"
      "nios.ddns_restrict_patterns_list.0" = "pattern1"
      "nios.ddns_restrict_patterns_list.1" = "pattern2"
    }
  }

  step {
    nios {
      fqdn                        = "{{random}}.com"
      view                        = "default"
      ddns_restrict_patterns_list = ["pattern3", "pattern4"]
    }
    check = {
      "nios.ddns_restrict_patterns_list.#" = "2"
      "nios.ddns_restrict_patterns_list.0" = "pattern3"
      "nios.ddns_restrict_patterns_list.1" = "pattern4"
    }
  }

}

case "ddns_restrict_protected" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      ddns_restrict_protected = true
    }
    check = {
      "nios.ddns_restrict_protected" = "true"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      ddns_restrict_protected = false
    }
    check = {
      "nios.ddns_restrict_protected" = "false"
    }
  }

}

case "ddns_restrict_secure" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      ddns_restrict_secure = true
    }
    check = {
      "nios.ddns_restrict_secure" = "true"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      ddns_restrict_secure = false
    }
    check = {
      "nios.ddns_restrict_secure" = "false"
    }
  }

}

case "ddns_restrict_static" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      ddns_restrict_static = true
    }
    check = {
      "nios.ddns_restrict_static" = "true"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      ddns_restrict_static = false
    }
    check = {
      "nios.ddns_restrict_static" = "false"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      disable = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "disable_forwarding" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn               = "{{random}}.com"
      view               = "default"
      disable_forwarding = true
    }
    check = {
      "nios.disable_forwarding" = "true"
    }
  }

  step {
    nios {
      fqdn               = "{{random}}.com"
      view               = "default"
      disable_forwarding = false
    }
    check = {
      "nios.disable_forwarding" = "false"
    }
  }

}

case "dns_integrity_enable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      dns_integrity_enable = true
      dns_integrity_member = "{{grid_master_hostname}}"
      grid_primary         = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.dns_integrity_enable" = "true"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      dns_integrity_enable = false
      dns_integrity_member = "{{grid_master_hostname}}"
      grid_primary         = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.dns_integrity_enable" = "false"
    }
  }

}

case "dns_integrity_frequency" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      dns_integrity_frequency = 1000
    }
    check = {
      "nios.dns_integrity_frequency" = "1000"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      dns_integrity_frequency = 2000
    }
    check = {
      "nios.dns_integrity_frequency" = "2000"
    }
  }

}

case "dns_integrity_member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      dns_integrity_member = "{{grid_master_hostname}}"
    }
    check = {
      "nios.dns_integrity_member" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      dns_integrity_member = "{{grid_master_hostname}}"
    }
    check = {
      "nios.dns_integrity_member" = "{{grid_master_hostname}}"
    }
  }

}

case "dns_integrity_verbose_logging" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                          = "{{random}}.com"
      view                          = "default"
      dns_integrity_verbose_logging = true
    }
    check = {
      "nios.dns_integrity_verbose_logging" = "true"
    }
  }

  step {
    nios {
      fqdn                          = "{{random}}.com"
      view                          = "default"
      dns_integrity_verbose_logging = false
    }
    check = {
      "nios.dns_integrity_verbose_logging" = "false"
    }
  }

}

case "dnssec_key_params" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn              = "{{random}}.com"
      view              = "default"
      dnssec_key_params = { ksk_algorithms = [{ algorithm = "RSASHA256", size = 2048 }], zsk_algorithms = [{ algorithm = "RSASHA256", size = 1024 }] }
    }
    check = {
      "nios.dnssec_key_params.ksk_algorithms.#"           = "1"
      "nios.dnssec_key_params.ksk_algorithms.0.algorithm" = "RSASHA256"
      "nios.dnssec_key_params.ksk_algorithms.0.size"      = "2048"
      "nios.dnssec_key_params.zsk_algorithms.#"           = "1"
      "nios.dnssec_key_params.zsk_algorithms.0.algorithm" = "RSASHA256"
      "nios.dnssec_key_params.zsk_algorithms.0.size"      = "1024"
    }
  }

  step {
    nios {
      fqdn              = "{{random}}.com"
      view              = "default"
      dnssec_key_params = { ksk_algorithms = [{ algorithm = "RSASHA512", size = 4096 }], zsk_algorithms = [{ algorithm = "RSASHA512", size = 2048 }] }
    }
    check = {
      "nios.dnssec_key_params.ksk_algorithms.#"           = "1"
      "nios.dnssec_key_params.ksk_algorithms.0.algorithm" = "RSASHA512"
      "nios.dnssec_key_params.ksk_algorithms.0.size"      = "4096"
      "nios.dnssec_key_params.zsk_algorithms.#"           = "1"
      "nios.dnssec_key_params.zsk_algorithms.0.algorithm" = "RSASHA512"
      "nios.dnssec_key_params.zsk_algorithms.0.size"      = "2048"
    }
  }

}

case "effective_check_names_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                         = "{{random}}.com"
      view                         = "default"
      effective_check_names_policy = "WARN"
    }
    check = {
      "nios.effective_check_names_policy" = "WARN"
    }
  }

  step {
    nios {
      fqdn                         = "{{random}}.com"
      view                         = "default"
      effective_check_names_policy = "FAIL"
    }
    check = {
      "nios.effective_check_names_policy" = "FAIL"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn      = "{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      fqdn      = "{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "external_primaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      external_primaries   = [{ address = "10.0.0.0", name = "example-server", tsig_key_alg = "HMAC-SHA256", tsig_key = "X4oRe92t54I+T98NdQpV2w==", use_tsig_key_name = true, tsig_key_name = "{{random2}}" }]
      ms_secondaries       = [{ address = "10.10.10.10", ns_name = "example-server", ns_ip = "1.1.1.1" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.#"         = "1"
      "nios.external_primaries.0.address" = "10.0.0.0"
      "nios.external_primaries.0.name"    = "example-server"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      external_primaries   = [{ address = "10.0.0.2", name = "example-server" }]
      ms_secondaries       = [{ address = "10.10.10.10", ns_name = "example-server", ns_ip = "1.1.1.1" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.0.address" = "10.0.0.2"
      "nios.external_primaries.0.name"    = "example-server"
    }
  }

}

case "external_secondaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      grid_primary         = [{ name = "{{grid_master_hostname}}" }]
      external_secondaries = [{ address = "10.0.0.0", name = "example.com", tsig_key_alg = "HMAC-SHA256", tsig_key = "X4oRe92t54I+T98NdQpV2w==", use_tsig_key_name = false, tsig_key_name = "{{random}}" }]
    }
    check = {
      "nios.external_secondaries.#"         = "1"
      "nios.external_secondaries.0.address" = "10.0.0.0"
      "nios.external_secondaries.0.name"    = "example.com"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      grid_primary         = [{ name = "{{grid_master_hostname}}" }]
      external_secondaries = [{ address = "10.0.0.2", name = "updated-example.com" }]
    }
    check = {
      "nios.external_secondaries.#"         = "1"
      "nios.external_secondaries.0.address" = "10.0.0.2"
      "nios.external_secondaries.0.name"    = "updated-example.com"
    }
  }

}

case "fqdn" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
    check = {
      "nios.fqdn" = "{{random}}.com"
    }
  }

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
    check = {
      "nios.fqdn" = "{{random}}.com"
    }
  }

}

case "grid_primary" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}", stealth = false }]
    }
    check = {
      "nios.grid_primary.#"         = "1"
      "nios.grid_primary.0.name"    = "{{grid_master_hostname}}"
      "nios.grid_primary.0.stealth" = "false"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}", stealth = true }, { name = "{{grid_member_hostname}}", stealth = false }]
    }
    check = {
      "nios.grid_primary.#"         = "2"
      "nios.grid_primary.0.name"    = "{{grid_master_hostname}}"
      "nios.grid_primary.0.stealth" = "true"
      "nios.grid_primary.1.name"    = "{{grid_member_hostname}}"
      "nios.grid_primary.1.stealth" = "false"
    }
  }

}

case "grid_secondaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      grid_secondaries = [{ name = "{{grid_member_hostname}}", stealth = false, grid_replicate = true, lead = false, enable_preferred_primaries = false }]
    }
    check = {
      "nios.grid_secondaries.#"                            = "1"
      "nios.grid_secondaries.0.name"                       = "{{grid_member_hostname}}"
      "nios.grid_secondaries.0.stealth"                    = "false"
      "nios.grid_secondaries.0.grid_replicate"             = "true"
      "nios.grid_secondaries.0.lead"                       = "false"
      "nios.grid_secondaries.0.enable_preferred_primaries" = "false"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_member_hostname}}" }]
      grid_secondaries = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.grid_secondaries.#"      = "1"
      "nios.grid_secondaries.0.name" = "{{grid_master_hostname}}"
    }
  }

}

case "last_queried_acl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      last_queried_acl = [{ address = "10.0.0.0" }]
    }
    check = {
      "nios.last_queried_acl.#"         = "1"
      "nios.last_queried_acl.0.address" = "10.0.0.0"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      last_queried_acl = [{ address = "10.0.0.2" }]
    }
    check = {
      "nios.last_queried_acl.#"         = "1"
      "nios.last_queried_acl.0.address" = "10.0.0.2"
    }
  }

}

case "locked" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn   = "{{random}}.com"
      view   = "default"
      locked = true
    }
    check = {
      "nios.locked" = "true"
    }
  }

  step {
    nios {
      fqdn   = "{{random}}.com"
      view   = "default"
      locked = false
    }
    check = {
      "nios.locked" = "false"
    }
  }

}

case "member_soa_mnames" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn              = "{{random}}.com"
      view              = "default"
      grid_primary      = [{ name = "{{grid_master_hostname}}" }]
      member_soa_mnames = [{ grid_primary = "{{grid_master_hostname}}", mname = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.member_soa_mnames.#"              = "1"
      "nios.member_soa_mnames.0.grid_primary" = "{{grid_master_hostname}}"
      "nios.member_soa_mnames.0.mname"        = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      fqdn              = "{{random}}.com"
      view              = "default"
      grid_primary      = [{ name = "{{grid_master_hostname}}" }]
      member_soa_mnames = [{ mname = "example.com" }]
    }
    check = {
      "nios.member_soa_mnames.#"              = "1"
      "nios.member_soa_mnames.0.grid_primary" = "{{grid_master_hostname}}"
      "nios.member_soa_mnames.0.mname"        = "example.com"
    }
  }

}

case "ms_ad_integrated" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      ms_ad_integrated = true
    }
    check = {
      "nios.ms_ad_integrated" = "true"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      ms_ad_integrated = false
    }
    check = {
      "nios.ms_ad_integrated" = "false"
    }
  }

}

case "ms_allow_transfer" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn              = "{{random}}.com"
      view              = "default"
      ms_allow_transfer = [{ address = "192.168.1.10", permission = "ALLOW" }]
    }
    check = {
      "nios.ms_allow_transfer.#"            = "1"
      "nios.ms_allow_transfer.0.address"    = "192.168.1.10"
      "nios.ms_allow_transfer.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      fqdn              = "{{random}}.com"
      view              = "default"
      ms_allow_transfer = [{ address = "192.168.1.20", permission = "DENY" }]
    }
    check = {
      "nios.ms_allow_transfer.#"            = "1"
      "nios.ms_allow_transfer.0.address"    = "192.168.1.20"
      "nios.ms_allow_transfer.0.permission" = "DENY"
    }
  }

}

case "ms_allow_transfer_mode" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                   = "{{random}}.com"
      view                   = "default"
      ms_allow_transfer_mode = "ANY"
    }
    check = {
      "nios.ms_allow_transfer_mode" = "ANY"
    }
  }

  step {
    nios {
      fqdn                   = "{{random}}.com"
      view                   = "default"
      ms_allow_transfer_mode = "NONE"
    }
    check = {
      "nios.ms_allow_transfer_mode" = "NONE"
    }
  }

}

case "ms_dc_ns_record_creation" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                     = "{{random}}.com"
      view                     = "default"
      ms_dc_ns_record_creation = [{ address = "10.0.0.0" }]
      ms_ad_integrated         = true
    }
    check = {
      "nios.ms_dc_ns_record_creation.#"         = "1"
      "nios.ms_dc_ns_record_creation.0.address" = "10.0.0.0"
      "nios.ms_ad_integrated"                   = "true"
    }
  }

  step {
    nios {
      fqdn                     = "{{random}}.com"
      view                     = "default"
      ms_dc_ns_record_creation = [{ address = "198.51.100.0" }]
      ms_ad_integrated         = true
    }
    check = {
      "nios.ms_dc_ns_record_creation.#"         = "1"
      "nios.ms_dc_ns_record_creation.0.address" = "198.51.100.0"
      "nios.ms_ad_integrated"                   = "true"
    }
  }

}

case "ms_ddns_mode" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      ms_ddns_mode = "ANY"
    }
    check = {
      "nios.ms_ddns_mode" = "ANY"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      ms_ddns_mode = "NONE"
    }
    check = {
      "nios.ms_ddns_mode" = "NONE"
    }
  }

}

case "ms_primaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      ms_primaries = [{ address = "10.10.10.10", ns_ip = "1.1.1.1", ns_name = "example-server" }]
    }
    check = {
      "nios.ms_primaries.#"         = "1"
      "nios.ms_primaries.0.address" = "10.10.10.10"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      ms_primaries = [{ address = "10.0.0.0", ns_ip = "1.1.1.1", ns_name = "example-server" }]
    }
    check = {
      "nios.ms_primaries.#"         = "1"
      "nios.ms_primaries.0.address" = "10.0.0.0"
    }
  }

}

case "ms_secondaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn           = "{{random}}.com"
      view           = "default"
      ms_secondaries = [{ address = "10.10.10.10", ns_name = "example-server", ns_ip = "1.1.1.1" }]
      grid_primary   = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.ms_secondaries.#"         = "1"
      "nios.ms_secondaries.0.address" = "10.10.10.10"
    }
  }

  step {
    nios {
      fqdn           = "{{random}}.com"
      view           = "default"
      ms_secondaries = [{ address = "example_server", ns_name = "example-server", ns_ip = "1.1.1.1" }]
      grid_primary   = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.ms_secondaries.#"         = "1"
      "nios.ms_secondaries.0.address" = "example_server"
    }
  }

}

case "ms_sync_disabled" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      ms_sync_disabled = true
      ms_primaries     = [{ address = "10.10.10.10", ns_name = "example-server", ns_ip = "1.1.1.1" }]
    }
    check = {
      "nios.ms_sync_disabled" = "true"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      ms_sync_disabled = false
      ms_primaries     = [{ address = "10.10.10.10", ns_name = "example-server", ns_ip = "1.1.1.1" }]
    }
    check = {
      "nios.ms_sync_disabled" = "false"
    }
  }

}

case "notify_delay" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      notify_delay = 5
    }
    check = {
      "nios.notify_delay" = "5"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      notify_delay = 20
    }
    check = {
      "nios.notify_delay" = "20"
    }
  }

}

case "ns_group" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ns_group_unknown" "test_ns_group" {
  #   nios = {
  #     name = "example-ns-group"
  #   }
  # }
  # resource "infoblox_ns_group_unknown" "test_ns_group_updated" {
  #   nios = {
  #     name = "updated-example-ns-group"
  #   }
  # }
  # PREREQ

  step {
    nios {
      fqdn     = "{{random}}.com"
      view     = "default"
      ns_group = "example-ns-group"
    }
    depends_on = [infoblox_ns_group_unknown.test_ns_group, infoblox_ns_group_unknown.test_ns_group_updated]
    check = {
      "nios.ns_group" = "example-ns-group"
    }
  }

  step {
    nios {
      fqdn     = "{{random}}.com"
      view     = "default"
      ns_group = "updated-example-ns-group"
    }
    depends_on = [infoblox_ns_group_unknown.test_ns_group, infoblox_ns_group_unknown.test_ns_group_updated]
    check = {
      "nios.ns_group" = "updated-example-ns-group"
    }
  }

}

case "prefix" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn   = "{{random}}.com"
      view   = "default"
      prefix = "128/26"
    }
    check = {
      "nios.prefix" = "128/26"
    }
  }

  step {
    nios {
      fqdn   = "{{random}}.com"
      view   = "default"
      prefix = "121/26"
    }
    check = {
      "nios.prefix" = "121/26"
    }
  }

}

case "record_name_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn               = "{{random}}.com"
      view               = "default"
      record_name_policy = "Allow Any"
    }
    check = {
      "nios.record_name_policy" = "Allow Any"
    }
  }

  step {
    nios {
      fqdn               = "{{random}}.com"
      view               = "default"
      record_name_policy = "Allow Underscore"
    }
    check = {
      "nios.record_name_policy" = "Allow Underscore"
    }
  }

}

case "scavenging_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                = "{{random}}.com"
      view                = "default"
      scavenging_settings = { enable_scavenging = true, expression_list = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "rtype", op1_type = "FIELD", op2 = "A", op2_type = "STRING" }, { op = "ENDLIST" }] }
    }
    check = {
      "nios.scavenging_settings.enable_scavenging"          = "true"
      "nios.scavenging_settings.expression_list.#"          = "3"
      "nios.scavenging_settings.expression_list.0.op"       = "AND"
      "nios.scavenging_settings.expression_list.0.op1_type" = "LIST"
      "nios.scavenging_settings.expression_list.1.op"       = "EQ"
      "nios.scavenging_settings.expression_list.1.op1"      = "rtype"
      "nios.scavenging_settings.expression_list.1.op1_type" = "FIELD"
      "nios.scavenging_settings.expression_list.1.op2"      = "A"
      "nios.scavenging_settings.expression_list.1.op2_type" = "STRING"
      "nios.scavenging_settings.expression_list.2.op"       = "ENDLIST"
    }
  }

  step {
    nios {
      fqdn                = "{{random}}.com"
      view                = "default"
      scavenging_settings = { enable_scavenging = true, expression_list = [{ op = "AND", op1_type = "LIST" }, { op = "EQ", op1 = "rtype", op1_type = "FIELD", op2 = "AAAA", op2_type = "STRING" }, { op = "ENDLIST" }] }
    }
    check = {
      "nios.scavenging_settings.enable_scavenging"          = "true"
      "nios.scavenging_settings.expression_list.#"          = "3"
      "nios.scavenging_settings.expression_list.0.op"       = "AND"
      "nios.scavenging_settings.expression_list.0.op1_type" = "LIST"
      "nios.scavenging_settings.expression_list.1.op"       = "EQ"
      "nios.scavenging_settings.expression_list.1.op1"      = "rtype"
      "nios.scavenging_settings.expression_list.1.op1_type" = "FIELD"
      "nios.scavenging_settings.expression_list.1.op2"      = "AAAA"
      "nios.scavenging_settings.expression_list.1.op2_type" = "STRING"
      "nios.scavenging_settings.expression_list.2.op"       = "ENDLIST"
    }
  }

}

case "soa_default_ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_default_ttl  = 8
      soa_expire       = 2419200
      soa_negative_ttl = 900
      soa_refresh      = 10800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_default_ttl" = "8"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_default_ttl  = 10
      soa_expire       = 2419200
      soa_negative_ttl = 900
      soa_refresh      = 10800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_default_ttl" = "10"
    }
  }

}

case "soa_email" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      soa_email    = "user1@example.com"
    }
    check = {
      "nios.soa_email" = "user1@example.com"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      soa_email    = "user2@example.com"
    }
    check = {
      "nios.soa_email" = "user2@example.com"
    }
  }

}

case "soa_expire" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_expire       = 24192
      soa_default_ttl  = 28800
      soa_negative_ttl = 900
      soa_refresh      = 10800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_expire" = "24192"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_expire       = 24100
      soa_default_ttl  = 28800
      soa_negative_ttl = 900
      soa_refresh      = 10800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_expire" = "24100"
    }
  }

}

case "soa_negative_ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_negative_ttl = 800
      soa_expire       = 2419200
      soa_default_ttl  = 28800
      soa_refresh      = 10800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_negative_ttl" = "800"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_negative_ttl = 900
      soa_expire       = 2419200
      soa_default_ttl  = 28800
      soa_refresh      = 10800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_negative_ttl" = "900"
    }
  }

}

case "soa_refresh" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_refresh      = 800
      soa_negative_ttl = 900
      soa_expire       = 2419200
      soa_default_ttl  = 28800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_refresh" = "800"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_refresh      = 900
      soa_negative_ttl = 900
      soa_expire       = 2419200
      soa_default_ttl  = 28800
      soa_retry        = 3600
    }
    check = {
      "nios.soa_refresh" = "900"
    }
  }

}

case "soa_retry" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_retry        = 1600
      soa_negative_ttl = 900
      soa_expire       = 2419200
      soa_default_ttl  = 28800
      soa_refresh      = 10800
    }
    check = {
      "nios.soa_retry" = "1600"
    }
  }

  step {
    nios {
      fqdn             = "{{random}}.com"
      view             = "default"
      grid_primary     = [{ name = "{{grid_master_hostname}}" }]
      soa_retry        = 1700
      soa_negative_ttl = 900
      soa_expire       = 2419200
      soa_default_ttl  = 28800
      soa_refresh      = 10800
    }
    check = {
      "nios.soa_retry" = "1700"
    }
  }

}

case "soa_serial_number" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                  = "{{random}}.com"
      view                  = "default"
      grid_primary          = [{ name = "{{grid_master_hostname}}" }]
      soa_serial_number     = 10
      set_soa_serial_number = true
      soa_retry             = 3600
      soa_negative_ttl      = 900
      soa_expire            = 2419200
      soa_default_ttl       = 28800
      soa_refresh           = 10800
    }
    check = {
      "nios.soa_serial_number" = "10"
    }
  }

  step {
    nios {
      fqdn                  = "{{random}}.com"
      view                  = "default"
      grid_primary          = [{ name = "{{grid_master_hostname}}" }]
      soa_serial_number     = 20
      set_soa_serial_number = true
      soa_retry             = 3600
      soa_negative_ttl      = 900
      soa_expire            = 2419200
      soa_default_ttl       = 28800
      soa_refresh           = 10800
    }
    check = {
      "nios.soa_serial_number" = "20"
    }
  }

}

case "srgs" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "test_shared_record_group" {
  #   nios = {
  #     name = "example_shared_record_group"
  #   }
  # }
  # resource "infoblox_shared_record_group_unknown" "test_shared_record_group_updated" {
  #   nios = {
  #     name = "updated_example_shared_record_group"
  #   }
  # }
  # PREREQ

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
      srgs = ["example_shared_record_group"]
    }
    # depends_on = [infoblox_shared_record_group_unknown.test_shared_record_group, infoblox_shared_record_group_unknown.test_shared_record_group_updated]
    check = {
      "nios.srgs.#" = "1"
      "nios.srgs.0" = "example_shared_record_group"
    }
  }

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
      srgs = ["updated_example_shared_record_group"]
    }
    # depends_on = [infoblox_shared_record_group_unknown.test_shared_record_group, infoblox_shared_record_group_unknown.test_shared_record_group_updated]
    check = {
      "nios.srgs.#" = "1"
      "nios.srgs.0" = "updated_example_shared_record_group"
    }
  }

}

case "update_forwarding" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      update_forwarding       = [{ struct = "addressac", address = "10.0.0.0", permission = "ALLOW" }]
      allow_update_forwarding = true
    }
    check = {
      "nios.update_forwarding.#"            = "1"
      "nios.update_forwarding.0.address"    = "10.0.0.0"
      "nios.update_forwarding.0.permission" = "ALLOW"
    }
  }

  step {
    nios {
      fqdn                    = "{{random}}.com"
      view                    = "default"
      update_forwarding       = [{ struct = "tsigac", tsig_key = "X4oRe92t54I+T98NdQpV2w==", tsig_key_name = "example-tsig-key", tsig_key_alg = "HMAC-SHA256" }]
      allow_update_forwarding = true
    }
    check = {
      "nios.update_forwarding.#"          = "1"
      "nios.update_forwarding.0.tsig_key" = "X4oRe92t54I+T98NdQpV2w=="
    }
  }

}

case "use_check_names_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                   = "{{random}}.com"
      view                   = "default"
      use_check_names_policy = true
    }
    check = {
      "nios.use_check_names_policy" = "true"
    }
  }

  step {
    nios {
      fqdn                   = "{{random}}.com"
      view                   = "default"
      use_check_names_policy = false
    }
    check = {
      "nios.use_check_names_policy" = "false"
    }
  }

}

case "use_external_primary" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      external_primaries   = [{ address = "10.0.0.0", name = "10.10.10.10" }]
      ms_secondaries       = [{ address = "10.10.10.10", ns_name = "10.10.10.10", ns_ip = "10.120.23.22" }]
      use_external_primary = true
    }
    check = {
      "nios.use_external_primary" = "true"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      use_external_primary = false
    }
    check = {
      "nios.use_external_primary" = "false"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
    check = {
      "nios.view" = "default"
    }
  }

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
    check = {
      "nios.view" = "default"
    }
  }

}

case "zone_format_ipv4" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn        = "110.0.0.0/24"
      view        = "default"
      zone_format = "IPV4"
    }
    check = {
      "nios.zone_format" = "IPV4"
    }
  }

}

case "zone_format_ipv6" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn        = "2002:5505::/64"
      view        = "default"
      zone_format = "IPV6"
    }
    check = {
      "nios.zone_format" = "IPV6"
    }
  }

}
