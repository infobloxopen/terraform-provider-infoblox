# Auto-generated resource acceptance-test cases for Ipv6range.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.network"                              = "14::/64"
      "nios.address_type"                         = "ADDRESS"
      "nios.disable"                              = "false"
      "nios.enable_discovery"                     = "false"
      "nios.recycle_leases"                       = "true"
      "nios.restart_if_needed"                    = "false"
      "nios.same_port_control_discovery_blackout" = "false"
      "nios.server_association_type"              = "NONE"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::20"
      end_addr     = "14::30"
      network_view = infoblox_network_view.test.nios.name
    }
  }

}

case "import" {
  backend  = "nios"
  parallel = true
  import   = true
  import_ignore = ["nios.ext_attrs_all", "nios.restart_if_needed"]
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::40"
      end_addr     = "14::50"
      network_view = infoblox_network_view.test.nios.name
    }
  }

}

case "address_type" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      address_type = "ADDRESS"
    }
    check = {
      "nios.address_type" = "ADDRESS"
    }
  }

  step {
    nios {
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1000::"
      ipv6_end_prefix   = "14:0:0:1fff::"
      network_view      = infoblox_network_view.test.nios.name
      address_type      = "PREFIX"
      ipv6_prefix_bits  = 80
    }
    check = {
      "nios.address_type" = "PREFIX"
    }
  }

  step {
    nios {
      network           = infoblox_ipv6_network.test.nios.network
      start_addr        = "14::1"
      end_addr          = "14::10"
      ipv6_start_prefix = "14:0:0:2000::"
      ipv6_end_prefix   = "14:0:0:2fff::"
      network_view      = infoblox_network_view.test.nios.name
      address_type      = "BOTH"
      ipv6_prefix_bits  = 80
    }
    check = {
      "nios.address_type" = "BOTH"
    }
  }

}

case "cloud_info" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.cloud_info.authority_type"   = "GM"
      "nios.cloud_info.delegated_scope"  = "NONE"
      "nios.cloud_info.owned_by_adaptor" = "false"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      cloud_info   = { delegated_member = null }
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.cloud_info.#" = "0"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      comment      = "Comment for the Ipv6range object"
    }
    check = {
      "nios.comment" = "Comment for the Ipv6range object"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      comment      = "Updated comment for the Ipv6range object"
    }
    check = {
      "nios.comment" = "Updated comment for the Ipv6range object"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      disable      = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
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
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network                       = infoblox_ipv6_network.test.nios.network
      start_addr                    = "14::1"
      end_addr                      = "14::10"
      network_view                  = infoblox_network_view.test.nios.name
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
      network                       = infoblox_ipv6_network.test.nios.network
      start_addr                    = "14::1"
      end_addr                      = "14::10"
      network_view                  = infoblox_network_view.test.nios.name
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
      network                       = infoblox_ipv6_network.test.nios.network
      start_addr                    = "14::1"
      end_addr                      = "14::10"
      network_view                  = infoblox_network_view.test.nios.name
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
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network                    = infoblox_ipv6_network.test.nios.network
      start_addr                 = "14::1"
      end_addr                   = "14::10"
      discovery_blackout_setting = { enable_blackout = true, blackout_duration = 100, blackout_schedule = { weekdays = ["TUESDAY", "MONDAY", "FRIDAY"], frequency = "WEEKLY", every = 15, minutes_past_hour = 6, disable = false, repeat = "RECUR", hour_of_day = 20 } }
      network_view               = infoblox_network_view.test.nios.name
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
      network                    = infoblox_ipv6_network.test.nios.network
      start_addr                 = "14::1"
      end_addr                   = "14::10"
      discovery_blackout_setting = { enable_blackout = true, blackout_duration = 200, blackout_schedule = { minutes_past_hour = 6, repeat = "ONCE", day_of_month = 30, month = 1, year = 2026, hour_of_day = 20 } }
      network_view               = infoblox_network_view.test.nios.name
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
  skip_if_env_empty = ["NIOS_DISCOVERY_MEMBER_HOSTNAME"]
  skip_reason       = "NIOS_DISCOVERY_MEMBER_HOSTNAME environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network          = infoblox_ipv6_network.test.nios.network
      start_addr       = "14::1"
      end_addr         = "14::10"
      network_view     = infoblox_network_view.test.nios.name
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.discovery_member" = "{{discovery_member_hostname}}"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
  }

}

