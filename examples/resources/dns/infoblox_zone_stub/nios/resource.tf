// Create a DNS Zone Stub with Basic Fields
resource "infoblox_zone_stub" "zone_stub_basic_fields" {
  nios = {
    fqdn = "example_stub_zone.example.com"
    stub_from = [
      {
        name    = "stub.example.com"
        address = "1.1.1.1"
      }
    ]
  }
}

// Create a DNS Zone Stub with Additional Fields
resource "infoblox_zone_stub" "zone_stub_additional_fields" {
  nios = {
    fqdn = "example_stub_zone2.example.com"
    stub_from = [
      {
        name    = "stub.example.com"
        address = "1.1.1.1"
      }
    ]

    // Additional Fields
    ms_ddns_mode = "ANY"
    prefix       = "stub-prefix"
    comment      = "This is a stub zone with additional fields"
    view         = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create an IPV4 DNS Zone Stub
resource "infoblox_zone_stub" "zone_stub_ipv4" {
  nios = {
    fqdn = "10.1.0.0/25"
    stub_from = [
      {
        name    = "stub.example.com"
        address = "1.1.1.1"
      }
    ]
    zone_format = "IPV4"
    prefix      = "zone-stub"
  }
}

// Create an IPV6 DNS Zone Stub
resource "infoblox_zone_stub" "zone_stub_ipv6_mapping" {
  nios = {
    fqdn = "3001:db8::/64"
    stub_from = [
      {
        name    = "stub.example.com"
        address = "1.1.1.1"
      }
    ]
    zone_format = "IPV6"
  }
}
