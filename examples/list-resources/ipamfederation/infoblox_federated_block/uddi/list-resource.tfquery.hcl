// List specific Federated Blocks using filters
list "infoblox_federated_block" "list_federated_block_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Example Federated Block created through Terraform"
    }
  }
  limit = 10
}

// List specific Federated Blocks using Tags
list "infoblox_federated_block" "list_federated_block_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Federated Blocks with resource details included
//
// NOTE: always scope this with filters or tag_filters (as above), never leave it unfiltered.
// The API's List endpoint returns IPv6 utilization_v6 values in a shape the current SDK
// cannot unmarshal, so an unfiltered call fails with a JSON error on any grid with
// pre-existing IPv6 Federated Blocks.
list "infoblox_federated_block" "list_federated_block_with_resource" {
  provider         = infoblox
  include_resource = true
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}
