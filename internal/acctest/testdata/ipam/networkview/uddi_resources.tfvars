# Auto-generated resource acceptance-test cases for Networkview.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name"                          = "{{random}}"
      "uddi.ddns_client_update"            = "client"
      "uddi.ddns_conflict_resolution_mode" = "check_with_dhcid"
      "uddi.ddns_generate_name"            = "false"
      "uddi.ddns_generated_prefix"         = "myhost"
      "uddi.ddns_send_updates"             = "true"
      "uddi.ddns_update_on_renew"          = "false"
      "uddi.ddns_use_conflict_resolution"  = "true"
      "uddi.hostname_rewrite_char"         = "-"
      "uddi.hostname_rewrite_enabled"      = "false"
      "uddi.hostname_rewrite_regex"        = "[^a-zA-Z0-9_.]"
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

case "asm_config" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      asm_config = { asm_threshold = 70, enable = true, enable_notification = true, forecast_period = 12, growth_factor = 40, growth_type = "count", history = 40, min_total = 30, min_unused = 30, reenable_date = "2020-01-10T10:11:22Z" }
    }
    check = {
      "uddi.asm_config.asm_threshold"       = "70"
      "uddi.asm_config.enable"              = "true"
      "uddi.asm_config.enable_notification" = "true"
      "uddi.asm_config.forecast_period"     = "12"
      "uddi.asm_config.growth_factor"       = "40"
      "uddi.asm_config.growth_type"         = "count"
      "uddi.asm_config.history"             = "40"
      "uddi.asm_config.min_total"           = "30"
      "uddi.asm_config.min_unused"          = "30"
      "uddi.asm_config.reenable_date"       = "2020-01-10T10:11:22Z"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      asm_config = { asm_threshold = 80, enable = false, enable_notification = false, forecast_period = 10, growth_factor = 50, growth_type = "percent", history = 50, min_total = 10, min_unused = 10, reenable_date = "2021-01-10T10:11:22Z" }
    }
    check = {
      "uddi.asm_config.asm_threshold"       = "80"
      "uddi.asm_config.enable"              = "false"
      "uddi.asm_config.enable_notification" = "false"
      "uddi.asm_config.forecast_period"     = "10"
      "uddi.asm_config.growth_factor"       = "50"
      "uddi.asm_config.growth_type"         = "percent"
      "uddi.asm_config.history"             = "50"
      "uddi.asm_config.min_total"           = "10"
      "uddi.asm_config.min_unused"          = "10"
      "uddi.asm_config.reenable_date"       = "2021-01-10T10:11:22Z"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = ""
    }
    check = {
      "uddi.comment" = ""
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.comment" = ""
    }
  }

}

case "ddns_client_update" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      ddns_client_update = "server"
    }
    check = {
      "uddi.ddns_client_update" = "server"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      ddns_client_update = "over_client_update"
    }
    check = {
      "uddi.ddns_client_update" = "over_client_update"
    }
  }

}

case "ddns_conflict_resolution_mode" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                          = "{{random}}"
      ddns_use_conflict_resolution  = false
      ddns_conflict_resolution_mode = "check_exists_with_dhcid"
    }
    check = {
      "uddi.ddns_use_conflict_resolution"  = "false"
      "uddi.ddns_conflict_resolution_mode" = "check_exists_with_dhcid"
    }
  }

  step {
    uddi {
      name                          = "{{random}}"
      ddns_use_conflict_resolution  = true
      ddns_conflict_resolution_mode = "check_with_dhcid"
    }
    check = {
      "uddi.ddns_use_conflict_resolution"  = "true"
      "uddi.ddns_conflict_resolution_mode" = "check_with_dhcid"
    }
  }

}

case "ddns_domain" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      ddns_domain = "abc"
    }
    check = {
      "uddi.ddns_domain" = "abc"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      ddns_domain = "xyz"
    }
    check = {
      "uddi.ddns_domain" = "xyz"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      ddns_domain = ""
    }
    check = {
      "uddi.ddns_domain" = ""
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.ddns_domain" = ""
    }
  }

}

case "ddns_generate_name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      ddns_generate_name = true
    }
    check = {
      "uddi.ddns_generate_name" = "true"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      ddns_generate_name = false
    }
    check = {
      "uddi.ddns_generate_name" = "false"
    }
  }

}

