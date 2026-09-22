# Auto-generated datasource acceptance-test cases for Sharednetwork.
case "filters" {
  backend           = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.91.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.authority", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.ddns_update_fixed_addresses", "nios.ddns_use_option81", "nios.deny_bootp", "nios.disable", "nios.enable_ddns", "nios.enable_pxe_lease_time", "nios.ignore_dhcp_option_list_request", "nios.ignore_id", "nios.lease_scavenge_time", "nios.name", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.update_dns_on_lease_renewal", "nios.use_authority", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_ddns_update_fixed_addresses", "nios.use_ddns_use_option81", "nios.use_deny_bootp", "nios.use_enable_ddns", "nios.use_ignore_client_identifier", "nios.use_ignore_dhcp_option_list_request", "nios.use_ignore_id", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_update_dns_on_lease_renewal"]

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }]
      network_view = infoblox_network_view.test_view.nios.name
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
  resource "infoblox_network" "test_network1" {
    nios = {
      network      = "201.93.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.authority", "nios.bootfile", "nios.bootserver", "nios.comment", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.ddns_update_fixed_addresses", "nios.ddns_use_option81", "nios.deny_bootp", "nios.disable", "nios.enable_ddns", "nios.enable_pxe_lease_time", "nios.ignore_dhcp_option_list_request", "nios.ignore_id", "nios.lease_scavenge_time", "nios.name", "nios.network_view", "nios.nextserver", "nios.pxe_lease_time", "nios.update_dns_on_lease_renewal", "nios.use_authority", "nios.use_bootfile", "nios.use_bootserver", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_ddns_update_fixed_addresses", "nios.use_ddns_use_option81", "nios.use_deny_bootp", "nios.use_enable_ddns", "nios.use_ignore_client_identifier", "nios.use_ignore_dhcp_option_list_request", "nios.use_ignore_id", "nios.use_lease_scavenge_time", "nios.use_logic_filter_rules", "nios.use_nextserver", "nios.use_options", "nios.use_pxe_lease_time", "nios.use_update_dns_on_lease_renewal"]

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network1.id }]
      network_view = infoblox_network_view.test_view.nios.name
      ext_attrs    = { Site = "{{random2}}" }
    }
  }

}
