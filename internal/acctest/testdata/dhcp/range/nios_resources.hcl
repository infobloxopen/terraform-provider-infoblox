# Auto-generated resource acceptance-test cases for Range.
case "basic" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.11"
      end_addr     = "10.0.0.12"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.start_addr"                      = "10.0.0.11"
      "nios.end_addr"                        = "10.0.0.12"
      "nios.always_update_dns"               = "false"
      "nios.enable_ddns"                     = "false"
      "nios.enable_discovery"                = "false"
      "nios.enable_ifmap_publishing"         = "false"
      "nios.enable_pxe_lease_time"           = "false"
      "nios.enable_snmp_warnings"            = "false"
      "nios.high_water_mark"                 = "95"
      "nios.high_water_mark_reset"           = "85"
      "nios.ignore_dhcp_option_list_request" = "false"
      "nios.lease_scavenge_time"             = "-1"
      "nios.low_water_mark"                  = "0"
      "nios.low_water_mark_reset"            = "10"
      "nios.network_view"                    = "{{random_view}}"
      "nios.update_dns_on_lease_renewal"     = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.13"
      end_addr     = "10.0.0.14"
      network_view = infoblox_network.test_network.nios.network_view
    }
  }

}

case "always_update_dns" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr        = "10.0.0.15"
      end_addr          = "10.0.0.16"
      network_view      = infoblox_network.test_network.nios.network_view
      always_update_dns = false
    }
    check = {
      "nios.always_update_dns" = "false"
    }
  }

  step {
    nios {
      start_addr        = "10.0.0.15"
      end_addr          = "10.0.0.16"
      network_view      = infoblox_network.test_network.nios.network_view
      always_update_dns = true
    }
    check = {
      "nios.always_update_dns" = "true"
    }
  }

}

case "bootfile" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.17"
      end_addr     = "10.0.0.18"
      network_view = infoblox_network.test_network.nios.network_view
      bootfile     = "boot.ini"
    }
    check = {
      "nios.bootfile" = "boot.ini"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.17"
      end_addr     = "10.0.0.18"
      network_view = infoblox_network.test_network.nios.network_view
      bootfile     = "bootfile_update.ini"
    }
    check = {
      "nios.bootfile" = "bootfile_update.ini"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.19"
      end_addr     = "10.0.0.20"
      network_view = infoblox_network.test_network.nios.network_view
      bootserver   = "bootserver"
    }
    check = {
      "nios.bootserver" = "bootserver"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.19"
      end_addr     = "10.0.0.20"
      network_view = infoblox_network.test_network.nios.network_view
      bootserver   = "bootserverupdate"
    }
    check = {
      "nios.bootserver" = "bootserverupdate"
    }
  }

}

case "cloud_info" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.21"
      end_addr     = "10.0.0.22"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.cloud_info.authority_type"   = "GM"
      "nios.cloud_info.delegated_scope"  = "NONE"
      "nios.cloud_info.owned_by_adaptor" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.23"
      end_addr     = "10.0.0.24"
      network_view = infoblox_network.test_network.nios.network_view
      comment      = "network range"
    }
    check = {
      "nios.comment" = "network range"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.23"
      end_addr     = "10.0.0.24"
      network_view = infoblox_network.test_network.nios.network_view
      comment      = "network range updated"
    }
    check = {
      "nios.comment" = "network range updated"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr      = "10.0.0.25"
      end_addr        = "10.0.0.26"
      network_view    = infoblox_network.test_network.nios.network_view
      ddns_domainname = "yourdomain.com"
    }
    check = {
      "nios.ddns_domainname" = "yourdomain.com"
    }
  }

  step {
    nios {
      start_addr      = "10.0.0.25"
      end_addr        = "10.0.0.26"
      network_view    = infoblox_network.test_network.nios.network_view
      ddns_domainname = "yourdomainupdate.com"
    }
    check = {
      "nios.ddns_domainname" = "yourdomainupdate.com"
    }
  }

}

