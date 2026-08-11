// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example-auth-zone.com."
    primary_type = "cloud"
  }
}

// Create Record SRV
resource "infoblox_record_srv" "test1" {
  uddi = {
    name_in_zone = "record_srv"
    rdata = {
      port     = 5060
      priority = 10
      target   = "sip.example.com"
      weight   = 5
    }
    comment = "test comment"
    ttl     = 300
    zone    = infoblox_zone_auth.example.id
    tags = {
      Site = "location-1"
    }
  }
}
