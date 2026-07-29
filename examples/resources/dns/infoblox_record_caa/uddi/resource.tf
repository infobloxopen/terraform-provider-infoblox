// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

// Create a CAA Record with Basic Fields
resource "infoblox_record_caa" "record_caa_with_basic_fields" {
  uddi = {
    rdata = {
      tag   = "issue"
      value = "ca.example.com"
    }
    zone    = infoblox_zone_auth.parent_zone.id
    comment = "CAA Record created by Terraform"
  }
}

// Create a CAA Record with Additional Fields
resource "infoblox_record_caa" "record_caa_with_additional_fields" {
  uddi = {
    rdata = {
      flags = 0
      tag   = "issue"
      value = "ca.example.com"
    }
    zone = infoblox_zone_auth.parent_zone.id

    // Additional Fields
    name_in_zone = "caa"
    comment      = "CAA Record created by Terraform"
    disabled     = false
    ttl          = 3600
    tags = {
      Site = "location-1"
    }
  }
}
