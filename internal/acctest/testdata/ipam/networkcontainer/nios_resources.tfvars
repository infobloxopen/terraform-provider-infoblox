# Auto-generated resource acceptance-test cases for Networkcontainer.
// Objects to be present on the grid for testing
// mac_filter - Filter mac 
// example-option-filter-1 - Filter Option
// rir-org-test1 - RIR Organization
// ISE Server has to be configured
// A discovery member has to be configured

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
    check = {
      "nios.network"                              = "{{random_cidr_network}}"
      "nios.authority"                            = "false"
      "nios.auto_create_reversezone"              = "false"
      "nios.ddns_generate_hostname"               = "false"
      "nios.ddns_server_always_updates"           = "true"
      "nios.ddns_ttl"                             = "0"
      "nios.ddns_update_fixed_addresses"          = "false"
      "nios.ddns_use_option81"                    = "false"
      "nios.deny_bootp"                           = "false"
      "nios.enable_ddns"                          = "false"
      "nios.enable_dhcp_thresholds"               = "false"
      "nios.enable_email_warnings"                = "false"
      "nios.enable_pxe_lease_time"                = "false"
      "nios.enable_snmp_warnings"                 = "false"
      "nios.high_water_mark"                      = "95"
      "nios.high_water_mark_reset"                = "85"
      "nios.ignore_dhcp_option_list_request"      = "false"
      "nios.ignore_id"                            = "NONE"
      "nios.lease_scavenge_time"                  = "-1"
      "nios.low_water_mark"                       = "0"
      "nios.low_water_mark_reset"                 = "10"
      "nios.mgm_private"                          = "false"
      "nios.network_view"                         = "default"
      "nios.recycle_leases"                       = "true"
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
      network = "{{random_cidr_network}}"
    }
  }

}

case "authority" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network   = "{{random_cidr_network}}"
      authority = false
    }
    check = {
      "nios.authority" = "false"
      "nios.network"   = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network   = "{{random_cidr_network}}"
      authority = true
    }
    check = {
      "nios.authority" = "true"
      "nios.network"   = "{{random_cidr_network}}"
    }
  }

}

case "bootfile" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network  = "{{random_cidr_network}}"
      bootfile = "bootfile"
    }
    check = {
      "nios.bootfile" = "bootfile"
      "nios.network"  = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network  = "{{random_cidr_network}}"
      bootfile = "bootfile_updated"
    }
    check = {
      "nios.bootfile" = "bootfile_updated"
      "nios.network"  = "{{random_cidr_network}}"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network    = "{{random_cidr_network}}"
      bootserver = "test_bootserver"
    }
    check = {
      "nios.bootserver" = "test_bootserver"
      "nios.network"    = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      bootserver = "test_bootserver_updated"
    }
    check = {
      "nios.bootserver" = "test_bootserver_updated"
      "nios.network"    = "{{random_cidr_network}}"
    }
  }

}

case "cloud_info" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
    check = {
      "nios.network" = "{{random_cidr_network}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
      comment = "test comment"
    }
    check = {
      "nios.comment" = "test comment"
      "nios.network" = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
      comment = "updated comment"
    }
    check = {
      "nios.comment" = "updated comment"
      "nios.network" = "{{random_cidr_network}}"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network         = "{{random_cidr_network}}"
      ddns_domainname = "test.com"
    }
    check = {
      "nios.ddns_domainname" = "test.com"
      "nios.network"         = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network         = "{{random_cidr_network}}"
      ddns_domainname = "testupdated.com"
    }
    check = {
      "nios.ddns_domainname" = "testupdated.com"
      "nios.network"         = "{{random_cidr_network}}"
    }
  }

}

case "ddns_generate_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                = "{{random_cidr_network}}"
      ddns_generate_hostname = false
    }
    check = {
      "nios.ddns_generate_hostname" = "false"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                = "{{random_cidr_network}}"
      ddns_generate_hostname = true
    }
    check = {
      "nios.ddns_generate_hostname" = "true"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

}

