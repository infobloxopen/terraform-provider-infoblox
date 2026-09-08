// Create a DNS View (Required as Parent)
resource "infoblox_view" "example_view" {
  uddi = {
    name = "example_dns_view1"
  }
}

// Create a Forward NSG (Prerequisite)
resource "infoblox_forward_nsg" "example_forward_nsg" {
  uddi = {
    name = "example-forward-nsg"
  }
}

// Objects to be present on the grid
// dns host, internal forwarders
// Create a DNS Zone Forward
resource "infoblox_zone_forward" "example" {
  uddi = {
    fqdn = "domain.com."

    // Other optional fields
    comment = "Example of a Forward Zone"
    tags = {
      Site = "location-1"
    }
    nsgs                = [infoblox_forward_nsg.example_forward_nsg.id]
    hosts               = ["dns/host/<id>"]
    internal_forwarders = ["dns/host/<id>"]
    view                = infoblox_view.example_view.id
  }
}
