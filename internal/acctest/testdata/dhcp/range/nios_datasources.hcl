# Auto-generated datasource acceptance-test cases for Range.
case "filters" {
  backend           = "nios"
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

  filter {
    type = "filters"
    values = {
      start_addr   = "nios.start_addr"
      network_view = "nios.network_view"
    }
  }

  pair_checks = ["nios.always_update_dns", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.deny_all_clients", "nios.deny_bootp", "nios.disable", "nios.discovery_member", "nios.enable_ddns", "nios.enable_dhcp_thresholds", "nios.enable_discovery", "nios.enable_email_warnings", "nios.enable_ifmap_publishing", "nios.enable_pxe_lease_time", "nios.enable_snmp_warnings", "nios.end_addr", "nios.failover_association", "nios.high_water_mark", "nios.high_water_mark_reset", "nios.ignore_dhcp_option_list_request", "nios.ignore_id", "nios.known_clients", "nios.lease_scavenge_time", "nios.low_water_mark", "nios.low_water_mark_reset", "nios.name", "nios.network", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.recycle_leases", "nios.restart_if_needed", "nios.same_port_control_discovery_blackout", "nios.server_association_type", "nios.split_scope_exclusion_percent", "nios.start_addr", "nios.template", "nios.unknown_clients", "nios.update_dns_on_lease_renewal", "nios.use_blackout_setting", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_deny_bootp", "nios.use_discovery_basic_polling_settings", "nios.use_email_list", "nios.use_enable_ddns", "nios.use_enable_dhcp_thresholds", "nios.use_enable_discovery", "nios.use_enable_ifmap_publishing", "nios.use_ignore_dhcp_option_list_request", "nios.use_ignore_id", "nios.use_known_clients", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_ms_options", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_recycle_leases", "nios.use_subscribe_settings", "nios.use_unknown_clients", "nios.use_update_dns_on_lease_renewal"]

  step {
    nios {
      start_addr   = "10.0.0.10"
      end_addr     = "10.0.0.20"
      network_view = infoblox_network.test_network.nios.network_view
    }
  }

}

case "ext_attr_filters" {
  backend           = "nios"
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

  filter {
    type = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.always_update_dns", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.deny_all_clients", "nios.deny_bootp", "nios.disable", "nios.discovery_member", "nios.enable_ddns", "nios.enable_dhcp_thresholds", "nios.enable_discovery", "nios.enable_email_warnings", "nios.enable_ifmap_publishing", "nios.enable_pxe_lease_time", "nios.enable_snmp_warnings", "nios.end_addr", "nios.failover_association", "nios.high_water_mark", "nios.high_water_mark_reset", "nios.ignore_dhcp_option_list_request", "nios.ignore_id", "nios.known_clients", "nios.lease_scavenge_time", "nios.low_water_mark", "nios.low_water_mark_reset", "nios.name", "nios.network", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.recycle_leases", "nios.restart_if_needed", "nios.same_port_control_discovery_blackout", "nios.server_association_type", "nios.split_scope_exclusion_percent", "nios.start_addr", "nios.template", "nios.unknown_clients", "nios.update_dns_on_lease_renewal", "nios.use_blackout_setting", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_deny_bootp", "nios.use_discovery_basic_polling_settings", "nios.use_email_list", "nios.use_enable_ddns", "nios.use_enable_dhcp_thresholds", "nios.use_enable_discovery", "nios.use_enable_ifmap_publishing", "nios.use_ignore_dhcp_option_list_request", "nios.use_ignore_id", "nios.use_known_clients", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_ms_options", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_recycle_leases", "nios.use_subscribe_settings", "nios.use_unknown_clients", "nios.use_update_dns_on_lease_renewal"]

  step {
    nios {
      start_addr   = "10.0.0.10"
      end_addr     = "10.0.0.20"
      network_view = infoblox_network.test_network.nios.network_view
      ext_attrs    = { Site = "{{random}}" }
    }
  }

}
