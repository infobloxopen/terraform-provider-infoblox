terraform {
  required_providers {
    infoblox = {
      source  = "infobloxopen/infoblox"
      version = "0.0.1"
    }
  }
}

provider "infoblox" {
  uddi = {
    portal_url = "https://csp.eu-stg-1.eu.stage.test.infoblox.com"
    portal_key = "7613ec68d403b3d72495f61b8be71aa9b51f07dd96d554c7799f2815ebba1f5b"
  }
}

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
