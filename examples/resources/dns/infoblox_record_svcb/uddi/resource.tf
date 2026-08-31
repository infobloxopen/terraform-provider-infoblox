// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

// Manage an SVCB Record
resource "infoblox_record_svcb" "example" {
  uddi = {
    rdata = {
      target_name = "record.com"
    }
    zone = infoblox_zone_auth.example.id

    // Other optional fields
    name_in_zone = "record"
    comment      = "Example comment"
    disabled     = false
    ttl          = 3600
    tags = {
      Site = "location-1"
    }
  }
}
