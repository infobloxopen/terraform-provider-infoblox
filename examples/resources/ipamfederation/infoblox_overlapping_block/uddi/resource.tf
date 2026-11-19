// Prerequisite: a Federated Realm to associate the overlapping block with
resource "infoblox_federated_realm" "example" {
  uddi = {
    name = "example_federated_realm"
  }
}

// Create an Overlapping Block with basic fields
resource "infoblox_overlapping_block" "example" {
  uddi = {
    address         = "10.0.0.0"
    cidr            = 24
    federated_realm = infoblox_federated_realm.example.id
  }
}

// Create an Overlapping Block with additional fields
resource "infoblox_overlapping_block" "example_with_additional_fields" {
  uddi = {
    address         = "10.1.0.0"
    cidr            = 24
    federated_realm = infoblox_federated_realm.example.id

    # Other optional fields
    name    = "example_overlapping_block_full"
    comment = "Example Overlapping Block created through Terraform"
    tags = {
      Site = "location-1"
    }
  }
}

// Prerequisite for the next-available example: a parent Federated Block
resource "infoblox_federated_block" "example_parent" {
  uddi = {
    address         = "10.2.0.0"
    cidr            = 16
    federated_realm = infoblox_federated_realm.example.id
  }
}

// Create an Overlapping Block by allocating the next available block under a parent Federated
// Block. The parent is referenced by attribute to create an implicit dependency.
resource "infoblox_overlapping_block" "example_next_available" {
  uddi = {
    cidr            = 24
    federated_realm = infoblox_federated_realm.example.id
    dynamic_allocation = {
      next_available_id = infoblox_federated_block.example_parent.id
    }
  }
}
