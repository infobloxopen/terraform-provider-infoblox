// Create Record CNAME
resource "infoblox_record_cname" "example" {
  uddi = {
    name_in_zone = "cname"
    rdata = {
      cname = "example.com"
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
