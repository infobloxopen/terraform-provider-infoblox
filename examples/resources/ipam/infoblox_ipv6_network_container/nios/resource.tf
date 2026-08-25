// Create IPv6 Network Container with Basic Fields
resource "infoblox_ipv6_network_container" "ipv6networkcontainer_with_basic_fields" {
  nios = {
    network = "10::/64"
  }
}

// Create IPv6 Network Container with Additional Fields
resource "infoblox_ipv6_network_container" "ipv6networkcontainer_with_additional_fields" {
  nios = {
    // Required attributes
    network = "11::/64"

    // Basic configuration
    network_view = "default"
    comment      = "IPv6 network container with additional fields"

    options = [
      {
        name         = "dhcp6.fqdn",
        num          = 39,
        value        = "test_options.com",
        vendor_class = "DHCPv6"
      }
    ]
    // DDNS settings
    enable_ddns            = true
    ddns_domainname        = "example.com"
    ddns_generate_hostname = true
    ddns_ttl               = 3600

    // Extensible attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create IPv6 Network Container with Function Call
resource "infoblox_ipv6_network_container" "example_func_call" {
  nios = {
    dynamic_allocation = {
      network      = "10::/64"
      network_view = "default"
      cidr         = 72
    }
    comment = "Network container created with function call"
  }
  depends_on = [
    infoblox_ipv6_network_container.ipv6networkcontainer_with_basic_fields
  ]
}
