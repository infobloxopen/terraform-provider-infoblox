# Auto-generated datasource acceptance-test cases for Networkcontainer.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      network = "nios.network"
    }
  }

  pair_checks = ["nios.authority", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.ddns_update_fixed_addresses", "nios.ddns_use_option81", "nios.delete_reason", "nios.deny_bootp", "nios.discovery_member", "nios.enable_ddns", "nios.enable_dhcp_thresholds", "nios.enable_discovery", "nios.enable_email_warnings", "nios.enable_immediate_discovery", "nios.enable_pxe_lease_time", "nios.enable_snmp_warnings", "nios.high_water_mark", "nios.high_water_mark_reset", "nios.ignore_dhcp_option_list_request", "nios.ignore_id", "nios.lease_scavenge_time", "nios.low_water_mark", "nios.low_water_mark_reset", "nios.mgm_private", "nios.network", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.recycle_leases", "nios.rir_organization", "nios.rir_registration_action", "nios.rir_registration_status", "nios.same_port_control_discovery_blackout", "nios.send_rir_request", "nios.unmanaged", "nios.update_dns_on_lease_renewal", "nios.use_authority", "nios.use_blackout_setting", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_ddns_update_fixed_addresses", "nios.use_ddns_use_option81", "nios.use_deny_bootp", "nios.use_discovery_basic_polling_settings", "nios.use_email_list", "nios.use_enable_ddns", "nios.use_enable_dhcp_thresholds", "nios.use_enable_discovery", "nios.use_ignore_dhcp_option_list_request", "nios.use_ignore_id", "nios.use_ipam_email_addresses", "nios.use_ipam_threshold_settings", "nios.use_ipam_trap_settings", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_mgm_private", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_recycle_leases", "nios.use_subscribe_settings", "nios.use_update_dns_on_lease_renewal", "nios.use_zone_associations"]

  step {
    nios {
      network = "{{random_cidr_network}}"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.authority", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.ddns_update_fixed_addresses", "nios.ddns_use_option81", "nios.delete_reason", "nios.deny_bootp", "nios.discovery_member", "nios.enable_ddns", "nios.enable_dhcp_thresholds", "nios.enable_discovery", "nios.enable_email_warnings", "nios.enable_immediate_discovery", "nios.enable_pxe_lease_time", "nios.enable_snmp_warnings", "nios.high_water_mark", "nios.high_water_mark_reset", "nios.ignore_dhcp_option_list_request", "nios.ignore_id", "nios.lease_scavenge_time", "nios.low_water_mark", "nios.low_water_mark_reset", "nios.mgm_private", "nios.network", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.recycle_leases", "nios.rir_organization", "nios.rir_registration_action", "nios.rir_registration_status", "nios.same_port_control_discovery_blackout", "nios.send_rir_request", "nios.unmanaged", "nios.update_dns_on_lease_renewal", "nios.use_authority", "nios.use_blackout_setting", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_ddns_update_fixed_addresses", "nios.use_ddns_use_option81", "nios.use_deny_bootp", "nios.use_discovery_basic_polling_settings", "nios.use_email_list", "nios.use_enable_ddns", "nios.use_enable_dhcp_thresholds", "nios.use_enable_discovery", "nios.use_ignore_dhcp_option_list_request", "nios.use_ignore_id", "nios.use_ipam_email_addresses", "nios.use_ipam_threshold_settings", "nios.use_ipam_trap_settings", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_mgm_private", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_recycle_leases", "nios.use_subscribe_settings", "nios.use_update_dns_on_lease_renewal", "nios.use_zone_associations"]

  step {
    nios {
      network   = "{{random_cidr_network}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