case "ddns_generate_hostname" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr             = "10.0.0.27"
      end_addr               = "10.0.0.28"
      network_view           = infoblox_network.test_network.nios.network_view
      ddns_generate_hostname = true
    }
    check = {
      "nios.ddns_generate_hostname" = "true"
    }
  }

  step {
    nios {
      start_addr             = "10.0.0.27"
      end_addr               = "10.0.0.28"
      network_view           = infoblox_network.test_network.nios.network_view
      ddns_generate_hostname = false
    }
    check = {
      "nios.ddns_generate_hostname" = "false"
    }
  }

}

case "deny_all_clients" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr       = "10.0.0.29"
      end_addr         = "10.0.0.30"
      network_view     = infoblox_network.test_network.nios.network_view
      deny_all_clients = true
    }
    check = {
      "nios.deny_all_clients" = "true"
    }
  }

  step {
    nios {
      start_addr       = "10.0.0.29"
      end_addr         = "10.0.0.30"
      network_view     = infoblox_network.test_network.nios.network_view
      deny_all_clients = false
    }
    check = {
      "nios.deny_all_clients" = "false"
    }
  }

}

case "deny_bootp" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.31"
      end_addr     = "10.0.0.32"
      network_view = infoblox_network.test_network.nios.network_view
      deny_bootp   = true
    }
    check = {
      "nios.deny_bootp" = "true"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.31"
      end_addr     = "10.0.0.32"
      network_view = infoblox_network.test_network.nios.network_view
      deny_bootp   = false
    }
    check = {
      "nios.deny_bootp" = "false"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.33"
      end_addr     = "10.0.0.34"
      network_view = infoblox_network.test_network.nios.network_view
      disable      = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.33"
      end_addr     = "10.0.0.34"
      network_view = infoblox_network.test_network.nios.network_view
      disable      = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "discovery_basic_poll_settings" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr                    = "10.0.0.35"
      end_addr                      = "10.0.0.36"
      network_view                  = infoblox_network.test_network.nios.network_view
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
      start_addr                    = "10.0.0.35"
      end_addr                      = "10.0.0.36"
      network_view                  = infoblox_network.test_network.nios.network_view
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
      start_addr                    = "10.0.0.35"
      end_addr                      = "10.0.0.36"
      network_view                  = infoblox_network.test_network.nios.network_view
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

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr                 = "10.0.0.37"
      end_addr                   = "10.0.0.38"
      network_view               = infoblox_network.test_network.nios.network_view
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
      start_addr                 = "10.0.0.37"
      end_addr                   = "10.0.0.38"
      network_view               = infoblox_network.test_network.nios.network_view
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

case "discovery_member" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr       = "10.0.0.39"
      end_addr         = "10.0.0.40"
      network_view     = infoblox_network.test_network.nios.network_view
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.discovery_member" = "{{discovery_member_hostname}}"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.39"
      end_addr     = "10.0.0.40"
      network_view = infoblox_network.test_network.nios.network_view
    }
  }

}

case "email_list" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.41"
      end_addr     = "10.0.0.42"
      network_view = infoblox_network.test_network.nios.network_view
      email_list   = ["example@infoblox.com"]
    }
    check = {
      "nios.email_list.#" = "1"
      "nios.email_list.0" = "example@infoblox.com"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.41"
      end_addr     = "10.0.0.42"
      network_view = infoblox_network.test_network.nios.network_view
      email_list   = ["example2@example.com"]
    }
    check = {
      "nios.email_list.0" = "example2@example.com"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.43"
      end_addr     = "10.0.0.44"
      network_view = infoblox_network.test_network.nios.network_view
      enable_ddns  = true
    }
    check = {
      "nios.enable_ddns" = "true"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.43"
      end_addr     = "10.0.0.44"
      network_view = infoblox_network.test_network.nios.network_view
      enable_ddns  = false
    }
    check = {
      "nios.enable_ddns" = "false"
    }
  }

}

case "enable_dhcp_thresholds" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr             = "10.0.0.45"
      end_addr               = "10.0.0.46"
      network_view           = infoblox_network.test_network.nios.network_view
      enable_dhcp_thresholds = true
    }
    check = {
      "nios.enable_dhcp_thresholds" = "true"
    }
  }

  step {
    nios {
      start_addr             = "10.0.0.45"
      end_addr               = "10.0.0.46"
      network_view           = infoblox_network.test_network.nios.network_view
      enable_dhcp_thresholds = false
    }
    check = {
      "nios.enable_dhcp_thresholds" = "false"
    }
  }

}