case "enable_discovery" {
  backend  = "nios"
  parallel = true
  skip_if_env_empty = ["NIOS_DISCOVERY_MEMBER_HOSTNAME"]
  skip_reason       = "NIOS_DISCOVERY_MEMBER_HOSTNAME environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "15::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network          = infoblox_ipv6_network.test2.nios.network
      start_addr       = "15::1"
      end_addr         = "15::10"
      network_view     = infoblox_network_view.test.nios.name
      enable_discovery = true
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.enable_discovery" = "true"
    }
  }

  step {
    nios {
      network          = infoblox_ipv6_network.test2.nios.network
      start_addr       = "15::1"
      end_addr         = "15::10"
      network_view     = infoblox_network_view.test.nios.name
      enable_discovery = false
      discovery_member = "{{discovery_member_hostname}}"
    }
    check = {
      "nios.enable_discovery" = "false"
    }
  }

}

case "end_addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.end_addr" = "14::10"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::20"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.end_addr" = "14::20"
    }
  }

}

case "enable_immediate_discovery" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network                    = infoblox_ipv6_network.test.nios.network
      start_addr                 = "14::101"
      end_addr                   = "14::110"
      network_view               = infoblox_network_view.test.nios.name
      enable_immediate_discovery = true
    }
    check = {
      "nios.enable_immediate_discovery" = "true"
    }
  }

  step {
    nios {
      network                    = infoblox_ipv6_network.test.nios.network
      start_addr                 = "14::101"
      end_addr                   = "14::110"
      network_view               = infoblox_network_view.test.nios.name
      enable_immediate_discovery = false
    }
    check = {
      "nios.enable_immediate_discovery" = "false"
    }
  }

}

case "exclude" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      exclude      = [{ start_address = "14::3", end_address = "14::5", comment = "Exclude range 14::3 - 14::5" }]
    }
    check = {
      "nios.exclude.#"               = "1"
      "nios.exclude.0.start_address" = "14::3"
      "nios.exclude.0.end_address"   = "14::5"
      "nios.exclude.0.comment"       = "Exclude range 14::3 - 14::5"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      exclude      = [{ start_address = "14::4", end_address = "14::6", comment = "Updated exclude range 14::4 - 14::6" }]
    }
    check = {
      "nios.exclude.#"               = "1"
      "nios.exclude.0.start_address" = "14::4"
      "nios.exclude.0.end_address"   = "14::6"
      "nios.exclude.0.comment"       = "Updated exclude range 14::4 - 14::6"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      ext_attrs    = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      ext_attrs    = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ipv6_end_prefix" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      address_type      = "PREFIX"
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1000::"
      ipv6_end_prefix   = "14:0:0:1fff::"
      network_view      = infoblox_network_view.test.nios.name
      ipv6_prefix_bits  = 80
    }
    check = {
      "nios.ipv6_end_prefix" = "14:0:0:1fff::"
    }
  }

  step {
    nios {
      address_type      = "PREFIX"
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1000::"
      ipv6_end_prefix   = "14:0:0:2000::"
      network_view      = infoblox_network_view.test.nios.name
      ipv6_prefix_bits  = 80
    }
    check = {
      "nios.ipv6_end_prefix" = "14:0:0:2000::"
    }
  }

}

case "ipv6_prefix_bits" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      address_type      = "PREFIX"
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1000::"
      ipv6_end_prefix   = "14:0:0:1fff::"
      network_view      = infoblox_network_view.test.nios.name
      ipv6_prefix_bits  = 80
    }
    check = {
      "nios.ipv6_prefix_bits" = "80"
    }
  }

  step {
    nios {
      address_type      = "PREFIX"
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1000::"
      ipv6_end_prefix   = "14:0:0:1fff::"
      network_view      = infoblox_network_view.test.nios.name
      ipv6_prefix_bits  = 81
    }
    check = {
      "nios.ipv6_prefix_bits" = "81"
    }
  }

}

