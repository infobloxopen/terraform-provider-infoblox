// Create a Federated Realm with the minimum set of fields
resource "infoblox_federated_realm" "example" {
  uddi = {
    name = "example_federated_realm"
  }
}

// Create a Federated Realm with additional fields
resource "infoblox_federated_realm" "example_with_additional_fields" {
  uddi = {
    name = "example_federated_realm_full"

    # Other optional fields
    comment = "Example Federated Realm created through Terraform"
    tags = {
      site = "Site A"
    }
  }
}
