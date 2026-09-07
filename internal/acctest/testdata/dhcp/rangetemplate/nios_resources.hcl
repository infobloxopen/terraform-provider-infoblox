// Objects to be presnet in the grid to run the Tcs]
// mac_filter, mac_filter2 - IPv4 MAC Filter
// nac_filter - IPv4 NAC Filter
// example-option-filter-1, example-option-filter-2 - IPVv4 Option Filters
// relay_agent_filter - IPv4 Relay Agent Filter
//test_filter_fingerprint, test_filter_fingerprint1 - IPv4 Fingerprint Filter
// example_failover_association, example_failover_association1 - Failover Association

# Auto-generated resource acceptance-test cases for Rangetemplate.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.name"                            = "{{random}}"
      "nios.cloud_api_compatible"            = "true"
      "nios.ddns_generate_hostname"          = "false"
      "nios.deny_all_clients"                = "false"
      "nios.deny_bootp"                      = "false"
      "nios.enable_ddns"                     = "false"
      "nios.enable_dhcp_thresholds"          = "false"
      "nios.enable_email_warnings"           = "false"
      "nios.enable_pxe_lease_time"           = "false"
      "nios.enable_snmp_warnings"            = "false"
      "nios.high_water_mark"                 = "95"
      "nios.high_water_mark_reset"           = "85"
      "nios.ignore_dhcp_option_list_request" = "false"
      "nios.lease_scavenge_time"             = "-1"
      "nios.low_water_mark"                  = "0"
      "nios.low_water_mark_reset"            = "10"
      "nios.recycle_leases"                  = "true"
      "nios.server_association_type"         = "NONE"
      "nios.update_dns_on_lease_renewal"     = "false"
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
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
  }

}

case "bootfile" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      bootfile             = "bootfile.txt"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.bootfile" = "bootfile.txt"
    }
  }

  step {
    nios {
      bootfile             = "bootfile12.txt"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.bootfile" = "bootfile12.txt"
    }
  }

}

case "bootserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      bootserver           = "bootserver"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.bootserver" = "bootserver"
    }
  }

  step {
    nios {
      bootserver           = "bootserver3"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.bootserver" = "bootserver3"
    }
  }

}

case "cloud_api_compatible" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping this test as it is a known issue."
  parallel    = true

  step {
    nios {
      cloud_api_compatible = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
    }
    check = {
      "nios.cloud_api_compatible" = "true"
    }
  }

  step {
    nios {
      cloud_api_compatible = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
    }
    check = {
      "nios.cloud_api_compatible" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      comment              = "comment for range template"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.comment" = "comment for range template"
    }
  }

  step {
    nios {
      comment              = "comment for range template updated"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.comment" = "comment for range template updated"
    }
  }

}

case "ddns_domainname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ddns_domainname      = "aa.bb.com"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.ddns_domainname" = "aa.bb.com"
    }
  }

  step {
    nios {
      ddns_domainname      = "qq.ww.com"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.ddns_domainname" = "qq.ww.com"
    }
  }

}

case "ddns_generate_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ddns_generate_hostname = true
      name                   = "{{random}}"
      number_of_addresses    = 100
      offset                 = 50
      cloud_api_compatible   = true
    }
    check = {
      "nios.ddns_generate_hostname" = "true"
    }
  }

  step {
    nios {
      ddns_generate_hostname = false
      name                   = "{{random}}"
      number_of_addresses    = 100
      offset                 = 50
      cloud_api_compatible   = true
    }
    check = {
      "nios.ddns_generate_hostname" = "false"
    }
  }

}

case "delegated_member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      delegated_member     = { name = "{{grid_master_hostname}}" }
    }
    check = {
      "nios.delegated_member.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      delegated_member     = { name = "{{grid_member_hostname}}" }
    }
    check = {
      "nios.delegated_member.name" = "{{grid_member_hostname}}"
    }
  }

}

case "deny_all_clients" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      deny_all_clients     = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.deny_all_clients" = "true"
    }
  }

  step {
    nios {
      deny_all_clients     = false
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.deny_all_clients" = "false"
    }
  }

}

case "deny_bootp" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      deny_bootp           = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.deny_bootp" = "true"
    }
  }

  step {
    nios {
      deny_bootp           = false
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.deny_bootp" = "false"
    }
  }

}

case "email_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      email_list           = ["bbb@info.com", "aaa@wapi.com"]
    }
    check = {
      "nios.email_list.0" = "bbb@info.com"
      "nios.email_list.1" = "aaa@wapi.com"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      email_list           = ["abc@info.com", "xyz@wapi.com"]
    }
    check = {
      "nios.email_list.0" = "abc@info.com"
      "nios.email_list.1" = "xyz@wapi.com"
    }
  }

}

