// Create an IPv6 Network (Required as Parent)
resource "infoblox_ipv6_network" "parent_network" {
  nios = {
    network      = "2001:db8:abcd:1231::/64"
    network_view = "default"
    comment      = "Parent network for the fixed addresses below"
  }
}

// Create an IPv6 Fixed Address with Basic Fields
resource "infoblox_ipv6_fixed_address" "create_ipv6_fixed_address_basic" {
  nios = {
    ipv6addr = "2001:db8:abcd:1231::2"
    duid     = "01:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
  }
  depends_on = [infoblox_ipv6_network.parent_network]
}

// Create an IPv6 Fixed Address with Additional Fields with PREFIX address type
resource "infoblox_ipv6_fixed_address" "create_ipv6_fixed_address_prefix_type" {
  nios = {
    // Basic Fields
    address_type    = "PREFIX"
    ipv6prefix      = "2001:db8:abcd:1231::"
    ipv6prefix_bits = 64
    match_client    = "MAC_ADDRESS"
    mac_address     = "01:6a:7b:8c:9d:5e"
    network_view    = "default"

    // Additional Fields
    comment = "IPv6 Fixed Address created with additional fields"

    options = [
      {
        name         = "dhcp6.domain-search"
        num          = 24
        value        = "\"example.com\""
        vendor_class = "DHCPv6"
      },
      {
        name         = "dhcp6.sntp-servers"
        num          = 31
        value        = "2001:4860:4860::8888"
        vendor_class = "DHCPv6"
      }
    ]
    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_ipv6_network.parent_network]
}

// Create an IPv6 Fixed Address with BOTH address type
resource "infoblox_ipv6_fixed_address" "create_ipv6_fixed_address_both_type" {
  nios = {
    // Basic Fields
    address_type    = "BOTH"
    ipv6addr        = "2001:db8:abcd:1231::2"
    ipv6prefix      = "2001:db8:abcd:1231::"
    ipv6prefix_bits = 64
    match_client    = "MAC_ADDRESS"
    mac_address     = "01:6a:7b:8c:9d:5e"
    network_view    = "default"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
  depends_on = [infoblox_ipv6_network.parent_network]
}

// Create an IPv6 Fixed Address with a dynamically allocated ipv6addr
resource "infoblox_ipv6_fixed_address" "create_ipv6_fixed_address_with_dynamic_allocation" {
  nios = {
    duid = "00:01:01:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
    dynamic_allocation = {
      network      = infoblox_ipv6_network.parent_network.nios.network
      network_view = "default"
    }
    comment = "Fixed Address created with a dynamically allocated ipv6addr"
  }
}
