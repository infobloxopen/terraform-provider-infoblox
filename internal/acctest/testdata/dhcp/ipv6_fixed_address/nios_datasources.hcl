# Auto-generated datasource acceptance-test cases for Ipv6fixedaddress.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      ipv6addr     = "nios.ipv6addr"
      mac_address  = "nios.mac_address"
      match_client = "nios.match_client"
    }
  }

  pair_checks = ["nios.address_type", "nios.allow_telnet", "nios.comment", "nios.device_description", "nios.device_location", "nios.device_type", "nios.device_vendor", "nios.disable", "nios.disable_discovery", "nios.domain_name", "nios.duid", "nios.enable_immediate_discovery", "nios.ipv6addr", "nios.ipv6prefix", "nios.ipv6prefix_bits", "nios.mac_address", "nios.match_client", "nios.name", "nios.network", "nios.network_view", "nios.preferred_lifetime", "nios.reserved_interface", "nios.template", "nios.use_cli_credentials", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_logic_filter_rules", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_snmp3_credential", "nios.use_snmp_credential", "nios.use_valid_lifetime", "nios.valid_lifetime"]

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      match_client = "MAC_ADDRESS"
      mac_address  = "00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test_ipv6_network" {
    nios = {
      network = "2001:db8:{{random_hextet}}:{{random_int}}::/64"
      network_view = infoblox_network_view.parent_network_view.nios.name
    }
  }
  resource "infoblox_network_view" "parent_network_view" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.address_type", "nios.allow_telnet", "nios.comment", "nios.device_description", "nios.device_location", "nios.device_type", "nios.device_vendor", "nios.disable", "nios.disable_discovery", "nios.domain_name", "nios.duid", "nios.enable_immediate_discovery", "nios.ipv6addr", "nios.ipv6prefix", "nios.ipv6prefix_bits", "nios.mac_address", "nios.match_client", "nios.name", "nios.network", "nios.network_view", "nios.preferred_lifetime", "nios.reserved_interface", "nios.template", "nios.use_cli_credentials", "nios.use_domain_name", "nios.use_domain_name_servers", "nios.use_logic_filter_rules", "nios.use_options", "nios.use_preferred_lifetime", "nios.use_snmp3_credential", "nios.use_snmp_credential", "nios.use_valid_lifetime", "nios.valid_lifetime"]

  step {
    nios {
      ipv6addr     = "2001:db8:{{random_hextet}}:{{random_int}}::1"
      match_client = "MAC_ADDRESS"
      mac_address  = "00:0c:29:ab:cd:ef"
      network      = infoblox_ipv6_network.test_ipv6_network.nios.network
      network_view = infoblox_network_view.parent_network_view.nios.name
      ext_attrs    = { Site = "{{random}}" }
    }
  }

}
