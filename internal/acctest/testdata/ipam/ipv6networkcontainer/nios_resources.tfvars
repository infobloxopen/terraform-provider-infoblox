# Auto-generated resource acceptance-test cases for Ipv6networkcontainer.
// Objects to be present on the grid for testing
// ipv6_option_filter , ipv6_option_filter1- IPv6 Option Filters
// rir-org-test1 - RIR Organization
// ISE Server has to be configured
// A discovery member has to be configured

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_ipv6_network}}"
    }
    check = {
      "nios.network"                              = "{{random_ipv6_network}}"
      "nios.ddns_generate_hostname"               = "false"
      "nios.ddns_server_always_updates"           = "true"
      "nios.ddns_ttl"                             = "0"
      "nios.enable_ddns"                          = "false"
      "nios.enable_discovery"                     = "false"
      "nios.mgm_private"                          = "false"
      "nios.network_view"                         = "default"
      "nios.rir_registration_status"              = "NOT_REGISTERED"
      "nios.same_port_control_discovery_blackout" = "false"
      "nios.unmanaged"                            = "false"
      "nios.update_dns_on_lease_renewal"          = "false"
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
      network = "{{random_ipv6_network}}"
    }
  }

}

case "auto_create_reversezone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                 = "{{random_ipv6_network}}"
      auto_create_reversezone = true
    }
    check = {
      "nios.auto_create_reversezone" = "true"
    }
  }

}

case "cloud_info" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_ipv6_network}}"
    }
    check = {
      "nios.network" = "{{random_ipv6_network}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_ipv6_network}}"
      comment = "test comment"
    }
    check = {
      "nios.comment" = "test comment"
      "nios.network" = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network = "{{random_ipv6_network}}"
      comment = "updated comment"
    }
    check = {
      "nios.comment" = "updated comment"
      "nios.network" = "{{random_ipv6_network}}"
    }
  }

}

case "delete_reason" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network       = "{{random_ipv6_network}}"
      delete_reason = "test-delete-reason"
    }
    check = {
      "nios.delete_reason" = "test-delete-reason"
    }
  }

  step {
    nios {
      network       = "{{random_ipv6_network}}"
      delete_reason = "updated-delete-reason"
    }
    check = {
      "nios.delete_reason" = "updated-delete-reason"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network         = "{{random_ipv6_network}}"
      ddns_domainname = "test.com"
    }
    check = {
      "nios.ddns_domainname" = "test.com"
      "nios.network"         = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network         = "{{random_ipv6_network}}"
      ddns_domainname = "testupdated.com"
    }
    check = {
      "nios.ddns_domainname" = "testupdated.com"
      "nios.network"         = "{{random_ipv6_network}}"
    }
  }

}

case "ddns_enable_option_fqdn" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                 = "{{random_ipv6_network}}"
      ddns_enable_option_fqdn = false
    }
    check = {
      "nios.ddns_enable_option_fqdn" = "false"
      "nios.network"                 = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network                 = "{{random_ipv6_network}}"
      ddns_enable_option_fqdn = true
    }
    check = {
      "nios.ddns_enable_option_fqdn" = "true"
      "nios.network"                 = "{{random_ipv6_network}}"
    }
  }

}

case "ddns_generate_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                = "{{random_ipv6_network}}"
      ddns_generate_hostname = false
    }
    check = {
      "nios.ddns_generate_hostname" = "false"
      "nios.network"                = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network                = "{{random_ipv6_network}}"
      ddns_generate_hostname = true
    }
    check = {
      "nios.ddns_generate_hostname" = "true"
      "nios.network"                = "{{random_ipv6_network}}"
    }
  }

}

case "ddns_server_always_updates" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                    = "{{random_ipv6_network}}"
      ddns_server_always_updates = true
      ddns_enable_option_fqdn    = true
    }
    check = {
      "nios.ddns_server_always_updates" = "true"
      "nios.ddns_enable_option_fqdn"    = "true"
      "nios.network"                    = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network                    = "{{random_ipv6_network}}"
      ddns_server_always_updates = false
      ddns_enable_option_fqdn    = true
    }
    check = {
      "nios.ddns_server_always_updates" = "false"
      "nios.ddns_enable_option_fqdn"    = "true"
      "nios.network"                    = "{{random_ipv6_network}}"
    }
  }

}

