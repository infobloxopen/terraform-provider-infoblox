// Create an auth zone with Basic Fields
resource "infoblox_zone_auth" "example_1" {
  uddi = {
    fqdn         = "example.com."
    primary_type = "cloud"
    comment      = "test comment"
    tags = {
      Site = "location-1"
    }
  }
}

// Create an auth zone served by an external primary
resource "infoblox_zone_auth" "example_external_primary" {
  uddi = {
    fqdn               = "external.example.com."
    primary_type       = "external"
    external_primaries = [{ fqdn = "tf-infoblox.com.", address = "192.168.11.11", type = "primary" }]
  }
}

// Create an auth zone with ACLs
resource "infoblox_zone_auth" "example_with_acls" {
  uddi = {
    fqdn         = "acl.example.com."
    primary_type = "cloud"

    query_acl = [
      { access  = "allow",
        element = "ip",
        address = "192.168.11.11"
      }
    ]
    transfer_acl = [
      { element = "acl",
        acl     = "dns/acl/0d20aafe-8490-4d2c-8367-9bc1b62b601c"
      }
    ]
    update_acl = [
      { access  = "deny",
        element = "tsig_key",
        tsig_key = {
          key = "keys/tsig/24b2fb48-666c-4e95-bc03-da6b5fef26c8"
        }
      }
    ]
  }
}
