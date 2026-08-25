// Manage HTTPS Records with basic fields
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

// Manage HTTPS Records with optional fields
resource "infoblox_record_https" "example" {
  uddi = {
    rdata = {
      target_name = "example.com"
    }
    zone = infoblox_zone_auth.example.id

    // Other optional fields
    name_in_zone = "https"
    comment      = "Example comment"
    disabled     = false
    ttl          = 3600
    tags = {
      Site = "location-1"
    }
  }
}