case "enable_email_warnings" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr            = "10.0.0.47"
      end_addr              = "10.0.0.48"
      network_view          = infoblox_network.test_network.nios.network_view
      enable_email_warnings = true
    }
    check = {
      "nios.enable_email_warnings" = "true"
    }
  }

  step {
    nios {
      start_addr            = "10.0.0.47"
      end_addr              = "10.0.0.48"
      network_view          = infoblox_network.test_network.nios.network_view
      enable_email_warnings = false
    }
    check = {
      "nios.enable_email_warnings" = "false"
    }
  }

}

case "enable_ifmap_publishing" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr              = "10.0.0.49"
      end_addr                = "10.0.0.50"
      network_view            = infoblox_network.test_network.nios.network_view
      enable_ifmap_publishing = true
    }
    check = {
      "nios.enable_ifmap_publishing" = "true"
    }
  }

  step {
    nios {
      start_addr              = "10.0.0.49"
      end_addr                = "10.0.0.50"
      network_view            = infoblox_network.test_network.nios.network_view
      enable_ifmap_publishing = false
    }
    check = {
      "nios.enable_ifmap_publishing" = "false"
    }
  }

}

case "enable_pxe_lease_time" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr            = "10.0.0.51"
      end_addr              = "10.0.0.52"
      network_view          = infoblox_network.test_network.nios.network_view
      enable_pxe_lease_time = true
      pxe_lease_time        = 3600
    }
    check = {
      "nios.enable_pxe_lease_time" = "true"
    }
  }

  step {
    nios {
      start_addr            = "10.0.0.51"
      end_addr              = "10.0.0.52"
      network_view          = infoblox_network.test_network.nios.network_view
      enable_pxe_lease_time = false
      pxe_lease_time        = 3600
    }
    check = {
      "nios.enable_pxe_lease_time" = "false"
    }
  }

}

case "enable_snmp_warnings" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      enable_snmp_warnings = true
      start_addr           = "10.0.0.53"
      end_addr             = "10.0.0.54"
      network_view         = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.enable_snmp_warnings" = "true"
    }
  }

  step {
    nios {
      enable_snmp_warnings = false
      start_addr           = "10.0.0.53"
      end_addr             = "10.0.0.54"
      network_view         = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.enable_snmp_warnings" = "false"
    }
  }

}

case "end_addr" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.55"
      end_addr     = "10.0.0.56"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.end_addr" = "10.0.0.56"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.55"
      end_addr     = "10.0.0.58"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.end_addr" = "10.0.0.58"
    }
  }

}

case "exclude" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.200"
      end_addr     = "10.0.0.210"
      network_view = infoblox_network.test_network.nios.network_view
      exclude      = [{ start_address = "10.0.0.202", end_address = "10.0.0.204" }]
    }
    check = {
      "nios.exclude.0.start_address" = "10.0.0.202"
      "nios.exclude.0.end_address"   = "10.0.0.204"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.200"
      end_addr     = "10.0.0.210"
      network_view = infoblox_network.test_network.nios.network_view
      exclude      = [{ start_address = "10.0.0.206", end_address = "10.0.0.209" }]
    }
    check = {
      "nios.exclude.0.start_address" = "10.0.0.206"
      "nios.exclude.0.end_address"   = "10.0.0.209"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.61"
      end_addr     = "10.0.0.62"
      network_view = infoblox_network.test_network.nios.network_view
      ext_attrs    = { Site = "{{random}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random}}"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.61"
      end_addr     = "10.0.0.62"
      network_view = infoblox_network.test_network.nios.network_view
      ext_attrs    = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

}

case "enable_discovery" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr       = "10.0.0.63"
      end_addr         = "10.0.0.64"
      network_view     = infoblox_network.test_network.nios.network_view
      enable_discovery = true
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.enable_discovery" = "true"
    }
  }

  step {
    nios {
      start_addr       = "10.0.0.63"
      end_addr         = "10.0.0.64"
      network_view     = infoblox_network.test_network.nios.network_view
      enable_discovery = false
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.enable_discovery" = "false"
    }
  }

}