case "enable_ddns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enable_ddns          = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.enable_ddns" = "true"
    }
  }

  step {
    nios {
      enable_ddns          = false
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.enable_ddns" = "false"
    }
  }

}

case "enable_dhcp_thresholds" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enable_dhcp_thresholds = true
      name                   = "{{random}}"
      number_of_addresses    = 100
      offset                 = 50
      cloud_api_compatible   = true
    }
    check = {
      "nios.enable_dhcp_thresholds" = "true"
    }
  }

  step {
    nios {
      enable_dhcp_thresholds = false
      name                   = "{{random}}"
      number_of_addresses    = 100
      offset                 = 50
      cloud_api_compatible   = true
    }
    check = {
      "nios.enable_dhcp_thresholds" = "false"
    }
  }

}

case "enable_email_warnings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enable_email_warnings = true
      name                  = "{{random}}"
      number_of_addresses   = 100
      offset                = 50
      cloud_api_compatible  = true
    }
    check = {
      "nios.enable_email_warnings" = "true"
    }
  }

  step {
    nios {
      enable_email_warnings = false
      name                  = "{{random}}"
      number_of_addresses   = 100
      offset                = 50
      cloud_api_compatible  = true
    }
    check = {
      "nios.enable_email_warnings" = "false"
    }
  }

}

case "enable_pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enable_pxe_lease_time = true
      pxe_lease_time        = 72000
      name                  = "{{random}}"
      number_of_addresses   = 100
      offset                = 50
      cloud_api_compatible  = true
    }
    check = {
      "nios.enable_pxe_lease_time" = "true"
    }
  }

  step {
    nios {
      enable_pxe_lease_time = false
      pxe_lease_time        = 72000
      name                  = "{{random}}"
      number_of_addresses   = 100
      offset                = 50
      cloud_api_compatible  = true
    }
    check = {
      "nios.enable_pxe_lease_time" = "false"
    }
  }

}

case "enable_snmp_warnings" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enable_snmp_warnings = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.enable_snmp_warnings" = "true"
    }
  }

  step {
    nios {
      enable_snmp_warnings = false
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.enable_snmp_warnings" = "false"
    }
  }

}

case "exclude" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      exclude              = [{ number_of_addresses = 10, offset = 20 }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.exclude.#"                     = "1"
      "nios.exclude.0.number_of_addresses" = "10"
      "nios.exclude.0.offset"              = "20"
    }
  }

  step {
    nios {
      exclude              = [{ number_of_addresses = 15, offset = 25, comment = "exclude for range template" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.exclude.#"                     = "1"
      "nios.exclude.0.number_of_addresses" = "15"
      "nios.exclude.0.offset"              = "25"
      "nios.exclude.0.comment"             = "exclude for range template"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      ext_attrs            = { "Tenant ID" = "{{random2}}" }
      cloud_api_compatible = true
    }
    check = {
      "nios.ext_attrs.Tenant ID" = "{{random2}}"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      ext_attrs            = { "Tenant ID" = "{{random3}}" }
      cloud_api_compatible = true
    }
    check = {
      "nios.ext_attrs.Tenant ID" = "{{random3}}"
    }
  }

}

case "failover_association" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      server_association_type = "FAILOVER"
      failover_association    = "example_failover_association"
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
    }
    check = {
      "nios.failover_association" = "example_failover_association"
    }
  }

  step {
    nios {
      server_association_type = "FAILOVER"
      failover_association    = "example_failover_association1"
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
    }
    check = {
      "nios.failover_association" = "example_failover_association1"
    }
  }

}

case "fingerprint_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      fingerprint_filter_rules = [{ filter = "test_filter_fingerprint", permission = "Allow" }]
      name                     = "{{random}}"
      number_of_addresses      = 100
      offset                   = 50
      cloud_api_compatible     = true
    }
    check = {
      "nios.fingerprint_filter_rules.#"            = "1"
      "nios.fingerprint_filter_rules.0.filter"     = "test_filter_fingerprint"
      "nios.fingerprint_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      fingerprint_filter_rules = [{ filter = "test_filter_fingerprint1", permission = "Deny" }]
      name                     = "{{random}}"
      number_of_addresses      = 100
      offset                   = 50
      cloud_api_compatible     = true
    }
    check = {
      "nios.fingerprint_filter_rules.#"            = "1"
      "nios.fingerprint_filter_rules.0.filter"     = "test_filter_fingerprint1"
      "nios.fingerprint_filter_rules.0.permission" = "Deny"
    }
  }

}

