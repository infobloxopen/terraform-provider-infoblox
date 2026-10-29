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
      network      = infoblox_ipv6_network.example_network_basic.nios.network
      network_view = "default"
      cidr         = 72
    }
    comment = "Network created with function call"
  }
}

// Create an IPAM IPv6 Network with Function Call from an IPv6 Network selected by Extensible Attributes
resource "infoblox_ipv6_network" "example_func_call_filter_network" {
  nios = {
    dynamic_allocation = {
      filter_params = {
        "*Site" : "location-1"
      }
      cidr = 72
    }
    comment = "Network created with function call from a network selected by extensible attributes"
  }
  depends_on = [infoblox_ipv6_network.example_network_additional]
}

// Create an IPAM IPv6 Network Container for allocation of IPv6 Networks
resource "infoblox_ipv6_network_container" "example_container" {
  nios = {
    network      = "12::/48"
    network_view = "default"
    ext_attrs = {
      Site = "location-2"
    }
  }
}

// Create an IPAM IPv6 Network with Function Call from an IPv6 Network Container
resource "infoblox_ipv6_network" "example_func_call_network_container" {
  nios = {
    dynamic_allocation = {
      network_container = infoblox_ipv6_network_container.example_container.nios.network
      network_view      = "default"
      cidr              = 64
    }
    comment = "Network created with function call from a network container"
  }
}

// Create an IPAM IPv6 Network with Function Call from an IPv6 Network Container selected by Extensible Attributes
resource "infoblox_ipv6_network" "example_func_call_filter_network_container" {
  nios = {
    dynamic_allocation = {
      filter_object = "networkcontainer"
      filter_params = {
        "*Site" : "location-2"
      }
      cidr = 64
    }
    comment = "Network created with function call from a network container selected by extensible attributes"
  }
  depends_on = [infoblox_ipv6_network_container.example_container]
}