case "enable_immediate_discovery" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr                 = "10.0.0.65"
      end_addr                   = "10.0.0.66"
      network_view               = infoblox_network.test_network.nios.network_view
      enable_immediate_discovery = true
    }
    check = {
      "nios.enable_immediate_discovery" = "true"
    }
  }

  step {
    nios {
      start_addr                 = "10.0.0.65"
      end_addr                   = "10.0.0.66"
      network_view               = infoblox_network.test_network.nios.network_view
      enable_immediate_discovery = false
    }
    check = {
      "nios.enable_immediate_discovery" = "false"
    }
  }

}

case "failover_association" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Requires non-grid master candidate to be in discovery polling mode"
  parallel    = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      failover_association    = "example_failover_association"
      start_addr              = "10.0.0.67"
      end_addr                = "10.0.0.68"
      network_view            = infoblox_network.test_network.nios.network_view
      server_association_type = "FAILOVER"
    }
    check = {
      "nios.failover_association" = "example_failover_association"
    }
  }

  step {
    nios {
      failover_association    = "example_failover_association1"
      start_addr              = "10.0.0.67"
      end_addr                = "10.0.0.68"
      network_view            = infoblox_network.test_network.nios.network_view
      server_association_type = "FAILOVER"
    }
    check = {
      "nios.failover_association" = "example_failover_association1"
    }
  }

}

case "fingerprint_filter_rules" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr               = "10.0.0.69"
      end_addr                 = "10.0.0.70"
      network_view             = infoblox_network.test_network.nios.network_view
      fingerprint_filter_rules = [{ filter = "test_filter_fingerprint", permission = "Allow" }]
    }
    check = {
      "nios.fingerprint_filter_rules.0.filter"     = "test_filter_fingerprint"
      "nios.fingerprint_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      start_addr               = "10.0.0.69"
      end_addr                 = "10.0.0.70"
      network_view             = infoblox_network.test_network.nios.network_view
      fingerprint_filter_rules = [{ filter = "test_filter_fingerprint1", permission = "Allow" }]
    }
    check = {
      "nios.fingerprint_filter_rules.0.filter"     = "test_filter_fingerprint1"
      "nios.fingerprint_filter_rules.0.permission" = "Allow"
    }
  }

}

case "high_water_mark" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr      = "10.0.0.71"
      end_addr        = "10.0.0.72"
      network_view    = infoblox_network.test_network.nios.network_view
      high_water_mark = 23
    }
    check = {
      "nios.high_water_mark" = "23"
    }
  }

  step {
    nios {
      start_addr      = "10.0.0.71"
      end_addr        = "10.0.0.72"
      network_view    = infoblox_network.test_network.nios.network_view
      high_water_mark = 42
    }
    check = {
      "nios.high_water_mark" = "42"
    }
  }

}

case "high_water_mark_reset" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      high_water_mark_reset = 23
      start_addr            = "10.0.0.73"
      end_addr              = "10.0.0.74"
      network_view          = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.high_water_mark_reset" = "23"
    }
  }

  step {
    nios {
      high_water_mark_reset = 42
      start_addr            = "10.0.0.73"
      end_addr              = "10.0.0.74"
      network_view          = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.high_water_mark_reset" = "42"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      ignore_dhcp_option_list_request = true
      start_addr                      = "10.0.0.75"
      end_addr                        = "10.0.0.76"
      network_view                    = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
    }
  }

  step {
    nios {
      ignore_dhcp_option_list_request = false
      start_addr                      = "10.0.0.75"
      end_addr                        = "10.0.0.76"
      network_view                    = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

case "ignore_id" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.177"
      end_addr     = "10.0.0.178"
      network_view = infoblox_network.test_network.nios.network_view
      ignore_id    = "CLIENT"
    }
    check = {
      "nios.ignore_id" = "CLIENT"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.177"
      end_addr     = "10.0.0.178"
      network_view = infoblox_network.test_network.nios.network_view
      ignore_id    = "MACADDR"
    }
    check = {
      "nios.ignore_id" = "MACADDR"
    }
  }

}