case "ipv6_start_prefix" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      address_type      = "PREFIX"
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1000::"
      ipv6_end_prefix   = "14:0:0:1fff::"
      ipv6_prefix_bits  = 80
      network_view      = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.ipv6_start_prefix" = "14:0:0:1000::"
    }
  }

  step {
    nios {
      address_type      = "PREFIX"
      network           = infoblox_ipv6_network.test.nios.network
      ipv6_start_prefix = "14:0:0:1005::"
      ipv6_end_prefix   = "14:0:0:1fff::"
      ipv6_prefix_bits  = 80
      network_view      = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.ipv6_start_prefix" = "14:0:0:1005::"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_filteroption" "filter1" {
    nios = {
      name = "{{random}}-1"
    }
  }
  resource "infoblox_ipv6_filteroption" "filter2" {
    nios = {
      name = "{{random}}-2"
    }
  }
  PREREQ

  step {
    nios {
      network            = infoblox_ipv6_network.test.nios.network
      start_addr         = "14::1"
      end_addr           = "14::10"
      network_view       = infoblox_network_view.test.nios.name
      logic_filter_rules = [{ filter = infoblox_ipv6_filteroption.filter1.nios.name, type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "{{random}}-1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      network            = infoblox_ipv6_network.test.nios.network
      start_addr         = "14::1"
      end_addr           = "14::10"
      network_view       = infoblox_network_view.test.nios.name
      logic_filter_rules = [{ filter = infoblox_ipv6_filteroption.filter2.nios.name, type = "Option" }]
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "{{random}}-2"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "member" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "140::/64"
      network_view = "default"
      members = [
        { name = "{{grid_master_hostname}}" },
        { name = "{{grid_member_hostname}}" },
      ]
    }
  }
  PREREQ

  step {
    nios {
      network                 = infoblox_ipv6_network.test.nios.network
      start_addr              = "140::1"
      end_addr                = "140::10"
      network_view            = "default"
      member                  = { name = "{{grid_master_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.member.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      network                 = infoblox_ipv6_network.test.nios.network
      start_addr              = "140::1"
      end_addr                = "140::10"
      network_view            = "default"
      member                  = { name = "{{grid_member_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.member.name" = "{{grid_member_hostname}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      name         = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      name         = "{{random3}}"
    }
    check = {
      "nios.name" = "{{random3}}"
    }
  }

}

case "network" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "15::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test1.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.network" = "14::/64"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test2.nios.network
      start_addr   = "15::1"
      end_addr     = "15::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.network" = "15::/64"
    }
  }

}

case "network_view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.network_view" = "{{random}}"
    }
  }

}

case "option_filter_rules" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_filteroption" "filter1" {
    nios = {
      name = "{{random}}-1"
    }
  }
  resource "infoblox_ipv6_filteroption" "filter2" {
    nios = {
      name = "{{random}}-2"
    }
  }
  PREREQ

  step {
    nios {
      network             = infoblox_ipv6_network.test.nios.network
      start_addr          = "14::1"
      end_addr            = "14::10"
      network_view        = infoblox_network_view.test.nios.name
      option_filter_rules = [{ filter = infoblox_ipv6_filteroption.filter1.nios.name, permission = "Allow" }]
    }
    check = {
      "nios.option_filter_rules.#"            = "1"
      "nios.option_filter_rules.0.filter"     = "{{random}}-1"
      "nios.option_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      network             = infoblox_ipv6_network.test.nios.network
      start_addr          = "14::1"
      end_addr            = "14::10"
      network_view        = infoblox_network_view.test.nios.name
      option_filter_rules = [{ filter = infoblox_ipv6_filteroption.filter2.nios.name, permission = "Deny" }]
    }
    check = {
      "nios.option_filter_rules.#"            = "1"
      "nios.option_filter_rules.0.filter"     = "{{random}}-2"
      "nios.option_filter_rules.0.permission" = "Deny"
    }
  }

}

case "port_control_blackout_setting" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network                       = infoblox_ipv6_network.test.nios.network
      start_addr                    = "14::1"
      end_addr                      = "14::10"
      network_view                  = infoblox_network_view.test.nios.name
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
      network                       = infoblox_ipv6_network.test.nios.network
      start_addr                    = "14::1"
      end_addr                      = "14::10"
      network_view                  = infoblox_network_view.test.nios.name
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

