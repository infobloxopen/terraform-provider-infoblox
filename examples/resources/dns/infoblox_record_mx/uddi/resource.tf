// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example-rec-mx.com."
    primary_type = "cloud"
  }
}

// Create Record MX
resource "infoblox_record_mx" "example" {
  uddi = {
    name_in_zone = "mx"
    rdata = {
      exchange   = "m1.example.com"
      preference = 10
    }
    zone = infoblox_zone_auth.example.id

    # Other optional fields
    comment  = "MX record created by Terraform"
    disabled = false
    ttl      = 3600
    tags = {
      Site = "location-1"
    }
  }
}