case "ignore_mac_addresses" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr           = "10.0.0.179"
      end_addr             = "10.0.0.180"
      network_view         = infoblox_network.test_network.nios.network_view
      ignore_mac_addresses = ["00:4a:2b:3c:1d:5e"]
    }
    check = {
      "nios.ignore_mac_addresses.0" = "00:4a:2b:3c:1d:5e"
    }
  }

  step {
    nios {
      start_addr           = "10.0.0.179"
      end_addr             = "10.0.0.180"
      network_view         = infoblox_network.test_network.nios.network_view
      ignore_mac_addresses = ["00:3a:2b:43:5d:52"]
    }
    check = {
      "nios.ignore_mac_addresses.0" = "00:3a:2b:43:5d:52"
    }
  }

}

case "known_clients" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr    = "10.0.0.81"
      end_addr      = "10.0.0.82"
      network_view  = infoblox_network.test_network.nios.network_view
      known_clients = "Deny"
    }
    check = {
      "nios.known_clients" = "Deny"
    }
  }

  step {
    nios {
      start_addr    = "10.0.0.81"
      end_addr      = "10.0.0.82"
      network_view  = infoblox_network.test_network.nios.network_view
      known_clients = "Allow"
    }
    check = {
      "nios.known_clients" = "Allow"
    }
  }

}

case "lease_scavenge_time" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr          = "10.0.0.83"
      end_addr            = "10.0.0.84"
      network_view        = infoblox_network.test_network.nios.network_view
      lease_scavenge_time = 86420
    }
    check = {
      "nios.lease_scavenge_time" = "86420"
    }
  }

  step {
    nios {
      start_addr          = "10.0.0.83"
      end_addr            = "10.0.0.84"
      network_view        = infoblox_network.test_network.nios.network_view
      lease_scavenge_time = 86430
    }
    check = {
      "nios.lease_scavenge_time" = "86430"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr         = "10.0.0.85"
      end_addr           = "10.0.0.86"
      network_view       = infoblox_network.test_network.nios.network_view
      logic_filter_rules = [{ filter = "mac_filter", type = "MAC" }]
    }
    check = {
      "nios.logic_filter_rules.0.filter" = "mac_filter"
      "nios.logic_filter_rules.0.type"   = "MAC"
    }
  }

  step {
    nios {
      start_addr         = "10.0.0.85"
      end_addr           = "10.0.0.86"
      network_view       = infoblox_network.test_network.nios.network_view
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

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr     = "10.0.0.87"
      end_addr       = "10.0.0.88"
      network_view   = infoblox_network.test_network.nios.network_view
      low_water_mark = 5
    }
    check = {
      "nios.low_water_mark" = "5"
    }
  }

  step {
    nios {
      start_addr     = "10.0.0.87"
      end_addr       = "10.0.0.88"
      network_view   = infoblox_network.test_network.nios.network_view
      low_water_mark = 1
    }
    check = {
      "nios.low_water_mark" = "1"
    }
  }

}

case "low_water_mark_reset" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      low_water_mark_reset = 5
      start_addr           = "10.0.0.89"
      end_addr             = "10.0.0.90"
      network_view         = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.low_water_mark_reset" = "5"
    }
  }

  step {
    nios {
      low_water_mark_reset = 1
      start_addr           = "10.0.0.89"
      end_addr             = "10.0.0.90"
      network_view         = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.low_water_mark_reset" = "1"
    }
  }

}

case "mac_filter_rules" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr       = "10.0.0.91"
      end_addr         = "10.0.0.92"
      network_view     = infoblox_network.test_network.nios.network_view
      mac_filter_rules = [{ filter = "mac_filter", permission = "Allow" }]
    }
    check = {
      "nios.mac_filter_rules.0.filter"     = "mac_filter"
      "nios.mac_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      start_addr       = "10.0.0.91"
      end_addr         = "10.0.0.92"
      network_view     = infoblox_network.test_network.nios.network_view
      mac_filter_rules = [{ filter = "mac_filter2", permission = "Deny" }]
    }
    check = {
      "nios.mac_filter_rules.0.filter"     = "mac_filter2"
      "nios.mac_filter_rules.0.permission" = "Deny"
    }
  }

}

