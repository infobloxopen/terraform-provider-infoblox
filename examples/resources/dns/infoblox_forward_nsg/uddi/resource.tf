// Create a Forward NSG with basic fields
resource "infoblox_forward_nsg" "basic" {
  uddi = {
    name = "example-forward-nsg"
  }
}

// Create a Forward NSG with additional fields
resource "infoblox_forward_nsg" "example" {
  uddi = {
    name    = "example-forward-nsg-full"
    comment = "An example Forward NSG created by Terraform"

    external_forwarders = [
      {
        address = "12.10.2.1"
        fqdn    = "ext.primary.example.com."
      },
    ]

    forwarders_only = false

    // hosts and internal_forwarders accept DNS host resource identifiers
    // hosts              = ["dns/host/<id>"]
    // internal_forwarders = ["dns/host/<id>"]

    // Reference another Forward NSG to chain resolution
    nsgs = [infoblox_forward_nsg.basic.id]

    tags = {
      Site = "location-1"
    }
  }
}
