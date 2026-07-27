// Create DNS View with Basic Fields
resource "infoblox_view" "create_view" {
  uddi = {
    name = "example_dns_view"
  }
}

// Create DNS View with Additional Fields
resource "infoblox_view" "create_view_with_additional_fields" {
  uddi = {
    name    = "example_custom_view"
    comment = "An example view"

    // ip_spaces = ["ipam/ip_space/<id>"]

    tags = {
      site = "Site A"
    }

    // match clients
    match_clients_acl = [
      {
        access  = "allow"
        element = "ip"
        address = "192.168.10.10"
      }
    ]

    // custom root name servers
    custom_root_ns_enabled = true
    custom_root_ns = [
      {
        address = "192.168.11.11"
        fqdn    = "example.com."
      }
    ]

    // EDNS client subnet (ECS) settings
    ecs_enabled    = true
    ecs_forwarding = false
    ecs_prefix_v4  = 24
    ecs_prefix_v6  = 56
    ecs_zones = [
      {
        access = "allow"
        fqdn   = "example.com."
      }
    ]

    // recursion access control
    recursion_acl = [
      {
        access  = "allow"
        element = "ip"
        address = "192.168.1.1"
      }
    ]

    // DNSSEC settings
    dnssec_enabled         = true
    dnssec_validate_expiry = true
    dnssec_trust_anchors = [
      {
        algorithm  = 8
        public_key = "AwEAAejpWrcCPGWEoiebhWKSdT6LcMGBsoXadKu1XNthMZUvx3P92HNE4J3q3EtAX8pnTsNShrsDvvgn4hmCsrURMLx/g+76JtLU5pdbtrGFjelHAuMrzLgFzpuA5Ct9THth5Hto6c0rl4yzz3qT3+I/rnUYrL/zd9zKWyMp1A9KlHqwCA3JbFZfl4IKBD2/g+GScEcpnDfUUVDU+7qRZkZ4BhBQ4a6Em73zggz/crcDtwc1cHcRP0DGbekZhF29+yjTPW4zKqGUHW8ZtP49ZMXOTY42epeiddFNy0Ze2jbTg99CnKvAxIKzYInUaPJ04rgMyeuVWpRKsVetJemhCaj9lEs="
        zone       = "example.com."
        sep        = true
      }
    ]

    // zone authority (SOA) settings
    zone_authority = {
      default_ttl       = 28800
      expire            = 2419200
      mname             = "ns.example.com"
      negative_ttl      = 900
      refresh           = 10800
      retry             = 3600
      rname             = "hostmaster@example.com"
      use_default_mname = false
    }
  }
}
