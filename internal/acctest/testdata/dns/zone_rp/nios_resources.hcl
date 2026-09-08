# Auto-generated resource acceptance-test cases for ZoneRp.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "default"
    }
    check = {
      "nios.fqdn"                                    = "{{random}}.com"
      "nios.view"                                    = "default"
      "nios.disable"                                 = "false"
      "nios.locked"                                  = "false"
      "nios.log_rpz"                                 = "true"
      "nios.rpz_drop_ip_rule_enabled"                = "false"
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv4" = "29"
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv6" = "112"
      "nios.rpz_severity"                            = "MAJOR"
      "nios.rpz_type"                                = "LOCAL"
      "nios.use_external_primary"                    = "false"
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

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      comment = "Comment for ZONE RP"
    }
    check = {
      "nios.comment" = "Comment for ZONE RP"
    }
  }

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      comment = "Updated Comment for ZONE RP"
    }
    check = {
      "nios.comment" = "Updated Comment for ZONE RP"
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
      grid_secondaries     = [{ name = "{{grid_member_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.#"                   = "1"
      "nios.external_primaries.0.address"           = "10.0.0.0"
      "nios.external_primaries.0.name"              = "example-server"
      "nios.external_primaries.0.use_tsig_key_name" = "true"
      "nios.external_primaries.0.tsig_key_name"     = "{{random2}}"
      "nios.external_primaries.0.tsig_key"          = "X4oRe92t54I+T98NdQpV2w=="
      "nios.external_primaries.0.tsig_key_alg"      = "HMAC-SHA256"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      external_primaries   = [{ address = "10.0.0.2", name = "example-updated-server" }]
      grid_secondaries     = [{ name = "{{grid_member_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.0.address" = "10.0.0.2"
      "nios.external_primaries.0.name"    = "example-updated-server"
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
      grid_primary         = [{ name = "{{grid_member_hostname}}" }]
      external_secondaries = [{ address = "10.0.0.0", name = "example.com", tsig_key_alg = "HMAC-SHA256", tsig_key = "X4oRe92t54I+T98NdQpV2w==", use_tsig_key_name = false, tsig_key_name = "{{random}}" }]
    }
    check = {
      "nios.external_secondaries.#"                   = "1"
      "nios.external_secondaries.0.address"           = "10.0.0.0"
      "nios.external_secondaries.0.name"              = "example.com"
      "nios.external_secondaries.0.use_tsig_key_name" = "false"
      "nios.external_secondaries.0.tsig_key"          = "X4oRe92t54I+T98NdQpV2w=="
      "nios.external_secondaries.0.tsig_key_alg"      = "HMAC-SHA256"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      grid_primary         = [{ name = "{{grid_member_hostname}}" }]
      external_secondaries = [{ address = "10.0.0.2", name = "updated-example.com" }]
    }
    check = {
      "nios.external_secondaries.#"         = "1"
      "nios.external_secondaries.0.address" = "10.0.0.2"
      "nios.external_secondaries.0.name"    = "updated-example.com"
    }
  }

}

case "fireeye_rule_mapping" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      fireeye_rule_mapping = { apt_override = "PASSTHRU", fireeye_alert_mapping = [{ alert_type = "DOMAIN_MATCH", lifetime = "86400", rpz_rule = "NODATA" }, { alert_type = "INFECTION_MATCH", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "MALWARE_CALLBACK", lifetime = "604800", rpz_rule = "PASSTHRU" }, { alert_type = "MALWARE_OBJECT", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "WEB_INFECTION", lifetime = "86400", rpz_rule = "PASSTHRU" }] }
      rpz_type             = "FIREEYE"
    }
    check = {
      "nios.fireeye_rule_mapping.apt_override"                       = "PASSTHRU"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.#"            = "5"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.0.alert_type" = "DOMAIN_MATCH"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.0.lifetime"   = "86400"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.0.rpz_rule"   = "NODATA"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      fireeye_rule_mapping = { apt_override = "NOOVERRIDE", fireeye_alert_mapping = [{ alert_type = "DOMAIN_MATCH", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "INFECTION_MATCH", lifetime = "0", rpz_rule = "NONE" }, { alert_type = "MALWARE_CALLBACK", lifetime = "604800", rpz_rule = "PASSTHRU" }, { alert_type = "MALWARE_OBJECT", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "WEB_INFECTION", lifetime = "86400", rpz_rule = "PASSTHRU" }] }
      rpz_type             = "FIREEYE"
    }
    check = {
      "nios.fireeye_rule_mapping.apt_override"                       = "NOOVERRIDE"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.#"            = "5"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.1.alert_type" = "INFECTION_MATCH"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.1.lifetime"   = "0"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.1.rpz_rule"   = "NONE"
    }
  }

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      fireeye_rule_mapping = { apt_override = "NXDOMAIN", fireeye_alert_mapping = [{ alert_type = "DOMAIN_MATCH", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "INFECTION_MATCH", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "MALWARE_CALLBACK", lifetime = "172800", rpz_rule = "NXDOMAIN" }, { alert_type = "MALWARE_OBJECT", lifetime = "86400", rpz_rule = "PASSTHRU" }, { alert_type = "WEB_INFECTION", lifetime = "86400", rpz_rule = "PASSTHRU" }] }
      rpz_type             = "FIREEYE"
    }
    check = {
      "nios.fireeye_rule_mapping.apt_override"                       = "NXDOMAIN"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.#"            = "5"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.2.alert_type" = "MALWARE_CALLBACK"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.2.lifetime"   = "172800"
      "nios.fireeye_rule_mapping.fireeye_alert_mapping.2.rpz_rule"   = "NXDOMAIN"
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
      grid_primary = [{ name = "{{grid_member_hostname}}", stealth = false }]
    }
    check = {
      "nios.grid_primary.#"      = "1"
      "nios.grid_primary.0.name" = "{{grid_member_hostname}}"
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

case "log_rpz" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      log_rpz = false
    }
    check = {
      "nios.log_rpz" = "false"
    }
  }

  step {
    nios {
      fqdn    = "{{random}}.com"
      view    = "default"
      log_rpz = true
    }
    check = {
      "nios.log_rpz" = "true"
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
      member_soa_mnames = [{ grid_primary = "{{grid_master_hostname}}", mname = "example.com" }]
    }
    check = {
      "nios.member_soa_mnames.0.grid_primary" = "{{grid_master_hostname}}"
      "nios.member_soa_mnames.0.mname"        = "example.com"
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
      "nios.member_soa_mnames.0.mname" = "example.com"
    }
  }

}

