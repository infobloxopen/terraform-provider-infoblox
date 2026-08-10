resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example-rec-mx.com."
    primary_type = "cloud"
  }
}

resource "infoblox_record_mx" "example" {
  uddi = {
    name_in_zone = "mx"
    rdata = {
      exchange   = "m1.example.com"
      preference = 10
    }
    zone = infoblox_zone_auth.example.id
    tags = {
      Site = "location-1"
    }
  }
}
