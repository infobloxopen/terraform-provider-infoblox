// Prerequisite: a Federated Realm to allocate the block under
resource "infoblox_federated_realm" "example" {
  uddi = {
    name = "example_federated_realm"
  }
}

// Create a Federated Block with basic fields
resource "infoblox_federated_block" "example" {
  uddi = {
    address         = "10.0.0.0"
    cidr            = 16
    federated_realm = infoblox_federated_realm.example.id
  }
}

// Create a Federated Block with additional fields
resource "infoblox_federated_block" "example_with_additional_fields" {
  uddi = {
    address         = "10.1.0.0"
    cidr            = 16
    federated_realm = infoblox_federated_realm.example.id

    # Other optional fields
    name    = "example_federated_block_full"
    comment = "Example Federated Block created through Terraform"
    tags = {
      Site = "location-1"
    }
    network_compliance = {
      minimum_netmask_length = 20
      maximum_netmask_length = 28
      default_netmask_length = 24
    }
  }
}

// Create a Federated Block by allocating the next available block under an existing parent
// Federated Block. The parent is referenced by attribute to create an implicit dependency.
resource "infoblox_federated_block" "example_next_available" {
  uddi = {
    cidr            = 20
    federated_realm = infoblox_federated_realm.example.id
    dynamic_allocation = {
      next_available_id = infoblox_federated_block.example.id
    }
  }
}
