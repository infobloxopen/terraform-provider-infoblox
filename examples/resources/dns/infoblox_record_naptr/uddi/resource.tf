// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "example.com"
    primary_type = "cloud"
  }
}

// Create Record NAPTR with Basic Fields
resource "infoblox_record_naptr" "example_basic" {
  uddi = {
    name_in_zone = "naptr"
    rdata = {
      order       = 100
      preference  = 10
      replacement = "."
      services    = "SIP+D2U"
    }
    zone = infoblox_zone_auth.example.id
  }
}

// Create Record NAPTR with Additional Fields
resource "infoblox_record_naptr" "example_additional_fields" {
  uddi = {
    name_in_zone = "naptr"
    rdata = {
      order       = 100
      preference  = 10
      replacement = "."
      services    = "SIP+D2U"
      flags       = "U"
      regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
    }
    zone     = infoblox_zone_auth.example.id
    comment  = "Example comment"
    disabled = false
    ttl      = 3600
    tags = {
      Site = "location-1"
    }
  }
}
