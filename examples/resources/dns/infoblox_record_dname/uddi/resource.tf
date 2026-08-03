// Create an auth zone (required as parent)
resource "infoblox_zone_auth" "parent_zone" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
  }
}

// Create a DNAME record with basic fields
resource "infoblox_record_dname" "example_1" {
  uddi = {
    zone = infoblox_zone_auth.parent_zone.id
    rdata = {
      target = "example-dname-1.com."
    }
  }
}