case "ddns_server_always_updates" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                    = "{{random_cidr_network}}"
      ddns_server_always_updates = true
      ddns_use_option81          = true
    }
    check = {
      "nios.ddns_server_always_updates" = "true"
      "nios.network"                    = "{{random_cidr_network}}"
      "nios.ddns_use_option81"          = "true"
    }
  }

  step {
    nios {
      network                    = "{{random_cidr_network}}"
      ddns_server_always_updates = false
      ddns_use_option81          = true
    }
    check = {
      "nios.ddns_server_always_updates" = "false"
      "nios.network"                    = "{{random_cidr_network}}"
      "nios.ddns_use_option81"          = "true"
    }
  }

}

case "ddns_ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network  = "{{random_cidr_network}}"
      ddns_ttl = 1
    }
    check = {
      "nios.ddns_ttl" = "1"
      "nios.network"  = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network  = "{{random_cidr_network}}"
      ddns_ttl = 2
    }
    check = {
      "nios.ddns_ttl" = "2"
      "nios.network"  = "{{random_cidr_network}}"
    }
  }

}

case "ddns_update_fixed_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                     = "{{random_cidr_network}}"
      ddns_update_fixed_addresses = true
    }
    check = {
      "nios.ddns_update_fixed_addresses" = "true"
      "nios.network"                     = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                     = "{{random_cidr_network}}"
      ddns_update_fixed_addresses = false
    }
    check = {
      "nios.ddns_update_fixed_addresses" = "false"
      "nios.network"                     = "{{random_cidr_network}}"
    }
  }

}

case "ddns_use_option81" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network           = "{{random_cidr_network}}"
      ddns_use_option81 = false
    }
    check = {
      "nios.ddns_use_option81" = "false"
      "nios.network"           = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network           = "{{random_cidr_network}}"
      ddns_use_option81 = true
    }
    check = {
      "nios.ddns_use_option81" = "true"
      "nios.network"           = "{{random_cidr_network}}"
    }
  }

}

case "delete_reason" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network       = "{{random_cidr_network}}"
      delete_reason = "test-delete-reason"
    }
    check = {
      "nios.delete_reason" = "test-delete-reason"
    }
  }

  step {
    nios {
      network       = "{{random_cidr_network}}"
      delete_reason = "updated-delete-reason"
    }
    check = {
      "nios.delete_reason" = "updated-delete-reason"
    }
  }

}

case "deny_bootp" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network    = "{{random_cidr_network}}"
      deny_bootp = false
    }
    check = {
      "nios.deny_bootp" = "false"
      "nios.network"    = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      deny_bootp = true
    }
    check = {
      "nios.deny_bootp" = "true"
      "nios.network"    = "{{random_cidr_network}}"
    }
  }

}

case "discovery_basic_poll_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                       = "{{random_cidr_network}}"
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
      network                       = "{{random_cidr_network}}"
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
      network                       = "{{random_cidr_network}}"
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
      network                    = "{{random_cidr_network}}"
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
      network                    = "{{random_cidr_network}}"
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

case "email_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network    = "{{random_cidr_network}}"
      email_list = ["test@infoblox.com"]
    }
    check = {
      "nios.email_list.0" = "test@infoblox.com"
      "nios.network"      = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      email_list = ["update@test.com"]
    }
    check = {
      "nios.email_list.0" = "update@test.com"
      "nios.network"      = "{{random_cidr_network}}"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network     = "{{random_cidr_network}}"
      enable_ddns = false
    }
    check = {
      "nios.enable_ddns" = "false"
      "nios.network"     = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network     = "{{random_cidr_network}}"
      enable_ddns = true
    }
    check = {
      "nios.enable_ddns" = "true"
      "nios.network"     = "{{random_cidr_network}}"
    }
  }

}

case "enable_dhcp_thresholds" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                = "{{random_cidr_network}}"
      enable_dhcp_thresholds = false
    }
    check = {
      "nios.enable_dhcp_thresholds" = "false"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                = "{{random_cidr_network}}"
      enable_dhcp_thresholds = true
    }
    check = {
      "nios.enable_dhcp_thresholds" = "true"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

}

