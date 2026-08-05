// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

resource "infoblox_record_aaaa" "example" {
  uddi = {
    name_in_zone = "aaaa"
    rdata = {
      address = "2001:db8::1"
    }
    zone = infoblox_zone_auth.parent_zone.id

    # Other optional fields
    comment  = "Example comment"
    disabled = false
    ttl      = 3600
    tags = {
      location = "site1"
    }
  }
}
