resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "193.in-addr.arpa."
    primary_type = "cloud"
  }
}

resource "infoblox_record_ptr" "example" {
  uddi = {
    rdata = {
      dname = "example.com"
    }
    zone = infoblox_zone_auth.example.id

    // Other optional fields
    name_in_zone = "1.0.168"
    comment      = "Created by Terraform"
    disabled     = false
    ttl          = 3600
    tags = {
      location = "site1"
    }
  }
}