case "member" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "102.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
      members      = [{ struct = "dhcpmember", name = "{{grid_master_hostname}}" }]
    }
  }
  PREREQ

  step {
    nios {
      start_addr              = "102.0.0.93"
      end_addr                = "102.0.0.94"
      network_view            = infoblox_network.test_network.nios.network_view
      member                  = { name = "{{grid_master_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.member.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      start_addr              = "102.0.0.93"
      end_addr                = "102.0.0.94"
      network_view            = infoblox_network.test_network.nios.network_view
      member                  = { name = "{{grid_member_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.member.name" = "{{grid_member_hostname}}"
    }
  }

}

case "ms_options" {
  backend           = "nios"
  skip              = true
  skip_reason       = "t.Skip: Skipping until MS Options can be properly tested with the API"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "100.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
      members      = [{ struct = "msdhcpserver", ipv4addr = "10.10.10.10" }]
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "100.0.0.182"
      end_addr     = "100.0.0.184"
      network_view = infoblox_network.test_network.nios.network_view
      ms_options   = [{ name = "domain-name", num = "15", value = "example.com" }, { name = "dhcp-lease-time", value = "7200" }]
      ms_server    = { ipv4addr = infoblox_network.test_network.nios.members[0].ipv4addr }
    }
    check = {
      "nios.ms_options" = "MS_OPTIONS_REPLACE_ME"
    }
  }

  step {
    nios {
      start_addr   = "100.0.0.182"
      end_addr     = "100.0.0.184"
      network_view = infoblox_network.test_network.nios.network_view
      ms_options   = [{ name = "braodcast-address", value = "127.0.0.1" }, { name = "dhcp-lease-time", value = "3600" }]
      ms_server    = { ipv4addr = infoblox_network.test_network.nios.members[0].ipv4addr }
    }
    check = {
      "nios.ms_options" = "MS_OPTIONS_UPDATE_REPLACE_ME"
    }
  }

}

case "ms_server" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "101.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
      members      = [{ struct = "msdhcpserver", ipv4addr = "10.10.10.10" }]
    }
  }
  PREREQ

  step {
    nios {
      start_addr              = "101.0.0.95"
      end_addr                = "101.0.0.96"
      network_view            = infoblox_network.test_network.nios.network_view
      ms_server               = { ipv4addr = infoblox_network.test_network.nios.members[0].ipv4addr }
      server_association_type = "MS_SERVER"
    }
    check = {
      "nios.ms_server.ipv4addr" = "10.10.10.10"
    }
  }

}

case "nac_filter_rules" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr       = "10.0.0.97"
      end_addr         = "10.0.0.98"
      network_view     = infoblox_network.test_network.nios.network_view
      nac_filter_rules = [{ filter = "nac_filter", permission = "Allow" }]
    }
    check = {
      "nios.nac_filter_rules.0.filter"     = "nac_filter"
      "nios.nac_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      start_addr       = "10.0.0.97"
      end_addr         = "10.0.0.98"
      network_view     = infoblox_network.test_network.nios.network_view
      nac_filter_rules = [{ filter = "nac_filter", permission = "Deny" }]
    }
    check = {
      "nios.nac_filter_rules.0.filter"     = "nac_filter"
      "nios.nac_filter_rules.0.permission" = "Deny"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.99"
      end_addr     = "10.0.0.100"
      network_view = infoblox_network.test_network.nios.network_view
      name         = "range 1"
    }
    check = {
      "nios.name" = "range 1"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.99"
      end_addr     = "10.0.0.100"
      network_view = infoblox_network.test_network.nios.network_view
      name         = "range 2"
    }
    check = {
      "nios.name" = "range 2"
    }
  }

}

