// Objects to be present on the grid for testing
// mac_filter - Filter mac
// example-option-filter-1 - Filter Option
// rir-org-test1 - RIR Organization
// ISE Server has to be configured
// A discovery member has to be configured
// Vlan(s) has to be configured - test-vlan-for-network and test-vlan-2-for-network

# Auto-generated resource acceptance-test cases for Network.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
    check = {
      "nios.network"                              = "{{random_cidr_network}}"
      "nios.network_view"                         = "default"
      "nios.authority"                            = "false"
      "nios.cloud_shared"                         = "false"
      "nios.ddns_generate_hostname"               = "false"
      "nios.ddns_server_always_updates"           = "true"
      "nios.ddns_ttl"                             = "0"
      "nios.ddns_update_fixed_addresses"          = "false"
      "nios.ddns_use_option81"                    = "false"
      "nios.deny_bootp"                           = "false"
      "nios.disable"                              = "false"
      "nios.enable_ddns"                          = "false"
      "nios.enable_dhcp_thresholds"               = "false"
      "nios.enable_discovery"                     = "false"
      "nios.enable_email_warnings"                = "false"
      "nios.enable_ifmap_publishing"              = "false"
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
      "nios.network"   = "{{random_cidr_network}}"
      "nios.authority" = "false"
    }
  }

  step {
    nios {
      network   = "{{random_cidr_network}}"
      authority = true
    }
    check = {
      "nios.network"   = "{{random_cidr_network}}"
      "nios.authority" = "true"
    }
  }

}

case "auto_create_reversezone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                 = "10.{{random_octet}}.{{random_octet}}.0/24"
      auto_create_reversezone = true
    }
    check = {
      "nios.network"                 = "10.{{random_octet}}.{{random_octet}}.0/24"
      "nios.auto_create_reversezone" = "true"
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
      "nios.network"  = "{{random_cidr_network}}"
      "nios.bootfile" = "bootfile"
    }
  }

  step {
    nios {
      network  = "{{random_cidr_network}}"
      bootfile = "bootfile_updated"
    }
    check = {
      "nios.network"  = "{{random_cidr_network}}"
      "nios.bootfile" = "bootfile_updated"
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
      "nios.network"    = "{{random_cidr_network}}"
      "nios.bootserver" = "test_bootserver"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      bootserver = "test_bootserver_updated"
    }
    check = {
      "nios.network"    = "{{random_cidr_network}}"
      "nios.bootserver" = "test_bootserver_updated"
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

case "cloud_shared" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network      = "{{random_cidr_network}}"
      cloud_shared = false
    }
    check = {
      "nios.network"      = "{{random_cidr_network}}"
      "nios.cloud_shared" = "false"
    }
  }

  step {
    nios {
      network      = "{{random_cidr_network}}"
      cloud_shared = true
    }
    check = {
      "nios.network"      = "{{random_cidr_network}}"
      "nios.cloud_shared" = "true"
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
      "nios.network" = "{{random_cidr_network}}"
      "nios.comment" = "test comment"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
      comment = "updated comment"
    }
    check = {
      "nios.network" = "{{random_cidr_network}}"
      "nios.comment" = "updated comment"
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
      "nios.network"         = "{{random_cidr_network}}"
      "nios.ddns_domainname" = "test.com"
    }
  }

  step {
    nios {
      network         = "{{random_cidr_network}}"
      ddns_domainname = "testupdated.com"
    }
    check = {
      "nios.network"         = "{{random_cidr_network}}"
      "nios.ddns_domainname" = "testupdated.com"
    }
  }

}