case "enable_email_warnings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network               = "{{random_cidr_network}}"
      enable_email_warnings = false
    }
    check = {
      "nios.enable_email_warnings" = "false"
      "nios.network"               = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network               = "{{random_cidr_network}}"
      enable_email_warnings = true
    }
    check = {
      "nios.enable_email_warnings" = "true"
      "nios.network"               = "{{random_cidr_network}}"
    }
  }

}

case "enable_pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network               = "{{random_cidr_network}}"
      pxe_lease_time        = 100
      enable_pxe_lease_time = false
    }
    check = {
      "nios.enable_pxe_lease_time" = "false"
      "nios.pxe_lease_time"        = "100"
      "nios.network"               = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network               = "{{random_cidr_network}}"
      pxe_lease_time        = 100
      enable_pxe_lease_time = true
    }
    check = {
      "nios.enable_pxe_lease_time" = "true"
      "nios.pxe_lease_time"        = "100"
      "nios.network"               = "{{random_cidr_network}}"
    }
  }

}

case "enable_snmp_warnings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network              = "{{random_cidr_network}}"
      enable_snmp_warnings = false
    }
    check = {
      "nios.enable_snmp_warnings" = "false"
      "nios.network"              = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      enable_snmp_warnings = true
    }
    check = {
      "nios.enable_snmp_warnings" = "true"
      "nios.network"              = "{{random_cidr_network}}"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ext_attrs = { Site = "{{random}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random}}"
      "nios.network"        = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
      "nios.network"        = "{{random_cidr_network}}"
    }
  }

}

case "high_water_mark" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network         = "{{random_cidr_network}}"
      high_water_mark = 95
    }
    check = {
      "nios.high_water_mark" = "95"
      "nios.network"         = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network         = "{{random_cidr_network}}"
      high_water_mark = 90
    }
    check = {
      "nios.high_water_mark" = "90"
      "nios.network"         = "{{random_cidr_network}}"
    }
  }

}

case "high_water_mark_reset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network               = "{{random_cidr_network}}"
      high_water_mark_reset = 85
    }
    check = {
      "nios.high_water_mark_reset" = "85"
      "nios.network"               = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network               = "{{random_cidr_network}}"
      high_water_mark_reset = 80
    }
    check = {
      "nios.high_water_mark_reset" = "80"
      "nios.network"               = "{{random_cidr_network}}"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                         = "{{random_cidr_network}}"
      ignore_dhcp_option_list_request = false
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
      "nios.network"                         = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                         = "{{random_cidr_network}}"
      ignore_dhcp_option_list_request = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
      "nios.network"                         = "{{random_cidr_network}}"
    }
  }

}

case "ignore_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ignore_id = "NONE"
    }
    check = {
      "nios.ignore_id" = "NONE"
      "nios.network"   = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ignore_id = "MACADDR"
    }
    check = {
      "nios.ignore_id" = "MACADDR"
      "nios.network"   = "{{random_cidr_network}}"
    }
  }

}

case "ignore_mac_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network              = "{{random_cidr_network}}"
      ignore_mac_addresses = ["aa:bb:cc:dd:ee:ff"]
    }
    check = {
      "nios.ignore_mac_addresses.0" = "aa:bb:cc:dd:ee:ff"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      ignore_mac_addresses = ["ff:ee:dd:cc:bb:aa"]
    }
    check = {
      "nios.ignore_mac_addresses.0" = "ff:ee:dd:cc:bb:aa"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

}

case "ipam_email_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network              = "{{random_cidr_network}}"
      ipam_email_addresses = ["testuser@infoblox.com"]
    }
    check = {
      "nios.ipam_email_addresses.0" = "testuser@infoblox.com"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      ipam_email_addresses = ["testuserupdated@infoblox.com"]
    }
    check = {
      "nios.ipam_email_addresses.0" = "testuserupdated@infoblox.com"
      "nios.network"                = "{{random_cidr_network}}"
    }
  }

}

