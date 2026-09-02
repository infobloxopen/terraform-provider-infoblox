// Create a Zone Auth (Required as parent)
resource "infoblox_zone_auth" "example_auth" {
  nios = {
    fqdn = "example_auth.com"
    view = "default"
  }
}

// Create a DNS zone delegated with Basic Fields
resource "infoblox_zone_delegated" "zone_delegated_basic_fields" {
  nios = {
    fqdn = "zone-delegated.example_auth.com"
    delegate_to = [
      {
        name    = "ns1.example.com"
        address = "10.10.10.10"
      }
    ]
  }

  depends_on = [infoblox_zone_auth.example_auth]
}

// Create a Zone Auth with IPv4 mapping (Required as parent)
resource "infoblox_zone_auth" "example_auth_reverse" {
  nios = {
    fqdn        = "111.0.0.0/24"
    view        = "default"
    zone_format = "IPV4"
  }
}

// Create a DNS zone delegated with IPv4 mapping
resource "infoblox_zone_delegated" "zone_delegated_ip4_mapping" {
  nios = {
    fqdn = "111.0.0.10/32"
    delegate_to = [
      {
        name    = "ns2.example.com"
        address = "2.2.2.2"
      }
    ]
    zone_format = "IPV4"
  }

  depends_on = [infoblox_zone_auth.example_auth_reverse]
}

// Create a DNS zone delegated with Additional Fields
resource "infoblox_zone_delegated" "zone_delegated_additional_fields" {
  nios = {
    fqdn = "zone-delegated-2.example_auth.com"
    delegate_to = [
      {
        name    = "ns2.example.com"
        address = "20.20.20.20"
      }
    ]

    // Additional Fields
    comment          = "This is a delegated zone for example.com"
    delegated_ttl    = 3600
    ms_ad_integrated = false
    ms_ddns_mode     = "ANY"
    prefix           = "example-prefix"
    zone_format      = "FORWARD"
    view             = "default"

    // Extensible Attributes
    ext_attrs = {
      Site = "location-1"
    }
  }

  depends_on = [infoblox_zone_auth.example_auth]
}
