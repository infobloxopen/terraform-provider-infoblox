# Auto-generated datasource acceptance-test cases for Ipv6network.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      network = "nios.network"
    }
  }

  pair_checks = ["nios.comment", "nios.ddns_domainname", "nios.ddns_enable_option_fqdn", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.disable", "nios.discovered_bridge_domain", "nios.discovered_tenant", "nios.discovery_member", "nios.domain_name", "nios.enable_ddns", "nios.enable_discovery", "nios.enable_ifmap_publishing", "nios.mgm_private", "nios.network", "nios.network_view", "nios.preferred_lifetime", "nios.recycle_leases", "nios.rir_organization", "nios.rir_registration_status", "nios.same_port_control_discovery_blackout", "nios.unmanaged", "nios.update_dns_on_lease_renewal", "nios.use_blackout_setting", "nios.use_ddns_domainname", "nios.use_ddns_enable_option_fqdn", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_discovery_basic_polling_settings", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_enable_ddns", "nios.use_enable_discovery", "nios.use_enable_ifmap_publishing", "nios.use_logic_filter_rules", "nios.use_mgm_private", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_recycle_leases", "nios.use_subscribe_settings", "nios.use_update_dns_on_lease_renewal", "nios.use_valid_lifetime", "nios.use_zone_associations", "nios.valid_lifetime"]

  step {
    nios {
      network = "{{random_ipv6_network}}"
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

  pair_checks = ["nios.comment", "nios.ddns_domainname", "nios.ddns_enable_option_fqdn", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.disable", "nios.discovered_bridge_domain", "nios.discovered_tenant", "nios.discovery_member", "nios.domain_name", "nios.enable_ddns", "nios.enable_discovery", "nios.enable_ifmap_publishing", "nios.mgm_private", "nios.network", "nios.network_view", "nios.preferred_lifetime", "nios.recycle_leases", "nios.rir_organization", "nios.rir_registration_status", "nios.same_port_control_discovery_blackout", "nios.unmanaged", "nios.update_dns_on_lease_renewal", "nios.use_blackout_setting", "nios.use_ddns_domainname", "nios.use_ddns_enable_option_fqdn", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_discovery_basic_polling_settings", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_enable_ddns", "nios.use_enable_discovery", "nios.use_enable_ifmap_publishing", "nios.use_logic_filter_rules", "nios.use_mgm_private", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_recycle_leases", "nios.use_subscribe_settings", "nios.use_update_dns_on_lease_renewal", "nios.use_valid_lifetime", "nios.use_zone_associations", "nios.valid_lifetime"]

  step {
    nios {
      network   = "{{random_ipv6_network}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
