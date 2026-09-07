// Create a Named ACL with Basic Fields
resource "infoblox_namedacl" "create_namedacl" {
  uddi = {
    name = "example_namedacl"
  }
}

// Create a Named ACL with Additional Fields
resource "infoblox_namedacl" "create_namedacl_with_additional_fields" {
  uddi = {
    name    = "example_namedacl_advanced"
    comment = "ACL to allow or deny access to specific network resources"

    // ACL entries using different element types
    list = [
      {
        element = "ip"
        access  = "allow"
        address = "192.168.1.0/24"
      },
      {
        element = "ip"
        access  = "deny"
        address = "10.0.0.1"
      },
      {
        element = "any"
        access  = "deny"
      },
    ]

    tags = {
      Site = "location-1"
    }
  }
}