case "high_water_mark" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      high_water_mark      = 55
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.high_water_mark" = "55"
    }
  }

  step {
    nios {
      high_water_mark      = 70
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.high_water_mark" = "70"
    }
  }

}

case "high_water_mark_reset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      high_water_mark_reset = 10
      name                  = "{{random}}"
      number_of_addresses   = 100
      offset                = 50
      cloud_api_compatible  = true
    }
    check = {
      "nios.high_water_mark_reset" = "10"
    }
  }

  step {
    nios {
      high_water_mark_reset = 20
      name                  = "{{random}}"
      number_of_addresses   = 100
      offset                = 50
      cloud_api_compatible  = true
    }
    check = {
      "nios.high_water_mark_reset" = "20"
    }
  }

}

case "ignore_dhcp_option_list_request" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      ignore_dhcp_option_list_request = true
      name                            = "{{random}}"
      number_of_addresses             = 100
      offset                          = 50
      cloud_api_compatible            = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "true"
    }
  }

  step {
    nios {
      ignore_dhcp_option_list_request = false
      name                            = "{{random}}"
      number_of_addresses             = 100
      offset                          = 50
      cloud_api_compatible            = true
    }
    check = {
      "nios.ignore_dhcp_option_list_request" = "false"
    }
  }

}

case "known_clients" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      known_clients        = "Allow"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.known_clients" = "Allow"
    }
  }

  step {
    nios {
      known_clients        = "Deny"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.known_clients" = "Deny"
    }
  }

}

case "lease_scavenge_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      lease_scavenge_time  = 86700
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.lease_scavenge_time" = "86700"
    }
  }

  step {
    nios {
      lease_scavenge_time  = 98400
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.lease_scavenge_time" = "98400"
    }
  }

}

case "logic_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      logic_filter_rules   = [{ filter = "example-option-filter-1", type = "Option" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "example-option-filter-1"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

  step {
    nios {
      logic_filter_rules   = [{ filter = "example-option-filter-2", type = "Option" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.logic_filter_rules.#"        = "1"
      "nios.logic_filter_rules.0.filter" = "example-option-filter-2"
      "nios.logic_filter_rules.0.type"   = "Option"
    }
  }

}

case "low_water_mark" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      low_water_mark       = 71
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.low_water_mark" = "71"
    }
  }

  step {
    nios {
      low_water_mark       = 33
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.low_water_mark" = "33"
    }
  }

}

case "low_water_mark_reset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      low_water_mark_reset = 36
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.low_water_mark_reset" = "36"
    }
  }

  step {
    nios {
      low_water_mark_reset = 14
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.low_water_mark_reset" = "14"
    }
  }

}

case "mac_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      mac_filter_rules     = [{ filter = "mac_filter", permission = "Allow" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.mac_filter_rules.#"            = "1"
      "nios.mac_filter_rules.0.filter"     = "mac_filter"
      "nios.mac_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      mac_filter_rules     = [{ filter = "mac_filter2", permission = "Deny" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.mac_filter_rules.#"            = "1"
      "nios.mac_filter_rules.0.filter"     = "mac_filter2"
      "nios.mac_filter_rules.0.permission" = "Deny"
    }
  }

}

case "member" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      member               = { name = "{{grid_master_hostname}}" }
    }
    check = {
      "nios.member.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      member               = { name = "{{grid_member_hostname}}" }
    }
    check = {
      "nios.member.name" = "{{grid_member_hostname}}"
    }
  }

}

case "ms_options" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping this test as it requires MS_Server setup."
  parallel    = true

  step {
    nios {
      ms_options           = ""
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.ms_options" = "MS_OPTIONS_REPLACE_ME"
    }
  }

  step {
    nios {
      ms_options           = ""
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.ms_options" = "MS_OPTIONS_UPDATE_REPLACE_ME"
    }
  }

}

case "ms_server" {
  backend     = "nios"
  skip        = true
  skip_reason = "t.Skip: Skipping this test as it requires MS_Server setup."
  parallel    = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.ms_server.ipv4addr" = "10.120.23.22"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.ms_server.ipv4addr" = "10.120.23.23"
    }
  }

}

case "nac_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      nac_filter_rules     = [{ filter = "nac_filter", permission = "Allow" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.nac_filter_rules.#"            = "1"
      "nios.nac_filter_rules.0.filter"     = "nac_filter"
      "nios.nac_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      nac_filter_rules     = [{ filter = "nac_filter", permission = "Deny" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.nac_filter_rules.#"            = "1"
      "nios.nac_filter_rules.0.filter"     = "nac_filter"
      "nios.nac_filter_rules.0.permission" = "Deny"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                 = "{{random2}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "nextserver" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      nextserver           = "next-server-1"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.nextserver" = "next-server-1"
    }
  }

  step {
    nios {
      nextserver           = "next-server-2"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.nextserver" = "next-server-2"
    }
  }

}

