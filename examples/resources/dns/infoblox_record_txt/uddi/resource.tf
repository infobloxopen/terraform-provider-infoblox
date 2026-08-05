// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

// Create Record TXT
resource "infoblox_record_txt" "example" {
  uddi = {
    name_in_zone = "txt"
    rdata = {
      text = "Example TXT Record"
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
