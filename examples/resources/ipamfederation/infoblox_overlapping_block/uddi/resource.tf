// Create a Federated Realm (required as a parent)
resource "infoblox_federated_realm" "example" {
  uddi = {
    name = "example_test_federated_realm"
  }
}

// Create an Overlapping Block with basic fields
resource "infoblox_overlapping_block" "example" {
  uddi = {
    address         = "10.10.0.0"
    cidr            = 24
    federated_realm = infoblox_federated_realm.example.id
  }
}

// Create an Overlapping Block with additional fields
resource "infoblox_overlapping_block" "example_with_additional_fields" {
  uddi = {
    address         = "10.20.0.0"
    cidr            = 24
    federated_realm = infoblox_federated_realm.example.id

    # Other optional fields
    name    = "example_overlapping_block"
    comment = "Example Overlapping Block created through Terraform"
    tags = {
      Site = "location-1"
    }
  }
}

// Create an IPv6 Overlapping Block
resource "infoblox_overlapping_block" "example_ipv6" {
  uddi = {
    address         = "2001:DB8:a1::0"
    cidr            = 48
    federated_realm = infoblox_federated_realm.example.id
    name            = "example_overlapping_block_ipv6"
  }
}
