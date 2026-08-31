resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

resource "infoblox_record_svcb" "example" {
  uddi = {
    rdata = {
      target_name = "example.com"
    }
    zone = infoblox_zone_auth.example.id

    // Other optional fields
    name_in_zone = "svcb"
    comment      = "Example comment"
    disabled     = false
    ttl          = 3600
    tags = {
      Site = "location-1"
    }
  }
}
