// Create a DNS View (Required as Parent)
resource "infoblox_view" "example" {
  uddi = {
    name = "example-view"
  }
}

// Create a Primary Authoritative Zone (Required as Parent)
resource "infoblox_zone_auth" "example" {
  uddi = {
    fqdn         = "domain.com."
    primary_type = "cloud"
    view         = infoblox_view.example.id
  }
}

// Create a DNS Zone Delegated
resource "infoblox_zone_delegated" "example" {
  uddi = {
    fqdn = "del.domain.com."
    delegation_servers = [{
      address = "12.0.0.0"
      fqdn    = "ns1.com."
    }]

    // Other optional fields
    view    = infoblox_view.example.id
    comment = "Delegation zone created through Terraform"
    tags = {
      Site = "location-1"
    }
    disabled = true
  }

  depends_on = [infoblox_view.example, infoblox_zone_auth.example]
}