case "ddns_generate_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                = "{{random_cidr_network}}"
      ddns_generate_hostname = true
    }
    check = {
      "nios.network"                = "{{random_cidr_network}}"
      "nios.ddns_generate_hostname" = "true"
    }
  }

  step {
    nios {
      network                = "{{random_cidr_network}}"
      ddns_generate_hostname = false
    }
    check = {
      "nios.network"                = "{{random_cidr_network}}"
      "nios.ddns_generate_hostname" = "false"
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
      "nios.network"                    = "{{random_cidr_network}}"
      "nios.ddns_server_always_updates" = "true"
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
      "nios.network"                    = "{{random_cidr_network}}"
      "nios.ddns_server_always_updates" = "false"
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
      "nios.network"  = "{{random_cidr_network}}"
      "nios.ddns_ttl" = "1"
    }
  }

  step {
    nios {
      network  = "{{random_cidr_network}}"
      ddns_ttl = 600
    }
    check = {
      "nios.network"  = "{{random_cidr_network}}"
      "nios.ddns_ttl" = "600"
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
      "nios.network"                     = "{{random_cidr_network}}"
      "nios.ddns_update_fixed_addresses" = "true"
    }
  }

  step {
    nios {
      network                     = "{{random_cidr_network}}"
      ddns_update_fixed_addresses = false
    }
    check = {
      "nios.network"                     = "{{random_cidr_network}}"
      "nios.ddns_update_fixed_addresses" = "false"
    }
  }

}

case "ddns_use_option81" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network           = "{{random_cidr_network}}"
      ddns_use_option81 = true
    }
    check = {
      "nios.network"           = "{{random_cidr_network}}"
      "nios.ddns_use_option81" = "true"
    }
  }

  step {
    nios {
      network           = "{{random_cidr_network}}"
      ddns_use_option81 = false
    }
    check = {
      "nios.network"           = "{{random_cidr_network}}"
      "nios.ddns_use_option81" = "false"
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
      "nios.network"    = "{{random_cidr_network}}"
      "nios.deny_bootp" = "false"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      deny_bootp = true
    }
    check = {
      "nios.network"    = "{{random_cidr_network}}"
      "nios.deny_bootp" = "true"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
      disable = false
    }
    check = {
      "nios.network" = "{{random_cidr_network}}"
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
      disable = true
    }
    check = {
      "nios.network" = "{{random_cidr_network}}"
      "nios.disable" = "true"
    }
  }

}

case "discovered_bridge_domain" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                  = "{{random_cidr_network}}"
      discovered_bridge_domain = "bridge-domain-1"
    }
    check = {
      "nios.network"                  = "{{random_cidr_network}}"
      "nios.discovered_bridge_domain" = "bridge-domain-1"
    }
  }

  step {
    nios {
      network                  = "{{random_cidr_network}}"
      discovered_bridge_domain = "bridge-domain-2"
    }
    check = {
      "nios.network"                  = "{{random_cidr_network}}"
      "nios.discovered_bridge_domain" = "bridge-domain-2"
    }
  }

}

case "discovered_tenant" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network           = "{{random_cidr_network}}"
      discovered_tenant = "tenant-1"
    }
    check = {
      "nios.discovered_tenant" = "tenant-1"
    }
  }

  step {
    nios {
      network           = "{{random_cidr_network}}"
      discovered_tenant = "tenant-2"
    }
    check = {
      "nios.discovered_tenant" = "tenant-2"
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
      "nios.network"                                                                   = "{{random_cidr_network}}"
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
      "nios.network"                                                                   = "{{random_cidr_network}}"
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
      "nios.network"                                                        = "{{random_cidr_network}}"
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

case "discovery_member" {
  backend     = "nios"
  parallel    = true
  
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

case "email_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network    = "{{random_cidr_network}}"
      email_list = ["admin@example.com"]
    }
    check = {
      "nios.network"      = "{{random_cidr_network}}"
      "nios.email_list.0" = "admin@example.com"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      email_list = ["admin@updated.com"]
    }
    check = {
      "nios.network"      = "{{random_cidr_network}}"
      "nios.email_list.0" = "admin@updated.com"
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
      "nios.network"     = "{{random_cidr_network}}"
      "nios.enable_ddns" = "false"
    }
  }

  step {
    nios {
      network     = "{{random_cidr_network}}"
      enable_ddns = true
    }
    check = {
      "nios.network"     = "{{random_cidr_network}}"
      "nios.enable_ddns" = "true"
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
      "nios.network"                = "{{random_cidr_network}}"
      "nios.enable_dhcp_thresholds" = "false"
    }
  }

  step {
    nios {
      network                = "{{random_cidr_network}}"
      enable_dhcp_thresholds = true
    }
    check = {
      "nios.network"                = "{{random_cidr_network}}"
      "nios.enable_dhcp_thresholds" = "true"
    }
  }

}

