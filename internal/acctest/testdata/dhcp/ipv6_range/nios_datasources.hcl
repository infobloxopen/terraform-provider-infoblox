# Auto-generated datasource acceptance-test cases for Ipv6range.
case "filters" {
  backend = "nios"
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

  filter {
    type   = "filters"
    values = {
      network      = "nios.network"
      start_addr   = "nios.start_addr"
      end_addr     = "nios.end_addr"
      network_view = "nios.network_view"
    }
  }

  pair_checks = ["nios.address_type", "nios.comment", "nios.disable", "nios.discovery_member", "nios.enable_discovery", "nios.enable_immediate_discovery", "nios.end_addr", "nios.ipv6_end_prefix", "nios.ipv6_prefix_bits", "nios.ipv6_start_prefix", "nios.name", "nios.network", "nios.network_view", "nios.recycle_leases", "nios.same_port_control_discovery_blackout", "nios.server_association_type", "nios.start_addr", "nios.template", "nios.use_blackout_setting", "nios.use_discovery_basic_polling_settings", "nios.use_enable_discovery", "nios.use_logic_filter_rules", "nios.use_recycle_leases", "nios.use_subscribe_settings"]

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
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

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.address_type", "nios.comment", "nios.disable", "nios.discovery_member", "nios.enable_discovery", "nios.enable_immediate_discovery", "nios.end_addr", "nios.ipv6_end_prefix", "nios.ipv6_prefix_bits", "nios.ipv6_start_prefix", "nios.name", "nios.network", "nios.network_view", "nios.recycle_leases", "nios.same_port_control_discovery_blackout", "nios.server_association_type", "nios.start_addr", "nios.template", "nios.use_blackout_setting", "nios.use_discovery_basic_polling_settings", "nios.use_enable_discovery", "nios.use_logic_filter_rules", "nios.use_recycle_leases", "nios.use_subscribe_settings"]

  step {
    nios {
      network      = infoblox_ipv6_network.test.nios.network
      start_addr   = "14::1"
      end_addr     = "14::10"
      network_view = infoblox_network_view.test.nios.name
      ext_attrs    = { Site = "{{random}}" }
    }
  }

}