case "ddns_ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network  = "{{random_ipv6_network}}"
      ddns_ttl = 1
    }
    check = {
      "nios.ddns_ttl" = "1"
      "nios.network"  = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network  = "{{random_ipv6_network}}"
      ddns_ttl = 2
    }
    check = {
      "nios.ddns_ttl" = "2"
      "nios.network"  = "{{random_ipv6_network}}"
    }
  }

}

case "discovery_basic_poll_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                       = "{{random_ipv6_network}}"
      discovery_basic_poll_settings = { auto_arp_refresh_before_switch_port_polling = true, cli_collection = false, complete_ping_sweep = false, device_profile = false, switch_port_data_collection_polling = "PERIODIC" }
    }
    check = {
      "nios.discovery_basic_poll_settings.cli_collection"                              = "false"
      "nios.discovery_basic_poll_settings.switch_port_data_collection_polling"         = "PERIODIC"
      "nios.discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling" = "true"
      "nios.discovery_basic_poll_settings.complete_ping_sweep"                         = "false"
      "nios.discovery_basic_poll_settings.device_profile"                              = "false"
    }
  }

  step {
    nios {
      network                       = "{{random_ipv6_network}}"
      discovery_basic_poll_settings = { auto_arp_refresh_before_switch_port_polling = true, cli_collection = true, complete_ping_sweep = false, device_profile = false, switch_port_data_collection_polling = "SCHEDULED" }
    }
    check = {
      "nios.discovery_basic_poll_settings.cli_collection"                              = "true"
      "nios.discovery_basic_poll_settings.switch_port_data_collection_polling"         = "SCHEDULED"
      "nios.discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling" = "true"
      "nios.discovery_basic_poll_settings.complete_ping_sweep"                         = "false"
      "nios.discovery_basic_poll_settings.device_profile"                              = "false"
    }
  }

  step {
    nios {
      network                       = "{{random_ipv6_network}}"
      discovery_basic_poll_settings = { auto_arp_refresh_before_switch_port_polling = true, cli_collection = true, complete_ping_sweep = false, device_profile = false, switch_port_data_collection_polling = "DISABLED" }
    }
    check = {
      "nios.discovery_basic_poll_settings.cli_collection"                              = "true"
      "nios.discovery_basic_poll_settings.switch_port_data_collection_polling"         = "DISABLED"
      "nios.discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling" = "true"
      "nios.discovery_basic_poll_settings.complete_ping_sweep"                         = "false"
      "nios.discovery_basic_poll_settings.device_profile"                              = "false"
    }
  }

}

case "discovery_blackout_setting" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                    = "{{random_ipv6_network}}"
      discovery_blackout_setting = { enable_blackout = true, blackout_duration = 100, blackout_schedule = { weekdays = ["TUESDAY", "MONDAY", "FRIDAY"], frequency = "WEEKLY", every = 15, minutes_past_hour = 6, disable = false, repeat = "RECUR", hour_of_day = 20 } }
    }
    check = {
      "nios.discovery_blackout_setting.enable_blackout"                     = "true"
      "nios.discovery_blackout_setting.blackout_duration"                   = "100"
      "nios.discovery_blackout_setting.blackout_schedule.weekdays.0"        = "TUESDAY"
      "nios.discovery_blackout_setting.blackout_schedule.weekdays.1"        = "MONDAY"
      "nios.discovery_blackout_setting.blackout_schedule.weekdays.2"        = "FRIDAY"
      "nios.discovery_blackout_setting.blackout_schedule.frequency"         = "WEEKLY"
      "nios.discovery_blackout_setting.blackout_schedule.every"             = "15"
      "nios.discovery_blackout_setting.blackout_schedule.minutes_past_hour" = "6"
      "nios.discovery_blackout_setting.blackout_schedule.disable"           = "false"
      "nios.discovery_blackout_setting.blackout_schedule.repeat"            = "RECUR"
    }
  }

  step {
    nios {
      network                    = "{{random_ipv6_network}}"
      discovery_blackout_setting = { enable_blackout = true, blackout_duration = 200, blackout_schedule = { minutes_past_hour = 6, repeat = "ONCE", day_of_month = 30, month = 1, year = 2026, hour_of_day = 20 } }
    }
    check = {
      "nios.discovery_blackout_setting.enable_blackout"                     = "true"
      "nios.discovery_blackout_setting.blackout_duration"                   = "200"
      "nios.discovery_blackout_setting.blackout_schedule.minutes_past_hour" = "6"
      "nios.discovery_blackout_setting.blackout_schedule.repeat"            = "ONCE"
      "nios.discovery_blackout_setting.blackout_schedule.day_of_month"      = "30"
      "nios.discovery_blackout_setting.blackout_schedule.month"             = "1"
      "nios.discovery_blackout_setting.blackout_schedule.year"              = "2026"
    }
  }

}