case "ipam_threshold_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      ipam_threshold_settings = { reset_value = 85, trigger_value = 95 }
    }
    check = {
      "nios.ipam_threshold_settings.reset_value"   = "85"
      "nios.ipam_threshold_settings.trigger_value" = "95"
      "nios.network"                               = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      ipam_threshold_settings = { reset_value = 75, trigger_value = 80 }
    }
    check = {
      "nios.ipam_threshold_settings.reset_value"   = "75"
      "nios.ipam_threshold_settings.trigger_value" = "80"
      "nios.network"                               = "{{random_cidr_network}}"
    }
  }

}

case "ipam_trap_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_cidr_network}}"
      ipam_trap_settings = { enable_email_warnings = false, enable_snmp_warnings = true }
    }
    check = {
      "nios.ipam_trap_settings.enable_email_warnings" = "false"
      "nios.ipam_trap_settings.enable_snmp_warnings"  = "true"
    }
  }

  step {
    nios {
      network            = "{{random_cidr_network}}"
      ipam_trap_settings = { enable_email_warnings = true, enable_snmp_warnings = false }
    }
    check = {
      "nios.ipam_trap_settings.enable_email_warnings" = "true"
      "nios.ipam_trap_settings.enable_snmp_warnings"  = "false"
    }
  }

}

case "lease_scavenge_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network             = "{{random_cidr_network}}"
      lease_scavenge_time = -1
    }
    check = {
      "nios.lease_scavenge_time" = "-1"
    }
  }

  step {
    nios {
      network             = "{{random_cidr_network}}"
      lease_scavenge_time = 86400
    }
    check = {
      "nios.lease_scavenge_time" = "86400"
    }
  }

}

case "low_water_mark" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network        = "{{random_cidr_network}}"
      low_water_mark = 0
    }
    check = {
      "nios.low_water_mark" = "0"
    }
  }

  step {
    nios {
      network        = "{{random_cidr_network}}"
      low_water_mark = 50
    }
    check = {
      "nios.low_water_mark" = "50"
    }
  }

}

case "low_water_mark_reset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network              = "{{random_cidr_network}}"
      low_water_mark_reset = 10
    }
    check = {
      "nios.low_water_mark_reset" = "10"
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      low_water_mark_reset = 20
    }
    check = {
      "nios.low_water_mark_reset" = "20"
    }
  }

}

case "mgm_private" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network     = "{{random_cidr_network}}"
      mgm_private = false
    }
    check = {
      "nios.mgm_private" = "false"
      "nios.network"     = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network     = "{{random_cidr_network}}"
      mgm_private = true
    }
    check = {
      "nios.mgm_private" = "true"
      "nios.network"     = "{{random_cidr_network}}"
    }
  }

}

case "network" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
    check = {
      "nios.network" = "{{random_cidr_network}}"
    }
  }

}

case "network_view" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network      = "{{random_cidr_network}}"
      network_view = "default"
    }
    check = {
      "nios.network_view" = "default"
      "nios.network"      = "{{random_cidr_network}}"
    }
  }

}

case "nextserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network    = "{{random_cidr_network}}"
      nextserver = "1.1.1.1"
    }
    check = {
      "nios.nextserver" = "1.1.1.1"
      "nios.network"    = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      nextserver = "1.1.1.2"
    }
    check = {
      "nios.nextserver" = "1.1.1.2"
      "nios.network"    = "{{random_cidr_network}}"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
      options = [{ name = "time-offset", num = 2, value = "50" }, { name = "subnet-mask", value = "1.1.1.1" }]
    }
    check = {
      "nios.options.0.name"  = "time-offset"
      "nios.options.0.num"   = "2"
      "nios.options.0.value" = "50"
      "nios.options.1.name"  = "subnet-mask"
      "nios.options.1.value" = "1.1.1.1"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
      options = [{ name = "dhcp-lease-time", value = "7200" }, { name = "subnet-mask", value = "1.1.1.1" }]
    }
    check = {
      "nios.options.0.name"  = "dhcp-lease-time"
      "nios.options.0.value" = "7200"
      "nios.options.1.name"  = "subnet-mask"
      "nios.options.1.value" = "1.1.1.1"
    }
  }

}