case "ns_group" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_nsgroup" "test" {
    nios = {
      name         = "{{random}}-nsg"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }
  resource "infoblox_nsgroup" "test_updated" {
    nios = {
      name         = "{{random}}-nsg-updated"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }
  PREREQ

  step {
    nios {
      fqdn     = "{{random}}.com"
      view     = "default"
      ns_group = "${infoblox_nsgroup.test.nios.name}"
    }
    check = {
      "nios.ns_group" = "{{random}}-nsg"
    }
  }

  step {
    nios {
      fqdn     = "{{random}}.com"
      view     = "default"
      ns_group = "${infoblox_nsgroup.test_updated.nios.name}"
    }
    check = {
      "nios.ns_group" = "{{random}}-nsg-updated"
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
      prefix = "STUB-b"
    }
    check = {
      "nios.prefix" = "STUB-b"
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

case "rpz_drop_ip_rule_enabled" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                     = "{{random}}.com"
      view                     = "default"
      rpz_drop_ip_rule_enabled = true
    }
    check = {
      "nios.rpz_drop_ip_rule_enabled" = "true"
    }
  }

  step {
    nios {
      fqdn                     = "{{random}}.com"
      view                     = "default"
      rpz_drop_ip_rule_enabled = false
    }
    check = {
      "nios.rpz_drop_ip_rule_enabled" = "false"
    }
  }

}

