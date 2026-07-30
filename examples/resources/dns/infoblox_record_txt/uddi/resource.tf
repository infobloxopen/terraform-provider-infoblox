// Create Record TXT
resource "infoblox_record_txt" "example" {
  uddi = {
    name_in_zone = "txt"
    rdata = {
      text = "example.com"
    }
    zone = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"

    # Other optional fields
    comment  = "Example comment"
    disabled = false
    ttl      = 3600
    tags = {
      location = "site1"
    }
  }
}
