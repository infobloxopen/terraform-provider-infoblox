// Prerequisites: a Federated Realm and a parent Federated Block to preview allocations from
resource "infoblox_federated_realm" "example" {
  uddi = {
    name = "example_federated_realm"
  }
}

resource "infoblox_federated_block" "example" {
  uddi = {
    address         = "10.0.0.0"
    cidr            = 16
    federated_realm = infoblox_federated_realm.example.id
    tags = {
      Site = "location-1"
    }
  }
}

// Preview the next available overlapping block addresses under a parent Federated Block.
// This does not allocate anything.
data "infoblox_next_available_overlapping_block" "example_next_available_overlapping_block" {
  id                      = infoblox_federated_block.example.id
  cidr                    = 24
  overlapping_block_count = 3
}

// Preview the next available overlapping block addresses under the Federated Blocks matching a tag
data "infoblox_next_available_overlapping_block" "example_next_available_overlapping_block_by_tag" {
  depends_on              = [infoblox_federated_block.example]
  cidr                    = 24
  overlapping_block_count = 2
  tag_filters = {
    Site = "location-1"
  }
}
