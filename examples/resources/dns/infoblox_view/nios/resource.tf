// Create DNS View with Basic Fields
resource "infoblox_view" "create_view" {
  nios = {
    name = "example_view"
  }
}

// Create DNS View with Additional Fields
resource "infoblox_view" "create_view_with_additional_fields" {
  nios = {
    name         = "example_custom_view"
    comment      = "DNS View"
    network_view = "default"

    // forwarders settings
    forwarders   = ["10.192.81.23"]
    forward_only = true

    // match clients and destinations
    match_destinations = [
      {
        struct     = "addressac"
        address    = "192.168.0.45"
        permission = "ALLOW"
      },
    ]
    match_clients = [
      {
        struct     = "addressac"
        address    = "92.168.0.23"
        permission = "ALLOW"
      },
    ]

    // extensible attributes
    ext_attrs = {
      Site = "location-1"
    }
  }
}
