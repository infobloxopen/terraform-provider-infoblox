resource "infoblox_record_srv" "test1" {
  uddi = {
    name_in_zone = "record_srv.example.com"
    rdata = {
      port     = 5060
      priority = 10
      target   = "sip.example.com"
      weight   = 5
    }
    comment = "test comment"
    ttl     = 300
    zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    tags = {
      Site = "location-1"
    }
  }
}
