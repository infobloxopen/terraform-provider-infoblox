resource "infoblox_record_ns" "example" {
  uddi = {
    rdata = {
      dname = "ns1.example.com."
    }
    zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    name_in_zone = "ns"
    comment      = "Example NS record"
    disabled     = false
    ttl          = 3600
    tags = {
      Site = "location-1"
    }
  }
}