case "enable_discovery" {
  backend     = "nios"
  parallel    = true
  
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
      "nios.network"               = "{{random_cidr_network}}"
      "nios.enable_email_warnings" = "false"
    }
  }

  step {
    nios {
      network               = "{{random_cidr_network}}"
      enable_email_warnings = true
    }
    check = {
      "nios.network"               = "{{random_cidr_network}}"
      "nios.enable_email_warnings" = "true"
    }
  }

}

case "enable_ifmap_publishing" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      enable_ifmap_publishing = true
    }
    check = {
      "nios.network"                 = "{{random_cidr_network}}"
      "nios.enable_ifmap_publishing" = "true"
    }
  }

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      enable_ifmap_publishing = false
    }
    check = {
      "nios.network"                 = "{{random_cidr_network}}"
      "nios.enable_ifmap_publishing" = "false"
    }
  }

}

case "enable_immediate_discovery" {
  backend     = "nios"
  parallel    = true
  
  step {
    nios {
      network                    = "{{random_cidr_network}}"
      enable_immediate_discovery = true
    }
    check = {
      "nios.enable_immediate_discovery" = "true"
    }
  }

  step {
    nios {
      network                    = "{{random_cidr_network}}"
      enable_immediate_discovery = false
    }
    check = {
      "nios.enable_immediate_discovery" = "false"
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
      "nios.network"               = "{{random_cidr_network}}"
      "nios.pxe_lease_time"        = "100"
      "nios.enable_pxe_lease_time" = "false"
    }
  }

  step {
    nios {
      network               = "{{random_cidr_network}}"
      pxe_lease_time        = 100
      enable_pxe_lease_time = true
    }
    check = {
      "nios.network"               = "{{random_cidr_network}}"
      "nios.pxe_lease_time"        = "100"
      "nios.enable_pxe_lease_time" = "true"
    }
  }

}

case "enable_snmp_warnings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network              = "{{random_cidr_network}}"
      enable_snmp_warnings = true
    }
    check = {
      "nios.network"              = "{{random_cidr_network}}"
      "nios.enable_snmp_warnings" = "true"
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      enable_snmp_warnings = false
    }
    check = {
      "nios.network"              = "{{random_cidr_network}}"
      "nios.enable_snmp_warnings" = "false"
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
      "nios.network"        = "{{random_cidr_network}}"
      "nios.ext_attrs.Site" = "{{random}}"
    }
  }

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.network"        = "{{random_cidr_network}}"
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

}

case "high_water_mark" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network         = "{{random_cidr_network}}"
      high_water_mark = 80
    }
    check = {
      "nios.network"         = "{{random_cidr_network}}"
      "nios.high_water_mark" = "80"
    }
  }

  step {
    nios {
      network         = "{{random_cidr_network}}"
      high_water_mark = 90
    }
    check = {
      "nios.network"         = "{{random_cidr_network}}"
      "nios.high_water_mark" = "90"
    }
  }

}

