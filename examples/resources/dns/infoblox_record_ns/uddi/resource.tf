// Create an Auth Zone (Required as Parent)
resource "infoblox_zone_auth" "parent_zone" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

resource "infoblox_record_ns" "example" {
  uddi = {
    rdata = {
      dname = "ns1.${infoblox_zone_auth.parent_zone.uddi.fqdn}"
    }
    zone         = infoblox_zone_auth.parent_zone.id
    name_in_zone = "ns"
    comment      = "Example NS record"
    disabled     = false
    ttl          = 3600
    tags = {
      Site = "location-1"
    }
  }
}
