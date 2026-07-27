// Create IPV4 forward mapping zone with Basic Fields
resource "infoblox_zone_auth" "create_zone1" {
  nios = {
    fqdn = "example1.com"
    view = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create IPV4 reverse mapping zone with Basic Fields
resource "infoblox_zone_auth" "create_zone2" {
  nios = {
    fqdn        = "10.0.0.0/24"
    view        = "default"
    zone_format = "IPV4"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create IPV6 reverse mapping zone with Basic Fields
resource "infoblox_zone_auth" "create_zone3" {
  nios = {
    fqdn        = "2002:1100::/64"
    view        = "default"
    zone_format = "IPV6"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create IPV4 forward mapping zone with Additional Fields
resource "infoblox_zone_auth" "create_zone4" {
  nios = {
    // Basic Fields
    fqdn = "example2.com"
    view = "default"

    // Additional Fields
    grid_primary = [
      {
        name = "infoblox.localdomain",
      }
    ]
    restart_if_needed = true

    soa_default_ttl  = 37000
    soa_expire       = 92000
    soa_negative_ttl = 900
    soa_refresh      = 2100
    soa_retry        = 800

    allow_query = [
      {
        struct     = "addressac"
        address    = "10.0.0.0"
        permission = "ALLOW"
      }
    ]

    comment = "IPV4 forward auth zone"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