case "ddns_generated_prefix" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                  = "{{random}}"
      ddns_generated_prefix = "host-prefix"
    }
    check = {
      "uddi.ddns_generated_prefix" = "host-prefix"
    }
  }

  step {
    uddi {
      name                  = "{{random}}"
      ddns_generated_prefix = "host-another-prefix"
    }
    check = {
      "uddi.ddns_generated_prefix" = "host-another-prefix"
    }
  }

}

case "dhcp_options" {
  backend     = "uddi"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_dhcp_option_code_unknown" "test" {
  #   uddi = {
  #     code = 234
  #     name = "test_dhcp_option_code"
  #     option_space = infoblox_dhcp_option_space_unknown.test.id
  #     type = "boolean"
  #   }
  # }
  # resource "infoblox_dhcp_option_group_unknown" "test" {
  #   uddi = {
  #     name = "\"og-\"+name"
  #     protocol = "ip4"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      name         = "{{random}}"
      dhcp_options = [{ type = "option", option_code = "dhcp/option_code/4b9ddf44-e2e3-4f1c-a3aa-4e508933feed", option_value = "456456" }]
    }
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.type"         = "option"
      "uddi.dhcp_options.0.option_value" = "456456"
    }
  }

  step {
    uddi {
      name         = "{{random}}"
    }
    check = {
      "uddi.dhcp_options.#" = "0"
    }
  }

}

case "dhcp_options_v6" {
  backend     = "uddi"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_dhcp_option_code_unknown" "test" {
  #   uddi = {
  #     code = 234
  #     name = "test_dhcp_option_code"
  #     option_space = infoblox_dhcp_option_space_unknown.test.id
  #     type = "boolean"
  #   }
  # }
  # resource "infoblox_dhcp_option_group_unknown" "test" {
  #   uddi = {
  #     name = "\"og-\"+name"
  #     protocol = "ip6"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      name            = "{{random}}"
      dhcp_options_v6 = [{ type = "option", option_code = "dhcp/option_code/46bf2e4a-25c8-4ac3-b9f7-4244c8265f2e", option_value = "255" }]
    }
    check = {
      "uddi.dhcp_options_v6.#"              = "1"
      "uddi.dhcp_options_v6.0.type"         = "option"
      "uddi.dhcp_options_v6.0.option_value" = "255"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
    }
  }

}

case "ddns_send_updates" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name              = "{{random}}"
      ddns_send_updates = true
    }
    check = {
      "uddi.ddns_send_updates" = "true"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      ddns_send_updates = false
    }
    check = {
      "uddi.ddns_send_updates" = "false"
    }
  }

}

case "ddns_ttl_percent" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name             = "{{random}}"
      ddns_ttl_percent = 20
    }
    check = {
      "uddi.ddns_ttl_percent" = "20"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      ddns_ttl_percent = 40
    }
    check = {
      "uddi.ddns_ttl_percent" = "40"
    }
  }

}

case "ddns_update_on_renew" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                 = "{{random}}"
      ddns_update_on_renew = true
    }
    check = {
      "uddi.ddns_update_on_renew" = "true"
    }
  }

  step {
    uddi {
      name                 = "{{random}}"
      ddns_update_on_renew = false
    }
    check = {
      "uddi.ddns_update_on_renew" = "false"
    }
  }

}

case "dhcp_config" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      dhcp_config = { abandoned_reclaim_time = 1000, abandoned_reclaim_time_v6 = 2000, allow_unknown = true, allow_unknown_v6 = true, ignore_client_uid = true, lease_time = 50, lease_time_v6 = 60 }
    }
    check = {
      "uddi.dhcp_config.abandoned_reclaim_time"    = "1000"
      "uddi.dhcp_config.abandoned_reclaim_time_v6" = "2000"
      "uddi.dhcp_config.allow_unknown"             = "true"
      "uddi.dhcp_config.allow_unknown_v6"          = "true"
      "uddi.dhcp_config.ignore_client_uid"         = "true"
      "uddi.dhcp_config.lease_time"                = "50"
      "uddi.dhcp_config.lease_time_v6"             = "60"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      dhcp_config = { abandoned_reclaim_time = 1500, abandoned_reclaim_time_v6 = 2500, allow_unknown = false, allow_unknown_v6 = false, ignore_client_uid = false, lease_time = 55, lease_time_v6 = 65 }
    }
    check = {
      "uddi.dhcp_config.abandoned_reclaim_time"    = "1500"
      "uddi.dhcp_config.abandoned_reclaim_time_v6" = "2500"
      "uddi.dhcp_config.allow_unknown"             = "false"
      "uddi.dhcp_config.allow_unknown_v6"          = "false"
      "uddi.dhcp_config.ignore_client_uid"         = "false"
      "uddi.dhcp_config.lease_time"                = "55"
      "uddi.dhcp_config.lease_time_v6"             = "65"
    }
  }

}

