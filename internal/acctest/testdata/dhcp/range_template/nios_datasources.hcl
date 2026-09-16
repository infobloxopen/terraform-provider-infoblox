# Auto-generated datasource acceptance-test cases for Rangetemplate.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.bootfile", "nios.bootserver", "nios.cloud_api_compatible", "nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.deny_all_clients", "nios.deny_bootp", "nios.enable_ddns", "nios.enable_dhcp_thresholds", "nios.enable_email_warnings", "nios.enable_pxe_lease_time", "nios.enable_snmp_warnings", "nios.failover_association", "nios.high_water_mark", "nios.high_water_mark_reset", "nios.ignore_dhcp_option_list_request", "nios.known_clients", "nios.lease_scavenge_time", "nios.low_water_mark", "nios.low_water_mark_reset", "nios.name", "nios.nextserver", "nios.number_of_addresses", "nios.offset", "nios.pxe_lease_time", "nios.recycle_leases", "nios.server_association_type", "nios.unknown_clients", "nios.update_dns_on_lease_renewal", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_deny_bootp", "nios.use_email_list", "nios.use_enable_ddns", "nios.use_enable_dhcp_thresholds", "nios.use_ignore_dhcp_option_list_request", "nios.use_known_clients", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_ms_options", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_recycle_leases", "nios.use_unknown_clients", "nios.use_update_dns_on_lease_renewal"]

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
    }
  }

}

case "ext_attr_filters" {
  backend     = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      "Tenant ID" = "nios.ext_attrs[\"Tenant ID\"]"
    }
  }

  pair_checks = ["nios.bootfile", "nios.bootserver", "nios.cloud_api_compatible", "nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.deny_all_clients", "nios.deny_bootp", "nios.enable_ddns", "nios.enable_dhcp_thresholds", "nios.enable_email_warnings", "nios.enable_pxe_lease_time", "nios.enable_snmp_warnings", "nios.failover_association", "nios.high_water_mark", "nios.high_water_mark_reset", "nios.ignore_dhcp_option_list_request", "nios.known_clients", "nios.lease_scavenge_time", "nios.low_water_mark", "nios.low_water_mark_reset", "nios.name", "nios.nextserver", "nios.number_of_addresses", "nios.offset", "nios.pxe_lease_time", "nios.recycle_leases", "nios.server_association_type", "nios.unknown_clients", "nios.update_dns_on_lease_renewal", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_deny_bootp", "nios.use_email_list", "nios.use_enable_ddns", "nios.use_enable_dhcp_thresholds", "nios.use_ignore_dhcp_option_list_request", "nios.use_known_clients", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_ms_options", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_recycle_leases", "nios.use_unknown_clients", "nios.use_update_dns_on_lease_renewal"]

  step {
    nios {
      name                 = "{{random}}"
      number_of_addresses  = 100
      offset               = 50
      cloud_api_compatible = true
      ext_attrs = { "Tenant ID" = "{{random2}}" }
    }
  }
}