case "rpz_drop_ip_rule_min_prefix_length_ipv4" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                                    = "{{random}}.com"
      view                                    = "default"
      rpz_drop_ip_rule_min_prefix_length_ipv4 = 20
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv4" = "20"
    }
  }

  step {
    nios {
      fqdn                                    = "{{random}}.com"
      view                                    = "default"
      rpz_drop_ip_rule_min_prefix_length_ipv4 = 30
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv4" = "30"
    }
  }

}

case "rpz_drop_ip_rule_min_prefix_length_ipv6" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                                    = "{{random}}.com"
      view                                    = "default"
      rpz_drop_ip_rule_min_prefix_length_ipv6 = 40
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv6" = "40"
    }
  }

  step {
    nios {
      fqdn                                    = "{{random}}.com"
      view                                    = "default"
      rpz_drop_ip_rule_min_prefix_length_ipv6 = 50
    }
    check = {
      "nios.rpz_drop_ip_rule_min_prefix_length_ipv6" = "50"
    }
  }

}

case "rpz_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      rpz_policy      = "DISABLED"
      substitute_name = "substitute.fqdn"
    }
    check = {
      "nios.rpz_policy" = "DISABLED"
    }
  }

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      rpz_policy      = "NODATA"
      substitute_name = "substitute.fqdn"
    }
    check = {
      "nios.rpz_policy" = "NODATA"
    }
  }

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      rpz_policy      = "PASSTHRU"
      substitute_name = "substitute.fqdn"
    }
    check = {
      "nios.rpz_policy" = "PASSTHRU"
    }
  }

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      rpz_policy      = "SUBSTITUTE"
      substitute_name = "substitute.fqdn"
    }
    check = {
      "nios.rpz_policy" = "SUBSTITUTE"
    }
  }

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      rpz_policy      = "NXDOMAIN"
      substitute_name = "substitute.fqdn"
    }
    check = {
      "nios.rpz_policy" = "NXDOMAIN"
    }
  }

}

case "rpz_severity" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      rpz_severity = "CRITICAL"
    }
    check = {
      "nios.rpz_severity" = "CRITICAL"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      rpz_severity = "INFORMATIONAL"
    }
    check = {
      "nios.rpz_severity" = "INFORMATIONAL"
    }
  }

  step {
    nios {
      fqdn         = "{{random}}.com"
      view         = "default"
      rpz_severity = "WARNING"
    }
    check = {
      "nios.rpz_severity" = "WARNING"
    }
  }

}

case "rpz_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                 = "{{random}}.com"
      view                 = "default"
      rpz_type             = "FEED"
      external_primaries   = [{ address = "10.0.0.0", name = "example-server" }]
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.rpz_type" = "FEED"
    }
  }

}

case "set_soa_serial_number" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn                  = "{{random}}.com"
      view                  = "default"
      set_soa_serial_number = true
    }
    check = {
      "nios.set_soa_serial_number" = "true"
    }
  }

  step {
    nios {
      fqdn                  = "{{random}}.com"
      view                  = "default"
      set_soa_serial_number = false
    }
    check = {
      "nios.set_soa_serial_number" = "false"
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
      grid_primary = [{ name = "{{grid_member_hostname}}" }]
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
      grid_primary = [{ name = "{{grid_member_hostname}}" }]
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

case "substitute_name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      substitute_name = "alternate.fqdn"
      rpz_policy      = "SUBSTITUTE"
    }
    check = {
      "nios.substitute_name" = "alternate.fqdn"
    }
  }

  step {
    nios {
      fqdn            = "{{random}}.com"
      view            = "default"
      substitute_name = "updated-Alternate.fqdn"
      rpz_policy      = "SUBSTITUTE"
    }
    check = {
      "nios.substitute_name" = "updated-Alternate.fqdn"
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
      external_primaries   = [{ address = "10.0.0.0", name = "example-server" }]
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.use_external_primary" = "true"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    nios = {
      name = "test_view"
    }
  }
  PREREQ

  step {
    nios {
      fqdn = "{{random}}.com"
      view = "test_view"
    }
    depends_on = [infoblox_view.test_view]
    check = {
      "nios.view" = "test_view"
    }
  }

}