case "network" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  resource "infoblox_network" "test_network2" {
    nios = {
      network      = "200.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.101"
      end_addr     = "10.0.0.102"
      network      = infoblox_network.test_network.nios.network
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.network" = "10.0.0.0/24"
    }
  }

  step {
    nios {
      start_addr   = "200.0.0.20"
      end_addr     = "200.0.0.30"
      network      = infoblox_network.test_network2.nios.network
      network_view = infoblox_network.test_network2.nios.network_view
    }
    check = {
      "nios.network" = "200.0.0.0/24"
    }
  }

}

case "network_view" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.103"
      end_addr     = "10.0.0.104"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.network_view" = "{{random_view}}"
    }
  }

}

case "nextserver" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.105"
      end_addr     = "10.0.0.106"
      network_view = infoblox_network.test_network.nios.network_view
      nextserver   = "next_server.com"
    }
    check = {
      "nios.nextserver" = "next_server.com"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.105"
      end_addr     = "10.0.0.106"
      network_view = infoblox_network.test_network.nios.network_view
      nextserver   = "next_server_update.com"
    }
    check = {
      "nios.nextserver" = "next_server_update.com"
    }
  }

}

case "option_filter_rules" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr          = "10.0.0.107"
      end_addr            = "10.0.0.108"
      network_view        = infoblox_network.test_network.nios.network_view
      option_filter_rules = [{ filter = "example-option-filter-1", permission = "Allow" }]
    }
    check = {
      "nios.option_filter_rules.0.filter"     = "example-option-filter-1"
      "nios.option_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      start_addr          = "10.0.0.107"
      end_addr            = "10.0.0.108"
      network_view        = infoblox_network.test_network.nios.network_view
      option_filter_rules = [{ filter = "example-option-filter-2", permission = "Deny" }]
    }
    check = {
      "nios.option_filter_rules.0.filter"     = "example-option-filter-2"
      "nios.option_filter_rules.0.permission" = "Deny"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.109"
      end_addr     = "10.0.0.110"
      network_view = infoblox_network.test_network.nios.network_view
      options      = [{ name = "time-offset", num = 2, value = "50" }, { name = "subnet-mask", value = "1.1.1.1" }]
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
      start_addr   = "10.0.0.109"
      end_addr     = "10.0.0.110"
      network_view = infoblox_network.test_network.nios.network_view
      options      = [{ name = "dhcp-lease-time", value = "7200" }, { name = "subnet-mask", value = "1.1.1.1" }]
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

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr                    = "10.0.0.111"
      end_addr                      = "10.0.0.112"
      network_view                  = infoblox_network.test_network.nios.network_view
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
      start_addr                    = "10.0.0.111"
      end_addr                      = "10.0.0.112"
      network_view                  = infoblox_network.test_network.nios.network_view
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

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr     = "10.0.0.113"
      end_addr       = "10.0.0.114"
      network_view   = infoblox_network.test_network.nios.network_view
      pxe_lease_time = 3400
    }
    check = {
      "nios.pxe_lease_time" = "3400"
    }
  }

  step {
    nios {
      start_addr     = "10.0.0.113"
      end_addr       = "10.0.0.114"
      network_view   = infoblox_network.test_network.nios.network_view
      pxe_lease_time = 3600
    }
    check = {
      "nios.pxe_lease_time" = "3600"
    }
  }

}

case "recycle_leases" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr     = "10.0.0.115"
      end_addr       = "10.0.0.116"
      network_view   = infoblox_network.test_network.nios.network_view
      recycle_leases = true
    }
    check = {
      "nios.recycle_leases" = "true"
    }
  }

  step {
    nios {
      start_addr     = "10.0.0.115"
      end_addr       = "10.0.0.116"
      network_view   = infoblox_network.test_network.nios.network_view
      recycle_leases = false
    }
    check = {
      "nios.recycle_leases" = "false"
    }
  }

}

case "relay_agent_filter_rules" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr               = "10.0.0.117"
      end_addr                 = "10.0.0.118"
      network_view             = infoblox_network.test_network.nios.network_view
      relay_agent_filter_rules = [{ filter = "relay_agent_filter", permission = "Allow" }]
    }
    check = {
      "nios.relay_agent_filter_rules.0.filter"     = "relay_agent_filter"
      "nios.relay_agent_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      start_addr               = "10.0.0.117"
      end_addr                 = "10.0.0.118"
      network_view             = infoblox_network.test_network.nios.network_view
      relay_agent_filter_rules = [{ filter = "relay_agent_filter", permission = "Deny" }]
    }
    check = {
      "nios.relay_agent_filter_rules.0.filter"     = "relay_agent_filter"
      "nios.relay_agent_filter_rules.0.permission" = "Deny"
    }
  }

}