case "domain_name_servers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network             = "{{random_ipv6_network}}"
      domain_name_servers = ["100::1"]
    }
    check = {
      "nios.domain_name_servers.0" = "100::1"
      "nios.network"               = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network             = "{{random_ipv6_network}}"
      domain_name_servers = ["100::2"]
    }
    check = {
      "nios.domain_name_servers.0" = "100::2"
      "nios.network"               = "{{random_ipv6_network}}"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network     = "{{random_ipv6_network}}"
      enable_ddns = false
    }
    check = {
      "nios.enable_ddns" = "false"
      "nios.network"     = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network     = "{{random_ipv6_network}}"
      enable_ddns = true
    }
    check = {
      "nios.enable_ddns" = "true"
      "nios.network"     = "{{random_ipv6_network}}"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network   = "{{random_ipv6_network}}"
      ext_attrs = { Site = "{{random}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random}}"
      "nios.network"        = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network   = "{{random_ipv6_network}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
      "nios.network"        = "{{random_ipv6_network}}"
    }
  }

}

case "mgm_private" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network     = "{{random_ipv6_network}}"
      mgm_private = false
    }
    check = {
      "nios.mgm_private" = "false"
      "nios.network"     = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network     = "{{random_ipv6_network}}"
      mgm_private = true
    }
    check = {
      "nios.mgm_private" = "true"
      "nios.network"     = "{{random_ipv6_network}}"
    }
  }

}

case "network" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_ipv6_network}}"
    }
    check = {
      "nios.network" = "{{random_ipv6_network}}"
    }
  }

}

case "network_view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network      = "{{random_ipv6_network}}"
      network_view = "default"
    }
    check = {
      "nios.network_view" = "default"
      "nios.network"      = "{{random_ipv6_network}}"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_ipv6_network}}"
      options = [{ name = "dhcp6.fqdn", num = 39, value = "test.com", vendor_class = "DHCPv6" }]
    }
    check = {
      "nios.network"                = "{{random_ipv6_network}}"
      "nios.options.#"              = "1"
      "nios.options.0.name"         = "dhcp6.fqdn"
      "nios.options.0.num"          = "39"
      "nios.options.0.value"        = "test.com"
      "nios.options.0.vendor_class" = "DHCPv6"
    }
  }

  step {
    nios {
      network = "{{random_ipv6_network}}"
      options = [{ name = "dhcp-rebinding-time", num = 59, value = 100, vendor_class = "DHCP" }]
    }
    check = {
      "nios.network"                = "{{random_ipv6_network}}"
      "nios.options.#"              = "1"
      "nios.options.0.name"         = "dhcp-rebinding-time"
      "nios.options.0.num"          = "59"
      "nios.options.0.value"        = "100"
      "nios.options.0.vendor_class" = "DHCP"
    }
  }

}

case "port_control_blackout_setting" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                       = "{{random_ipv6_network}}"
      port_control_blackout_setting = { enable_blackout = true, blackout_duration = 100, blackout_schedule = { weekdays = ["TUESDAY", "MONDAY", "FRIDAY"], frequency = "WEEKLY", every = 15, minutes_past_hour = 6, disable = false, repeat = "RECUR", hour_of_day = 20 } }
    }
    check = {
      "nios.port_control_blackout_setting.enable_blackout"                     = "true"
      "nios.port_control_blackout_setting.blackout_duration"                   = "100"
      "nios.port_control_blackout_setting.blackout_schedule.weekdays.0"        = "TUESDAY"
      "nios.port_control_blackout_setting.blackout_schedule.weekdays.1"        = "MONDAY"
      "nios.port_control_blackout_setting.blackout_schedule.weekdays.2"        = "FRIDAY"
      "nios.port_control_blackout_setting.blackout_schedule.frequency"         = "WEEKLY"
      "nios.port_control_blackout_setting.blackout_schedule.every"             = "15"
      "nios.port_control_blackout_setting.blackout_schedule.minutes_past_hour" = "6"
      "nios.port_control_blackout_setting.blackout_schedule.disable"           = "false"
      "nios.port_control_blackout_setting.blackout_schedule.repeat"            = "RECUR"
    }
  }

  step {
    nios {
      network                       = "{{random_ipv6_network}}"
      port_control_blackout_setting = { enable_blackout = true, blackout_duration = 200, blackout_schedule = { minutes_past_hour = 6, repeat = "ONCE", day_of_month = 30, month = 1, year = 2026, hour_of_day = 20 } }
    }
    check = {
      "nios.port_control_blackout_setting.enable_blackout"                     = "true"
      "nios.port_control_blackout_setting.blackout_duration"                   = "200"
      "nios.port_control_blackout_setting.blackout_schedule.minutes_past_hour" = "6"
      "nios.port_control_blackout_setting.blackout_schedule.repeat"            = "ONCE"
      "nios.port_control_blackout_setting.blackout_schedule.day_of_month"      = "30"
      "nios.port_control_blackout_setting.blackout_schedule.month"             = "1"
      "nios.port_control_blackout_setting.blackout_schedule.year"              = "2026"
    }
  }

}

