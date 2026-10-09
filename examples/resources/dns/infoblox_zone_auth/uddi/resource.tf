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

// Create a Named ACL (Required for the transfer ACL)
resource "infoblox_namedacl" "example" {
  uddi = {
    name = "example_namedacl"
  }
}

// Create a TSIG Key (Required for the update ACL)
resource "infoblox_tsig_key" "example" {
  uddi = {
    name   = "tsig-key-example.example.com."
    secret = "wuQuR0A08ApqKT65yaGiqWHalHxS7Ie8LF2VTUFZFZo="
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
        acl     = infoblox_namedacl.example.id
      }
    ]
    update_acl = [
      { access  = "deny",
        element = "tsig_key",
        tsig_key = {
          key = infoblox_tsig_key.example.id
        }
      }
    ]
  }
}