case "number_of_addresses" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.number_of_addresses" = "100"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 500
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.number_of_addresses" = "500"
    }
  }

}

case "offset" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.offset" = "50"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 2000
      cloud_api_compatible = true
    }
    check = {
      "nios.offset" = "2000"
    }
  }

}

case "option_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      option_filter_rules  = [{ filter = "example-option-filter-1", permission = "Allow" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.option_filter_rules.#"            = "1"
      "nios.option_filter_rules.0.filter"     = "example-option-filter-1"
      "nios.option_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      option_filter_rules  = [{ filter = "example-option-filter-2", permission = "Deny" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.option_filter_rules.#"            = "1"
      "nios.option_filter_rules.0.filter"     = "example-option-filter-2"
      "nios.option_filter_rules.0.permission" = "Deny"
    }
  }

}

case "options" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      options              = [{ name = "domain-name", num = "15", value = "aa.bb.com" }, { name = "dhcp-lease-time", num = "51", value = "72000" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.options.#"       = "2"
      "nios.options.0.name"  = "domain-name"
      "nios.options.0.value" = "aa.bb.com"
      "nios.options.1.name"  = "dhcp-lease-time"
      "nios.options.1.value" = "72000"
    }
  }

  step {
    nios {
      options              = [{ name = "domain-name", num = "15", value = "cc.dd.com" }, { name = "dhcp-lease-time", num = "51", value = "82000" }]
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.options.#"       = "2"
      "nios.options.0.name"  = "domain-name"
      "nios.options.0.value" = "cc.dd.com"
      "nios.options.1.name"  = "dhcp-lease-time"
      "nios.options.1.value" = "82000"
    }
  }

}

case "pxe_lease_time" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      pxe_lease_time       = 3600
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.pxe_lease_time" = "3600"
    }
  }

  step {
    nios {
      pxe_lease_time       = 7200
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.pxe_lease_time" = "7200"
    }
  }

}

case "recycle_leases" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      recycle_leases       = true
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.recycle_leases" = "true"
    }
  }

  step {
    nios {
      recycle_leases       = false
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.recycle_leases" = "false"
    }
  }

}

case "relay_agent_filter_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      relay_agent_filter_rules = [{ filter = "relay_agent_filter", permission = "Allow" }]
      name                     = "{{random}}"
      number_of_addresses      = 100
      offset                   = 50
      cloud_api_compatible     = true
    }
    check = {
      "nios.relay_agent_filter_rules.#"            = "1"
      "nios.relay_agent_filter_rules.0.filter"     = "relay_agent_filter"
      "nios.relay_agent_filter_rules.0.permission" = "Allow"
    }
  }

  step {
    nios {
      relay_agent_filter_rules = [{ filter = "relay_agent_filter", permission = "Deny" }]
      name                     = "{{random}}"
      number_of_addresses      = 100
      offset                   = 50
      cloud_api_compatible     = true
    }
    check = {
      "nios.relay_agent_filter_rules.#"            = "1"
      "nios.relay_agent_filter_rules.0.filter"     = "relay_agent_filter"
      "nios.relay_agent_filter_rules.0.permission" = "Deny"
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

  step {
    nios {
      server_association_type = "FAILOVER"
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
      failover_association = "example_failover_association"
    }
    check = {
      "nios.server_association_type" = "FAILOVER"
    }
  }

  step {
    nios {
      server_association_type = "NONE"
      name                    = "{{random}}"
      number_of_addresses     = 100
      offset                  = 50
      cloud_api_compatible    = true
    }
    check = {
      "nios.server_association_type" = "NONE"
    }
  }

}

case "unknown_clients" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      unknown_clients      = "Deny"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.unknown_clients" = "Deny"
    }
  }

  step {
    nios {
      unknown_clients      = "Allow"
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
    check = {
      "nios.unknown_clients" = "Allow"
    }
  }

}

case "update_dns_on_lease_renewal" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      update_dns_on_lease_renewal = true
      name                        = "{{random}}"
      number_of_addresses         = 100
      offset                      = 50
      cloud_api_compatible        = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "true"
    }
  }

  step {
    nios {
      update_dns_on_lease_renewal = false
      name                        = "{{random}}"
      number_of_addresses         = 100
      offset                      = 50
      cloud_api_compatible        = true
    }
    check = {
      "nios.update_dns_on_lease_renewal" = "false"
    }
  }

}