case "preferred_lifetime" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      preferred_lifetime = 27000
      valid_lifetime     = 43200
    }
    check = {
      "nios.preferred_lifetime" = "27000"
      "nios.network"            = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      preferred_lifetime = 30000
      valid_lifetime     = 43200
    }
    check = {
      "nios.preferred_lifetime" = "30000"
      "nios.network"            = "{{random_ipv6_network}}"
    }
  }

}

case "restart_if_needed" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network           = "{{random_ipv6_network}}"
      restart_if_needed = false
    }
    check = {
      "nios.restart_if_needed" = "false"
    }
  }

  step {
    nios {
      network           = "{{random_ipv6_network}}"
      restart_if_needed = true
    }
    check = {
      "nios.restart_if_needed" = "true"
    }
  }

}

case "rir_registration_action" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6networkcontainer" "rir_parent" {
    nios = {
      network          = "2001:db8:{{random_hextet}}::/48"
      rir_organization = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name" = "TEST-NET-V6"
        "RIPE Description" = "Test IPv6 network"
        "RIPE Country" = "United States (US)"
        "RIPE Admin Contact" = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source" = "RIPE"
        "RIPE IPv6 Status" = "ASSIGNED"
      }
    }
  }
  PREREQ

  step {
    nios {
      network                 = "2001:db8:{{random_hextet}}:1::/64"
      rir_registration_action = "CREATE"
      rir_organization        = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-V6-CHILD"
        "RIPE Description"       = "Test IPv6 child network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv6 Status"       = "ASSIGNED"
      }
    }
    depends_on = [infoblox_ipv6networkcontainer.rir_parent]
    check = {
      "nios.rir_registration_action" = "CREATE"
      "nios.network"                 = "2001:db8:{{random_hextet}}:1::/64"
      "nios.rir_organization"        = "rir-org-test1"
    }
  }

  step {
    nios {
      network                 = "2001:db8:{{random_hextet}}:1::/64"
      rir_registration_action = "NONE"
      rir_organization        = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-V6-CHILD"
        "RIPE Description"       = "Test IPv6 child network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv6 Status"       = "ASSIGNED"
      }
    }
    depends_on = [infoblox_ipv6networkcontainer.rir_parent]
    check = {
      "nios.rir_registration_action" = "NONE"
      "nios.network"                 = "2001:db8:{{random_hextet}}:1::/64"
      "nios.rir_organization"        = "rir-org-test1"
    }
  }
}

case "rir_registration_status" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                              = "{{random_ipv6_network}}"
      rir_registration_status              = "NOT_REGISTERED"
      same_port_control_discovery_blackout = false
    }
    check = {
      "nios.rir_registration_status"              = "NOT_REGISTERED"
      "nios.same_port_control_discovery_blackout" = "false"
      "nios.network"                              = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network                              = "{{random_ipv6_network}}"
      rir_registration_status              = "NOT_REGISTERED"
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.rir_registration_status"              = "NOT_REGISTERED"
      "nios.same_port_control_discovery_blackout" = "true"
      "nios.network"                              = "{{random_ipv6_network}}"
    }
  }

}

case "send_rir_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_ipv6_network}}"
      send_rir_request = true
    }
    check = {
      "nios.send_rir_request" = "true"
      "nios.network"          = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network          = "{{random_ipv6_network}}"
      send_rir_request = false
    }
    check = {
      "nios.send_rir_request" = "false"
      "nios.network"          = "{{random_ipv6_network}}"
    }
  }

}