case "port_control_blackout_setting" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                       = "{{random_cidr_network}}"
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
      network                       = "{{random_cidr_network}}"
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

case "pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network        = "{{random_cidr_network}}"
      pxe_lease_time = 0
    }
    check = {
      "nios.pxe_lease_time" = "0"
      "nios.network"        = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network        = "{{random_cidr_network}}"
      pxe_lease_time = 40000
    }
    check = {
      "nios.pxe_lease_time" = "40000"
      "nios.network"        = "{{random_cidr_network}}"
    }
  }

}

case "recycle_leases" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network        = "{{random_cidr_network}}"
      recycle_leases = false
    }
    check = {
      "nios.recycle_leases" = "false"
      "nios.network"        = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network        = "{{random_cidr_network}}"
      recycle_leases = true
    }
    check = {
      "nios.recycle_leases" = "true"
      "nios.network"        = "{{random_cidr_network}}"
    }
  }

}

case "restart_if_needed" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network           = "{{random_cidr_network}}"
      restart_if_needed = false
    }
    check = {
      "nios.restart_if_needed" = "false"
    }
  }

  step {
    nios {
      network           = "{{random_cidr_network}}"
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
  resource "infoblox_networkcontainer" "rir_parent" {
    nios = {
      network          = "10.{{random_octet}}.0.0/16"
      rir_organization = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET"
        "RIPE Description"       = "Test network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
  }
  PREREQ

  step {
    nios {
      network                 = "10.{{random_octet}}.0.0/24"
      rir_registration_action = "CREATE"
      rir_organization        = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-CHILD"
        "RIPE Description"       = "Test child network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
    depends_on = [infoblox_networkcontainer.rir_parent]
    check = {
      "nios.rir_registration_action" = "CREATE"
      "nios.network"                 = "10.{{random_octet}}.0.0/24"
    }
  }

  step {
    nios {
      network                 = "10.{{random_octet}}.0.0/24"
      rir_registration_action = "NONE"
      rir_organization        = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-CHILD"
        "RIPE Description"       = "Test child network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
    depends_on = [infoblox_networkcontainer.rir_parent]
    check = {
      "nios.rir_registration_action" = "NONE"
      "nios.network"                 = "10.{{random_octet}}.0.0/24"
    }
  }

}

case "rir_registration_status" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                              = "{{random_cidr_network}}"
      rir_registration_status              = "NOT_REGISTERED"
      same_port_control_discovery_blackout = false
    }
    check = {
      "nios.rir_registration_status"              = "NOT_REGISTERED"
      "nios.same_port_control_discovery_blackout" = "false"
      "nios.network"                              = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                              = "{{random_cidr_network}}"
      rir_registration_status              = "NOT_REGISTERED"
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.rir_registration_status"              = "NOT_REGISTERED"
      "nios.same_port_control_discovery_blackout" = "true"
      "nios.network"                              = "{{random_cidr_network}}"
    }
  }

}

case "same_port_control_discovery_blackout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                              = "{{random_cidr_network}}"
      same_port_control_discovery_blackout = false
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "false"
      "nios.network"                              = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                              = "{{random_cidr_network}}"
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "true"
      "nios.network"                              = "{{random_cidr_network}}"
    }
  }

}

case "send_rir_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_cidr_network}}"
      send_rir_request = false
    }
    check = {
      "nios.send_rir_request" = "false"
    }
  }

  step {
    nios {
      network          = "{{random_cidr_network}}"
      send_rir_request = true
    }
    check = {
      "nios.send_rir_request" = "true"
    }
  }

}

case "subscribe_settings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_cidr_network}}"
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
      network            = "{{random_cidr_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["SECURITY_GROUP"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "SECURITY_GROUP"
    }
  }

}

case "unmanaged" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network   = "{{random_cidr_network}}"
      unmanaged = false
    }
    check = {
      "nios.unmanaged" = "false"
      "nios.network"   = "{{random_cidr_network}}"
    }
  }

}

