resource "infoblox_record_aaaa" "example" {
  uddi = {
    name_in_zone = "aaaa"
    rdata = {
      address = "2001:db8::1"
    }
    zone = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"

    # Other optional fields
    comment  = "Example comment"
    disabled = false
    ttl      = 3600
    tags = {
      location = "site1"
    }
  }
}
