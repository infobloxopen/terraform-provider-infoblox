resource "infoblox_record_a" "test1" {
  uddi = {
    name = "record_a.example.com"
    rdata = {
      address = "10.0.0.19"
    }
    comment = "test comment"
    zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    tags = {
      Site = "location-1"
    }
  }
}
