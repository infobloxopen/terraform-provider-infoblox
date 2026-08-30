// Create an IPv6 Fixed Address with Basic Fields
resource "infoblox_ipv6fixedaddress" "create_ipv6_fixed_address_basic" {
  nios = {
    ipv6addr = "2001:db8:abcd:1231::2"
    duid     = "01:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
  }
}

// Create an IPv6 Fixed Address with Additional Fields with PREFIX address type
resource "infoblox_ipv6fixedaddress" "create_ipv6_fixed_address_additional1" {
  nios = {
    // Basic Fields
    address_type    = "PREFIX"
    ipv6prefix      = "2001:db8:abcd:1232::"
    ipv6prefix_bits = 64
    match_client    = "MAC_ADDRESS"
    mac_address     = "01:6a:7b:8c:9d:5e"
    network_view    = "default"

    // Additional Fields
    comment = "IPv6 Fixed Address created with additional fields"

    options = [
      {
        name  = "domain-name"
        num   = 15
        value = "example.com"
      },
      {
        name  = "dhcp-renewal-time"
        num   = 58
        value = "720"
      }
    ]
    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPv6 Fixed Address with Additional Fields with BOTH address type
resource "infoblox_ipv6fixedaddress" "create_ipv6_fixed_address_additional2" {
  nios = {
    // Basic Fields
    address_type    = "BOTH"
    ipv6addr        = "2001:db8:abcd:1231::3"
    ipv6prefix      = "2001:db8:abcd:1231::"
    ipv6prefix_bits = 64
    match_client    = "MAC_ADDRESS"
    mac_address     = "00:6a:7b:8c:9d:6e"
    network_view    = "default"

    // Additional Fields
    preferred_lifetime  = 2400
    valid_lifetime      = 4800
    domain_name         = "example.com"
    domain_name_servers = ["2001:4860:4860::8888", "2001:4860:4860::8844"]
  }
}

// Create an IPv6 Fixed Address using function call to retrieve ipv6addr
resource "infoblox_ipv6fixedaddress" "create_ipv6_fixed_address_with_func_call" {
  nios = {
    duid    = "00:01:01:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
    comment = "Fixed Address created with ipv6addr retrieved via function call"
  }
}
