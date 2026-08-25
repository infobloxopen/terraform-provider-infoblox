// Create an IPAM IPv6 Network with Basic Fields
resource "infoblox_ipv6_network" "example_network_basic" {
  nios = {
    network = "10::/64"
  }
}

// Create an IPAM IPv6 Network with Additional Fields
resource "infoblox_ipv6_network" "example_network_additional" {
  nios = {
    // Required attributes
    network = "11::/64"

    // Basic configuration
    network_view = "default"
    comment      = "Created by Terraform"

    ddns_enable_option_fqdn    = true
    ddns_generate_hostname     = true
    ddns_server_always_updates = true
    ddns_ttl                   = 0
    disable                    = true

    enable_ddns             = true
    enable_ifmap_publishing = true
    ext_attrs = {
      Site = "location-1"
    }

    options = [
      {
        name         = "dhcp6.fqdn",
        num          = 39,
        value        = "test_options.com",
        vendor_class = "DHCPv6"
      }
    ]
    port_control_blackout_setting = {
      enable_blackout = false
    }
    preferred_lifetime          = 27000
    valid_lifetime              = 43200
    recycle_leases              = true
    update_dns_on_lease_renewal = true
  }
}

// Create an IPAM IPv6 Network with Function Call
resource "infoblox_ipv6_network" "example_func_call" {
  nios = {
    dynamic_allocation = {
      network      = "10::/64"
      network_view = "default"
      cidr         = 72
    }
    comment = "Network created with function call"
  }
  depends_on = [
    infoblox_ipv6_network.example_network_basic
  ]
}
