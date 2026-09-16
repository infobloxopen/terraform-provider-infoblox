case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      ipv4addr = "nios.ipv4addr"
    }
  }

  pair_checks = ["nios.agent_circuit_id", "nios.agent_remote_id", "nios.allow_telnet", "nios.always_update_dns", "nios.bootfile", "nios.bootserver", "nios.client_identifier_prepend_zero", "nios.comment", "nios.ddns_domainname", "nios.ddns_hostname", "nios.deny_bootp", "nios.device_description", "nios.device_location", "nios.device_type", "nios.device_vendor", "nios.dhcp_client_identifier", "nios.disable", "nios.disable_discovery", "nios.enable_ddns", "nios.enable_immediate_discovery", "nios.enable_pxe_lease_time", "nios.ignore_dhcp_option_list_request", "nios.ipv4addr", "nios.mac", "nios.match_client", "nios.name", "nios.network", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.reserved_interface", "nios.restart_if_needed", "nios.template", "nios.use_bootfile", "nios.use_bootserver", "nios.use_cli_credentials", "nios.use_ddns_domainname", "nios.use_deny_bootp", "nios.use_enable_ddns", "nios.use_ignore_dhcp_option_list_request", "nios.use_logic_filter_rules", "nios.use_ms_options", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_snmp3_credential", "nios.use_snmp_credential"]

  step {
    nios {
      ipv4addr         = "15.0.0.241"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
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

  pair_checks = ["nios.agent_circuit_id", "nios.agent_remote_id", "nios.allow_telnet", "nios.always_update_dns", "nios.bootfile", "nios.bootserver", "nios.client_identifier_prepend_zero", "nios.comment", "nios.ddns_domainname", "nios.ddns_hostname", "nios.deny_bootp", "nios.device_description", "nios.device_location", "nios.device_type", "nios.device_vendor", "nios.dhcp_client_identifier", "nios.disable", "nios.disable_discovery", "nios.enable_ddns", "nios.enable_immediate_discovery", "nios.enable_pxe_lease_time", "nios.ignore_dhcp_option_list_request", "nios.ipv4addr", "nios.mac", "nios.match_client", "nios.name", "nios.network", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.reserved_interface", "nios.restart_if_needed", "nios.template", "nios.use_bootfile", "nios.use_bootserver", "nios.use_cli_credentials", "nios.use_ddns_domainname", "nios.use_deny_bootp", "nios.use_enable_ddns", "nios.use_ignore_dhcp_option_list_request", "nios.use_logic_filter_rules", "nios.use_ms_options", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_snmp3_credential", "nios.use_snmp_credential"]

  step {
    nios {
      ipv4addr         = "15.0.0.242"
      match_client     = "CIRCUIT_ID"
      agent_circuit_id = "{{random_int}}"
      ext_attrs        = { Site = "{{random}}" }
    }
  }

}


case "ms_server_struct" {
  backend     = "nios"
  skip        = true
  skip_reason = "Requires a PUT Implementation in the Data Source"
}