case "recycle_leases" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network        = infoblox_ipv6_network.test.nios.network
      start_addr     = "14::1"
      end_addr       = "14::10"
      network_view   = infoblox_network_view.test.nios.name
      recycle_leases = true
    }
    check = {
      "nios.recycle_leases" = "true"
    }
  }

  step {
    nios {
      network        = infoblox_ipv6_network.test.nios.network
      start_addr     = "14::1"
      end_addr       = "14::10"
      network_view   = infoblox_network_view.test.nios.name
      recycle_leases = false
    }
    check = {
      "nios.recycle_leases" = "false"
    }
  }

}

case "restart_if_needed" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network           = infoblox_ipv6_network.test.nios.network
      start_addr        = "14::121"
      end_addr          = "14::130"
      network_view      = infoblox_network_view.test.nios.name
      restart_if_needed = true
    }
    check = {
      "nios.restart_if_needed" = "true"
    }
  }

  step {
    nios {
      network           = infoblox_ipv6_network.test.nios.network
      start_addr        = "14::121"
      end_addr          = "14::130"
      network_view      = infoblox_network_view.test.nios.name
      restart_if_needed = false
    }
    check = {
      "nios.restart_if_needed" = "false"
    }
  }

}

case "same_port_control_discovery_blackout" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network                              = infoblox_ipv6_network.test.nios.network
      start_addr                           = "14::1"
      end_addr                             = "14::10"
      network_view                         = infoblox_network_view.test.nios.name
      same_port_control_discovery_blackout = true
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "true"
    }
  }

  step {
    nios {
      network                              = infoblox_ipv6_network.test.nios.network
      start_addr                           = "14::1"
      end_addr                             = "14::10"
      network_view                         = infoblox_network_view.test.nios.name
      same_port_control_discovery_blackout = false
    }
    check = {
      "nios.same_port_control_discovery_blackout" = "false"
    }
  }

}

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "server_association_type" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "141::/64"
      network_view = "default"
      members = [{ name = "{{grid_member_hostname}}" }]
    }
  }
  PREREQ

  step {
    nios {
      network                 = infoblox_ipv6_network.test.nios.network
      start_addr              = "141::1"
      end_addr                = "141::10"
      network_view            = "default"
      member                  = { name = "{{grid_member_hostname}}" }
      server_association_type = "MEMBER"
    }
    check = {
      "nios.server_association_type" = "MEMBER"
    }
  }

  step {
    nios {
      network                 = infoblox_ipv6_network.test.nios.network
      start_addr              = "141::1"
      end_addr                = "141::10"
      network_view            = "default"
      server_association_type = "NONE"
    }
    check = {
      "nios.server_association_type" = "NONE"
    }
  }

}

case "start_addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.start_addr" = "14::1"
    }
  }

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::5"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
    check = {
      "nios.start_addr" = "14::5"
    }
  }

}

case "subscribe_settings" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "221::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  PREREQ

  step {
    nios {
      network            = infoblox_ipv6_network.test.nios.network
      start_addr         = "221::1"
      end_addr           = "221::10"
      network_view       = infoblox_ipv6_network.test.nios.network_view
      subscribe_settings = { enabled_attributes = ["DOMAINNAME", "ENDPOINT_PROFILE"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "DOMAINNAME"
      "nios.subscribe_settings.enabled_attributes.1" = "ENDPOINT_PROFILE"
    }
  }

  step {
    nios {
      network            = infoblox_ipv6_network.test.nios.network
      start_addr         = "221::1"
      end_addr           = "221::10"
      network_view       = infoblox_ipv6_network.test.nios.network_view
      subscribe_settings = { enabled_attributes = ["SECURITY_GROUP"] }
    }
    check = {
      "nios.subscribe_settings.enabled_attributes.0" = "SECURITY_GROUP"
    }
  }

}

case "template" {
  backend     = "nios"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network = "14::/64"
      network_view = infoblox_network_view.test.nios.name
    }
  }
  resource "infoblox_network_view" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_range_template" "test_template" {
    nios = {
      name = "{{random2}}"
      number_of_addresses = 10
      offset = 1
      cloud_api_compatible = true
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::131"
      end_addr     = "14::140"
      network_view = infoblox_network_view.test.nios.name
      template     = infoblox_ipv6_range_template.test_template.nios.name
    }
    check = {
      "nios.template" = "{{random2}}"
    }
  }

}
