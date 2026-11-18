// Create a Federated Realm (required as a parent)
resource "infoblox_federated_realm" "example" {
  uddi = {
    name = "example_federated_realm"
  }
}

// Create a Forward Looking Delegation with basic fields
resource "infoblox_forward_looking_delegation" "example" {
  uddi = {
    address          = "10.10.0.0"
    cidr             = 24
    federated_realms = [infoblox_federated_realm.example.id]
  }
}

// Create a Forward Looking Delegation with additional fields
resource "infoblox_forward_looking_delegation" "example_with_additional_fields" {
  uddi = {
    address          = "10.20.0.0"
    cidr             = 24
    federated_realms = [infoblox_federated_realm.example.id]

    # Other optional fields
    name    = "example_forward_looking_delegation"
    comment = "Example Forward Looking Delegation created through Terraform"
    tags = {
      Site = "location-1"
    }
  }
}

// Create an IPv6 Forward Looking Delegation
resource "infoblox_forward_looking_delegation" "example_ipv6" {
  uddi = {
    address          = "2001:db8:a1::"
    cidr             = 48
    federated_realms = [infoblox_federated_realm.example.id]
    name             = "example_forward_looking_delegation_ipv6"
  }
}
