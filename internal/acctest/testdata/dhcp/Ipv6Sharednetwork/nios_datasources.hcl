# Auto-generated datasource acceptance-test cases for Ipv6sharednetwork.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.ddns_use_option81", "nios.disable", "nios.domain_name", "nios.enable_ddns", "nios.name", "nios.network_view", "nios.preferred_lifetime", "nios.update_dns_on_lease_renewal", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_ddns_use_option81", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_enable_ddns", "nios.use_logic_filter_rules", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_update_dns_on_lease_renewal", "nios.use_valid_lifetime", "nios.valid_lifetime"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.ddns_domainname", "nios.ddns_generate_hostname", "nios.ddns_server_always_updates", "nios.ddns_ttl", "nios.ddns_use_option81", "nios.disable", "nios.domain_name", "nios.enable_ddns", "nios.name", "nios.network_view", "nios.preferred_lifetime", "nios.update_dns_on_lease_renewal", "nios.use_ddns_domainname", "nios.use_ddns_generate_hostname", "nios.use_ddns_ttl", "nios.use_ddns_use_option81", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_enable_ddns", "nios.use_logic_filter_rules", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_update_dns_on_lease_renewal", "nios.use_valid_lifetime", "nios.valid_lifetime"]

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random}}" }
    }
  }

}