case "high_water_mark_reset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network               = "{{random_cidr_network}}"
      high_water_mark_reset = 70
    }
    check = {
      "nios.network"               = "{{random_cidr_network}}"
      "nios.high_water_mark_reset" = "70"
    }
  }

  step {
    nios {
      network               = "{{random_cidr_network}}"
      high_water_mark_reset = 80
    }
    check = {
      "nios.network"               = "{{random_cidr_network}}"
      "nios.high_water_mark_reset" = "80"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                         = "{{random_cidr_network}}"
      ignore_dhcp_option_list_request = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
    }
  }

  step {
    nios {
      network                         = "{{random_cidr_network}}"
      ignore_dhcp_option_list_request = false
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
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
    }
  }

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ignore_id = "MACADDR"
    }
    check = {
      "nios.ignore_id" = "MACADDR"
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
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      ignore_mac_addresses = ["ff:ee:dd:cc:bb:aa"]
    }
    check = {
      "nios.ignore_mac_addresses.0" = "ff:ee:dd:cc:bb:aa"
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
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      ipam_email_addresses = ["testuser2@infoblox.com"]
    }
    check = {
      "nios.ipam_email_addresses.0" = "testuser2@infoblox.com"
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
      "nios.network"                               = "{{random_cidr_network}}"
      "nios.ipam_threshold_settings.reset_value"   = "85"
      "nios.ipam_threshold_settings.trigger_value" = "95"
    }
  }

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      ipam_threshold_settings = { reset_value = 75, trigger_value = 80 }
    }
    check = {
      "nios.network"                               = "{{random_cidr_network}}"
      "nios.ipam_threshold_settings.reset_value"   = "75"
      "nios.ipam_threshold_settings.trigger_value" = "80"
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
      "nios.network"                                  = "{{random_cidr_network}}"
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
      "nios.network"                                  = "{{random_cidr_network}}"
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
      "nios.network"             = "{{random_cidr_network}}"
      "nios.lease_scavenge_time" = "-1"
    }
  }

  step {
    nios {
      network             = "{{random_cidr_network}}"
      lease_scavenge_time = 86400
    }
    check = {
      "nios.network"             = "{{random_cidr_network}}"
      "nios.lease_scavenge_time" = "86400"
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

case "low_water_mark" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network        = "{{random_cidr_network}}"
      low_water_mark = 0
    }
    check = {
      "nios.network"        = "{{random_cidr_network}}"
      "nios.low_water_mark" = "0"
    }
  }

  step {
    nios {
      network        = "{{random_cidr_network}}"
      low_water_mark = 50
    }
    check = {
      "nios.network"        = "{{random_cidr_network}}"
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
      "nios.network"              = "{{random_cidr_network}}"
      "nios.low_water_mark_reset" = "10"
    }
  }

  step {
    nios {
      network              = "{{random_cidr_network}}"
      low_water_mark_reset = 20
    }
    check = {
      "nios.network"              = "{{random_cidr_network}}"
      "nios.low_water_mark_reset" = "20"
    }
  }

}

case "members" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
      members = [{ struct = "dhcpmember", name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.network"          = "{{random_cidr_network}}"
      "nios.members.0.struct" = "dhcpmember"
      "nios.members.0.name"   = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
      members = [{ struct = "msdhcpserver", ipv4addr = "10.10.10.10" }]
    }
    check = {
      "nios.network"            = "{{random_cidr_network}}"
      "nios.members.0.struct"   = "msdhcpserver"
      "nios.members.0.ipv4addr" = "10.10.10.10"
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
    }
  }

  step {
    nios {
      network     = "{{random_cidr_network}}"
      mgm_private = true
    }
    check = {
      "nios.mgm_private" = "true"
    }
  }

}

case "netmask" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network = "{{random_cidr_network}}"
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
      "nios.network"    = "{{random_cidr_network}}"
      "nios.nextserver" = "1.1.1.1"
    }
  }

  step {
    nios {
      network    = "{{random_cidr_network}}"
      nextserver = "1.1.1.2"
    }
    check = {
      "nios.network"    = "{{random_cidr_network}}"
      "nios.nextserver" = "1.1.1.2"
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
      options = [{ num = "51", value = "7200" }, { name = "subnet-mask", value = "1.1.1.1" }]
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
    }
  }

  step {
    nios {
      network        = "{{random_cidr_network}}"
      pxe_lease_time = 40000
    }
    check = {
      "nios.pxe_lease_time" = "40000"
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
    }
  }

  step {
    nios {
      network        = "{{random_cidr_network}}"
      recycle_leases = true
    }
    check = {
      "nios.recycle_leases" = "true"
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

case "rir_organization" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network          = "{{random_cidr_network}}"
      rir_organization = "rir-org-test1"
      ext_attrs = { "RIPE Network Name" = "test-network", "RIPE Description" = "test description", "RIPE Country" = "United States (US)", "RIPE Admin Contact" = "IB-RIPE", "RIPE Technical Contact" = "IB-RIPE", "RIPE IPv4 Status" = "ASSIGNED PA", "RIPE Registry Source" = "TEST" }
    }
    check = {
      "nios.rir_organization"            = "rir-org-test1"
      "nios.ext_attrs.RIPE Network Name" = "test-network"
      "nios.ext_attrs.RIPE Description"  = "test description"
    }
  }

  step {
    nios {
      network          = "{{random_cidr_network}}"
      rir_organization = "rir-org-test1"
      ext_attrs = { "RIPE Network Name" = "updated-network", "RIPE Description" = "updated description", "RIPE Country" = "United States (US)", "RIPE Admin Contact" = "IB-RIPE", "RIPE Technical Contact" = "IB-RIPE", "RIPE IPv4 Status" = "ASSIGNED PA", "RIPE Registry Source" = "TEST" }
    }
    check = {
      "nios.rir_organization"            = "rir-org-test1"
      "nios.ext_attrs.RIPE Network Name" = "updated-network"
      "nios.ext_attrs.RIPE Description"  = "updated description"
    }
  }

}

