// Create a Response Policy Zone with Basic Fields
resource "infoblox_zone_rp" "zone_rp_basic_fields" {
  nios = {
    fqdn = "example1.com"
    view = "default"
    ext_attrs = {
      Site = "location-1"
    }
  }
}

// Create a Response Policy Zone that rewrites matching queries to a substitute name
resource "infoblox_zone_rp" "zone_rp_additional_fields" {
  nios = {
    // Basic Fields
    fqdn = "example2.com"
    view = "default"

    // A zone must own its SOA (explicit grid primary) for the SOA timers below
    // to be honoured; otherwise NIOS resets them to the grid defaults.
    grid_primary = [
      {
        name = "infoblox.localdomain"
      }
    ]

    // Response policy behaviour
    rpz_policy      = "SUBSTITUTE"
    substitute_name = "alternate.fqdn"
    rpz_severity    = "WARNING"
    log_rpz         = true

    // SOA timers
    soa_default_ttl  = 37000
    soa_expire       = 92000
    soa_negative_ttl = 900
    soa_refresh      = 2100
    soa_retry        = 800

    comment = "Comment for Zone RP"
    ext_attrs = {
      Site = "location-2"
    }
  }
}

// Create a Response Policy Zone with an explicit grid primary and IP-trigger drop rules
resource "infoblox_zone_rp" "zone_rp_grid_primary" {
  nios = {
    fqdn = "example3.com"
    view = "default"

    grid_primary = [
      {
        name = "infoblox.localdomain"
      }
    ]

    rpz_drop_ip_rule_enabled                = true
    rpz_drop_ip_rule_min_prefix_length_ipv4 = 24
    rpz_drop_ip_rule_min_prefix_length_ipv6 = 64

    ext_attrs = {
      Site = "location-1"
    }
  }
}
