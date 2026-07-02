resource "unified_dns_record_a" "test1" {
  uddi = {
    name = "test-rec-19.example.com"
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
