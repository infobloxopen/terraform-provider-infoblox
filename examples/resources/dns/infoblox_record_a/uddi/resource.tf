// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

resource "infoblox_record_a" "test1" {
  uddi = {
    name_in_zone = "record_a"
    rdata = {
      address = "10.0.0.19"
    }
    comment = "test comment"
    zone    = infoblox_zone_auth.example.id
    tags = {
      Site = "location-1"
    }
  }
}