case "rir_registration_action" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      rir_registration_action = "NONE"
      comment                 = "initial comment"
    }
    check = {
      "nios.rir_registration_action" = "NONE"
      "nios.comment"                 = "initial comment"
    }
  }

  step {
    nios {
      network                 = "{{random_cidr_network}}"
      rir_registration_action = "NONE"
      comment                 = "updated comment"
    }
    check = {
      "nios.rir_registration_action" = "NONE"
      "nios.comment"                 = "updated comment"
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
    }
  }

  step {
    nios {
      network                              = "{{random_cidr_network}}"
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "true"
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
  backend     = "nios"
  parallel    = true

  step {
    nios {
      network            = "{{random_cidr_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["DOMAINNAME"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "DOMAINNAME"
    }
  }

  step {
    nios {
      network            = "{{random_cidr_network}}"
      network_view       = "test_network_view"
      subscribe_settings = { enabled_attributes = ["ENDPOINT_PROFILE"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "ENDPOINT_PROFILE"
    }
  }

}

case "template" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_network_template_unknown" "test_net_tmpl" {
  #   nios = {
  #     name = "{{random}}"
  #     netmask = 24
  #   }
  # }
  # PREREQ
  #
  step {
    nios {
      # template = infoblox_network_template_unknown.test_net_tmpl.nios.name
      network  = "10.{{random_octet}}.{{random_octet}}.0/24"
      template = "test-networktemplate-for-network"
    }
    # depends_on = [infoblox_network_template_unknown.test_net_tmpl]
    check = {
      "nios.network"  = "10.{{random_octet}}.{{random_octet}}.0/24"
      "nios.template" = "test-networktemplate-for-network"
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
    }
  }

  step {
    nios {
      network                     = "{{random_cidr_network}}"
      update_dns_on_lease_renewal = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
    }
  }

}

case "vlans" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_vlan_view" "test_vlan_view" {
  #   nios = {
  #     start_vlan_id = 50
  #     end_vlan_id = 100
  #     name = "test-vlanview-for-network"
  #   }
  # }
  # resource "infoblox_vlan" "test_vlan" {
  #   nios = {
  #     id = 50
  #     name = "test-vlan-for-network"
  #     parent = infoblox_vlan_view.test_vlan_view.nios.ref
  #   }
  # }
  # PREREQ

  step {
    nios {
      network = "{{random_cidr_network}}"
      # vlans   = [{ vlan = infoblox_vlan.test_vlan.nios.ref }]
      vlans = [{ vlan = "vlan/ZG5zLnZsYW4kLmNvbS5pbmZvYmxveC5kbnMudmxhbl92aWV3JHRlc3QtdmxhbnZpZXctZm9yLW5ldHdvcmsuNTAuMTAwLjUw:test-vlanview-for-network/test-vlan-for-network/50" }]
    }
    check = {
      "nios.vlans.0.vlan" = "vlan/ZG5zLnZsYW4kLmNvbS5pbmZvYmxveC5kbnMudmxhbl92aWV3JHRlc3QtdmxhbnZpZXctZm9yLW5ldHdvcmsuNTAuMTAwLjUw:test-vlanview-for-network/test-vlan-for-network/50"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
      vlans = [{ vlan = "vlan/ZG5zLnZsYW4kLmNvbS5pbmZvYmxveC5kbnMudmxhbl92aWV3JHRlc3QtdmxhbnZpZXctZm9yLW5ldHdvcmsuNTAuMTAwLjUx:test-vlanview-for-network/test-vlan-2-for-network/51" }]
    }
    check = {
      "nios.vlans.0.vlan" = "vlan/ZG5zLnZsYW4kLmNvbS5pbmZvYmxveC5kbnMudmxhbl92aWV3JHRlc3QtdmxhbnZpZXctZm9yLW5ldHdvcmsuNTAuMTAwLjUx:test-vlanview-for-network/test-vlan-2-for-network/51"
    }
  }

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
    check = {
      "nios.network" = "{{random_cidr_network}}"
    }
  }

}
