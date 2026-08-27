// Objects to be present on the grid
// ensg1 - external_ns_group
// Create a DNS zone forward with Basic Fields
resource "infoblox_zone_forward" "zone_forward_basic_fields" {
  nios = {
    fqdn              = "example1.example.com"
    external_ns_group = "ensg1"
  }
}

// Create a DNS zone forward with Additional Fields
resource "infoblox_zone_forward" "zone_forward_additional_fields" {
  nios = {
    fqdn = "example2.example.com"
    forward_to = [
      {
        name    = "ns1.example.com"
        address = "1.1.1.1"
      }
    ]
    forwarding_servers = [
      {
        name                    = "infoblox.172_28_82_33"
        forwarders_only         = true
        use_override_forwarders = true
        forward_to = [
          {
            name    = "kk.fwd.com"
            address = "10.2.1.31"
          }
        ]
      }
    ]
    view = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPV4 DNS zone forward
resource "infoblox_zone_forward" "zone_forward_ipv4" {
  nios = {
    fqdn = "192.1.0.0/24"
    forward_to = [
      {
        name    = "ns1.example.com"
        address = "1.1.1.1"
      }
    ]
    zone_format = "IPV4"
  }
}

// Create an IPV6 DNS zone forward
resource "infoblox_zone_forward" "zone_forward_ipv6_mapping" {
  nios = {
    fqdn = "3002:db8::/64"
    forward_to = [
      {
        name    = "ns1.example.com"
        address = "1.1.1.1"
      }
    ]
    zone_format = "IPV6"
  }
}