case "update_dns_on_lease_renewal" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                     = "{{random_cidr_network}}"
      update_dns_on_lease_renewal = false
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "false"
      "nios.network"                     = "{{random_cidr_network}}"
    }
  }

  step {
    nios {
      network                     = "{{random_cidr_network}}"
      update_dns_on_lease_renewal = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
      "nios.network"                     = "{{random_cidr_network}}"
    }
  }

}

case "discovery_member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_cidr_network}}"
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.discovery_member" = "{{discovery_member_hostname}}"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
  }

}

case "enable_discovery" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_cidr_network}}"
      discovery_member = "{{discovery_member_hostname}}"
      enable_discovery = true
    }
    check = {
      "nios.enable_discovery" = "true"
    }
  }

  step {
    nios {
      network          = "{{random_cidr_network}}"
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
      network                    = "{{random_cidr_network}}"
      discovery_member           = "{{discovery_member_hostname}}"
      enable_immediate_discovery = true
    }
    check = {
      "nios.enable_immediate_discovery" = "true"
    }
  }

  step {
    nios {
      network                    = "{{random_cidr_network}}"
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
      network            = "{{random_cidr_network}}"
      logic_filter_rules = [{ filter = "mac_filter", type = "MAC" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "mac_filter"
      "nios.logic_filter_rules.0.type"   = "MAC"
    }
  }

  step {
    nios {
      network            = "{{random_cidr_network}}"
      logic_filter_rules = [{ filter = "example-option-filter-1", type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "example-option-filter-1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "rir_organization" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_cidr_network}}"
      rir_organization = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET"
        "RIPE Description"       = "Test network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
    check = {
      "nios.rir_organization" = "rir-org-test1"
    }
  }

}

case "rir_organization_action" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_networkcontainer" "rir_parent" {
    nios = {
      network          = "11.{{random_octet}}.0.0/16"
      rir_organization = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET"
        "RIPE Description"       = "Test network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
  }
  PREREQ

  step {
    nios {
      network                 = "11.{{random_octet}}.0.0/24"
      rir_registration_action = "CREATE"
      rir_organization        = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-CHILD"
        "RIPE Description"       = "Test child network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
    depends_on = [infoblox_networkcontainer.rir_parent]
    check = {
      "nios.rir_registration_action" = "CREATE"
    }
  }

  step {
    nios {
      network                 = "11.{{random_octet}}.0.0/24"
      rir_registration_action = "NONE"
      rir_organization        = "rir-org-test1"
      ext_attrs = {
        "RIPE Network Name"      = "TEST-NET-CHILD"
        "RIPE Description"       = "Test child network"
        "RIPE Country"           = "United States (US)"
        "RIPE Admin Contact"     = "TEST-RIPE"
        "RIPE Technical Contact" = "TEST-RIPE"
        "RIPE Registry Source"   = "RIPE"
        "RIPE IPv4 Status"       = "ASSIGNED PA"
      }
    }
    depends_on = [infoblox_networkcontainer.rir_parent]
    check = {
      "nios.rir_registration_action" = "NONE"
    }
  }

}

case "mapped_ea_attributes" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network            = "{{random_cidr_network}}"
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
      network            = "{{random_cidr_network}}"
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
      network            = "{{random_cidr_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["USERNAME"], mapped_ea_attributes = [{ name = "MAC", mapped_ea = "Building" }] }
    }
    check = {
      "nios.subscribe_settings.mapped_ea_attributes.0.name"      = "MAC"
      "nios.subscribe_settings.mapped_ea_attributes.0.mapped_ea" = "Building"
    }
  }

}


case "next_available_network" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_networkcontainer" "alloc_parent" {
    nios = {
      network      = "10.{{random_octet}}.0.0/16"
      network_view = "default"
    }
  }
  PREREQ

  step {
    nios {
      dynamic_allocation = {
        network      = infoblox_networkcontainer.alloc_parent.nios.network
        network_view = "default"
        cidr         = 24
      }
      comment = "Created by Dynamic Allocation"
    }
    depends_on = [infoblox_networkcontainer.alloc_parent]
    check = {
      "nios.comment"      = "Created by Dynamic Allocation"
      "nios.network_view" = "default"
    }
  }

}