case "default_realms" {
  backend     = "uddi"
  # skip        = true
  # skip_reason = "requires_resource: infoblox_federated_realm not yet implemented"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_federated_realm_unknown" "realm1" {
  #   uddi = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      name           = "{{random}}"
      default_realms = ["federation/federated_realm/f76ecc5e-db40-455c-a5f3-2b6ea57785cd"]
    }
    check = {
      "uddi.default_realms.#" = "1"
    }
  }

}

case "header_option_filename" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = "HEADER_OPTION_FILEip_space_name"
    }
    check = {
      "uddi.header_option_filename" = "HEADER_OPTION_FILEip_space_name"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"
    }
    check = {
      "uddi.header_option_filename" = "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      header_option_filename = ""
    }
    check = {
      "uddi.header_option_filename" = ""
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.header_option_filename" = ""
    }
  }

}

case "header_option_server_address" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = "10.0.0.0"
    }
    check = {
      "uddi.header_option_server_address" = "10.0.0.0"
    }
  }

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = "12.0.0.0"
    }
    check = {
      "uddi.header_option_server_address" = "12.0.0.0"
    }
  }

  step {
    uddi {
      name                         = "{{random}}"
      header_option_server_address = ""
    }
    check = {
      "uddi.header_option_server_address" = ""
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.header_option_server_address" = ""
    }
  }

}

case "header_option_server_name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = "HEADER_OPTION_SERVER_ip_space_name"
    }
    check = {
      "uddi.header_option_server_name" = "HEADER_OPTION_SERVER_ip_space_name"
    }
  }

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"
    }
    check = {
      "uddi.header_option_server_name" = "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"
    }
  }

  step {
    uddi {
      name                      = "{{random}}"
      header_option_server_name = ""
    }
    check = {
      "uddi.header_option_server_name" = ""
    }
  }

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.header_option_server_name" = ""
    }
  }

}

case "hostname_rewrite_char" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                  = "{{random}}"
      hostname_rewrite_char = "+"
    }
    check = {
      "uddi.hostname_rewrite_char" = "+"
    }
  }

  step {
    uddi {
      name                  = "{{random}}"
      hostname_rewrite_char = "/"
    }
    check = {
      "uddi.hostname_rewrite_char" = "/"
    }
  }

}

case "hostname_rewrite_enabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                     = "{{random}}"
      hostname_rewrite_enabled = true
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "true"
    }
  }

  step {
    uddi {
      name                     = "{{random}}"
      hostname_rewrite_enabled = false
    }
    check = {
      "uddi.hostname_rewrite_enabled" = "false"
    }
  }

}