case "same_port_control_discovery_blackout" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr                           = "10.0.0.119"
      end_addr                             = "10.0.0.120"
      network_view                         = infoblox_network.test_network.nios.network_view
      same_port_control_discovery_blackout = false
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "false"
    }
  }

  step {
    nios {
      start_addr                           = "10.0.0.119"
      end_addr                             = "10.0.0.120"
      network_view                         = infoblox_network.test_network.nios.network_view
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "true"
    }
  }

}

case "server_association_type" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "190.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
      members      = [{ struct = "dhcpmember", name = "{{grid_member_hostname}}" }]
    }
  }
  PREREQ

  step {
    nios {
      start_addr              = "190.0.0.121"
      end_addr                = "190.0.0.122"
      network_view            = infoblox_network.test_network.nios.network_view
      server_association_type = "FAILOVER"
      failover_association    = "example_failover_association1"
      member                  = { name = infoblox_network.test_network.nios.members[0].name }
    }
    check = {
      "nios.server_association_type" = "FAILOVER"
      "nios.failover_association"    = "example_failover_association1"
    }
  }

  step {
    nios {
      start_addr              = "190.0.0.121"
      end_addr                = "190.0.0.122"
      network_view            = infoblox_network.test_network.nios.network_view
      server_association_type = "MEMBER"
      failover_association    = "example_failover_association1"
      member                  = { name = infoblox_network.test_network.nios.members[0].name }
    }
    check = {
      "nios.server_association_type" = "MEMBER"
      "nios.member.name"             = "{{grid_member_hostname}}"
    }
  }

}

case "start_addr" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr   = "10.0.0.213"
      end_addr     = "10.0.0.215"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.start_addr" = "10.0.0.213"
    }
  }

  step {
    nios {
      start_addr   = "10.0.0.212"
      end_addr     = "10.0.0.215"
      network_view = infoblox_network.test_network.nios.network_view
    }
    check = {
      "nios.start_addr" = "10.0.0.212"
    }
  }

}

case "subscribe_settings" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "110.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr         = "110.0.0.127"
      end_addr           = "110.0.0.128"
      network_view       = infoblox_network.test_network.nios.network_view
      subscribe_settings = { enabled_attributes = ["DOMAINNAME"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "DOMAINNAME"
    }
  }

  step {
    nios {
      start_addr         = "110.0.0.127"
      end_addr           = "110.0.0.128"
      network_view       = infoblox_network.test_network.nios.network_view
      subscribe_settings = { enabled_attributes = ["ENDPOINT_PROFILE"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "ENDPOINT_PROFILE"
    }
  }

}

case "unknown_clients" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr      = "10.0.0.129"
      end_addr        = "10.0.0.130"
      network_view    = infoblox_network.test_network.nios.network_view
      unknown_clients = "Allow"
    }
    check = {
      "nios.unknown_clients" = "Allow"
    }
  }

  step {
    nios {
      start_addr      = "10.0.0.129"
      end_addr        = "10.0.0.130"
      network_view    = infoblox_network.test_network.nios.network_view
      unknown_clients = "Deny"
    }
    check = {
      "nios.unknown_clients" = "Deny"
    }
  }

}

case "update_dns_on_lease_renewal" {
  backend  = "nios"
  parallel = true

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "10.0.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      start_addr                  = "10.0.0.131"
      end_addr                    = "10.0.0.132"
      network_view                = infoblox_network.test_network.nios.network_view
      update_dns_on_lease_renewal = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
    }
  }

  step {
    nios {
      start_addr                  = "10.0.0.131"
      end_addr                    = "10.0.0.132"
      network_view                = infoblox_network.test_network.nios.network_view
      update_dns_on_lease_renewal = false
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "false"
    }
  }

}