case "same_port_control_discovery_blackout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                              = "{{random_ipv6_network}}"
      same_port_control_discovery_blackout = false
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "false"
      "nios.network"                              = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network                              = "{{random_ipv6_network}}"
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "true"
      "nios.network"                              = "{{random_ipv6_network}}"
    }
  }

}

case "unmanaged" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network   = "{{random_ipv6_network}}"
      unmanaged = false
    }
    check = {
      "nios.unmanaged" = "false"
    }
  }

}

case "update_dns_on_lease_renewal" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                     = "{{random_ipv6_network}}"
      update_dns_on_lease_renewal = false
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "false"
      "nios.network"                     = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network                     = "{{random_ipv6_network}}"
      update_dns_on_lease_renewal = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
      "nios.network"                     = "{{random_ipv6_network}}"
    }
  }

}

case "valid_lifetime" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network        = "{{random_ipv6_network}}"
      valid_lifetime = 43200
    }
    check = {
      "nios.valid_lifetime" = "43200"
      "nios.network"        = "{{random_ipv6_network}}"
    }
  }

  step {
    nios {
      network        = "{{random_ipv6_network}}"
      valid_lifetime = 50000
    }
    check = {
      "nios.valid_lifetime" = "50000"
      "nios.network"        = "{{random_ipv6_network}}"
    }
  }

}

case "discovery_member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_ipv6_network}}"
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.discovery_member" = "{{discovery_member_hostname}}"
    }
  }

}

case "enable_discovery" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_ipv6_network}}"
      discovery_member = "{{discovery_member_hostname}}"
      enable_discovery = true
    }
    check = {
      "nios.enable_discovery" = "true"
    }
  }

  step {
    nios {
      network          = "{{random_ipv6_network}}"
      discovery_member = "{{discovery_member_hostname}}"
      enable_discovery = false
    }
    check = {
      "nios.enable_discovery" = "false"
    }
  }

}

case "enable_immediate_discovery" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                    = "{{random_ipv6_network}}"
      discovery_member           = "{{discovery_member_hostname}}"
      enable_immediate_discovery = true
    }
    check = {
      "nios.enable_immediate_discovery" = "true"
    }
  }

  step {
    nios {
      network                    = "{{random_ipv6_network}}"
      discovery_member           = "{{discovery_member_hostname}}"
      enable_immediate_discovery = false
    }
    check = {
      "nios.enable_immediate_discovery" = "false"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      logic_filter_rules = [{ filter = "ipv6_option_filter", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      logic_filter_rules = [{ filter = "ipv6_option_filter1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "ipv6_option_filter1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "rir_organization" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_ipv6_network}}"
      rir_organization = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-V6"
        "RIPE Description"       = "Test IPv6 network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv6 Status"       = "ASSIGNED"
      }
    }
    check = {
      "nios.rir_organization" = "rir-org-test1"
    }
  }
}

case "subscribe_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["DOMAINNAME", "ENDPOINT_PROFILE"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "DOMAINNAME"
      "nios.subscribe_settings.enabled_attributes.1" = "ENDPOINT_PROFILE"
    }
  }

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["SECURITY_GROUP"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "SECURITY_GROUP"
    }
  }

}

case "mapped_ea_attributes" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["USERNAME"], mapped_ea_attributes = [{ name = "IP_ADDRESS", mapped_ea = "Site" }] }
    }
    check = {
      "nios.subscribe_settings.mapped_ea_attributes.0.name"      = "IP_ADDRESS"
      "nios.subscribe_settings.mapped_ea_attributes.0.mapped_ea" = "Site"
    }
  }

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["USERNAME"], mapped_ea_attributes = [{ name = "MAC", mapped_ea = "Site" }] }
    }
    check = {
      "nios.subscribe_settings.mapped_ea_attributes.0.name"      = "MAC"
      "nios.subscribe_settings.mapped_ea_attributes.0.mapped_ea" = "Site"
    }
  }

  step {
    nios {
      network            = "{{random_ipv6_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["USERNAME"], mapped_ea_attributes = [{ name = "MAC", mapped_ea = "Building" }] }
    }
    check = {
      "nios.subscribe_settings.mapped_ea_attributes.0.name"      = "MAC"
      "nios.subscribe_settings.mapped_ea_attributes.0.mapped_ea" = "Building"
    }
  }

}
