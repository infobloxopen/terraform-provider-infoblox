// Create a Named List with Basic Fields
resource "infoblox_named_list" "create_named_list" {
  uddi = {
    name = "example_named_list"
    type = "custom_list"
    items_described = [
      {
        item        = "example.com"
        description = "Example domain"
      },
    ]
  }
}

// Create a Named List with Additional Fields
resource "infoblox_named_list" "create_named_list_with_additional_fields" {
  uddi = {
    name             = "example_named_list_advanced"
    type             = "custom_list"
    description      = "Named list for blocking malicious domains"
    confidence_level = "HIGH"
    threat_level     = "MEDIUM"

    items_described = [
      {
        item        = "malicious-example.com"
        description = "Known malicious domain"
      },
      {
        item        = "192.0.2.0/24"
        description = "Suspicious IP range"
      },
    ]

    tags = {
      Site = "location-1"
    }
  }
}