case "hostname_rewrite_regex" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                   = "{{random}}"
      hostname_rewrite_regex = "[^a-z]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^a-z]"
    }
  }

  step {
    uddi {
      name                   = "{{random}}"
      hostname_rewrite_regex = "[^0-9]"
    }
    check = {
      "uddi.hostname_rewrite_regex" = "[^0-9]"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      inheritance_sources = { asm_config = { action = "inherit", asm_enable_block = { action = "inherit" }, asm_growth_block = { action = "inherit" }, asm_threshold = { action = "inherit" }, forecast_period = { action = "inherit" }, history = { action = "inherit" }, min_total = { action = "inherit" }, min_unused = { action = "inherit" } }, ddns_client_update = { action = "inherit" }, ddns_conflict_resolution_mode = { action = "inherit" }, ddns_enabled = { action = "inherit" }, ddns_hostname_block = { action = "inherit" }, ddns_ttl_percent = { action = "inherit" }, ddns_update_on_renew = { action = "inherit" }, ddns_use_conflict_resolution = { action = "inherit" }, header_option_filename = { action = "inherit" }, header_option_server_address = { action = "inherit" }, header_option_server_name = { action = "inherit" }, hostname_rewrite_block = { action = "inherit" }, vendor_specific_option_option_space = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.asm_config.asm_enable_block.action"         = "inherit"
      "uddi.inheritance_sources.asm_config.asm_growth_block.action"         = "inherit"
      "uddi.inheritance_sources.asm_config.asm_threshold.action"            = "inherit"
      "uddi.inheritance_sources.asm_config.forecast_period.action"          = "inherit"
      "uddi.inheritance_sources.asm_config.history.action"                  = "inherit"
      "uddi.inheritance_sources.asm_config.min_total.action"                = "inherit"
      "uddi.inheritance_sources.asm_config.min_unused.action"               = "inherit"
      "uddi.inheritance_sources.ddns_client_update.action"                  = "inherit"
      "uddi.inheritance_sources.ddns_conflict_resolution_mode.action"       = "inherit"
      "uddi.inheritance_sources.ddns_ttl_percent.action"                    = "inherit"
      "uddi.inheritance_sources.ddns_update_on_renew.action"                = "inherit"
      "uddi.inheritance_sources.ddns_use_conflict_resolution.action"        = "inherit"
      "uddi.inheritance_sources.header_option_filename.action"              = "inherit"
      "uddi.inheritance_sources.header_option_server_address.action"        = "inherit"
      "uddi.inheritance_sources.header_option_server_name.action"           = "inherit"
      "uddi.inheritance_sources.hostname_rewrite_block.action"              = "inherit"
      "uddi.inheritance_sources.vendor_specific_option_option_space.action" = "inherit"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      inheritance_sources = { asm_config = { action = "override", asm_enable_block = { action = "override" }, asm_growth_block = { action = "override" }, asm_threshold = { action = "override" }, forecast_period = { action = "override" }, history = { action = "override" }, min_total = { action = "override" }, min_unused = { action = "override" } }, ddns_client_update = { action = "override" }, ddns_conflict_resolution_mode = { action = "override" }, ddns_enabled = { action = "inherit" }, ddns_hostname_block = { action = "override" }, ddns_ttl_percent = { action = "override" }, ddns_update_on_renew = { action = "override" }, ddns_use_conflict_resolution = { action = "override" }, header_option_filename = { action = "override" }, header_option_server_address = { action = "override" }, header_option_server_name = { action = "override" }, hostname_rewrite_block = { action = "override" }, vendor_specific_option_option_space = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.asm_config.asm_enable_block.action"         = "override"
      "uddi.inheritance_sources.asm_config.asm_growth_block.action"         = "override"
      "uddi.inheritance_sources.asm_config.asm_threshold.action"            = "override"
      "uddi.inheritance_sources.asm_config.forecast_period.action"          = "override"
      "uddi.inheritance_sources.asm_config.history.action"                  = "override"
      "uddi.inheritance_sources.asm_config.min_total.action"                = "override"
      "uddi.inheritance_sources.asm_config.min_unused.action"               = "override"
      "uddi.inheritance_sources.ddns_client_update.action"                  = "override"
      "uddi.inheritance_sources.ddns_conflict_resolution_mode.action"       = "override"
      "uddi.inheritance_sources.ddns_ttl_percent.action"                    = "override"
      "uddi.inheritance_sources.ddns_update_on_renew.action"                = "override"
      "uddi.inheritance_sources.ddns_use_conflict_resolution.action"        = "override"
      "uddi.inheritance_sources.header_option_filename.action"              = "override"
      "uddi.inheritance_sources.header_option_server_address.action"        = "override"
      "uddi.inheritance_sources.header_option_server_name.action"           = "override"
      "uddi.inheritance_sources.hostname_rewrite_block.action"              = "override"
      "uddi.inheritance_sources.vendor_specific_option_option_space.action" = "override"
    }
  }

}

case "multiple_default_realms" {
  backend     = "uddi"
  skip        = true
  skip_reason = "t.Skip: Skipping test temporarily due to Multiple realms not being supported in the current test environment."
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm_unknown" "realm1" {
    uddi = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    uddi {
      name           = "{{random}}"
      default_realms = [infoblox_federated_realm_unknown.realm1.id, infoblox_federated_realm_unknown.realm2.id, infoblox_federated_realm_unknown.realm3.id, infoblox_federated_realm_unknown.realm4.id, infoblox_federated_realm_unknown.realm5.id]
    }
    check = {
      "uddi.default_realms.#" = "5"
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      default_realms = [infoblox_federated_realm_unknown.realm1_updated.id, infoblox_federated_realm_unknown.realm2_updated.id, infoblox_federated_realm_unknown.realm3_updated.id, infoblox_federated_realm_unknown.realm4_updated.id, infoblox_federated_realm_unknown.realm5_updated.id]
    }
    check = {
      "uddi.default_realms.#" = "5"
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
