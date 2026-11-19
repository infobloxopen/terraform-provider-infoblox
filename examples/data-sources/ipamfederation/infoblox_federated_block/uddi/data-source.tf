// Retrieve Federated Blocks filtered by an attribute
data "infoblox_federated_block" "get_federated_block_using_filters" {
  filters = {
    comment = "Example Federated Block created through Terraform"
  }
}

// Retrieve Federated Blocks filtered by tag
data "infoblox_federated_block" "get_federated_block_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// NOTE: An unfiltered "retrieve all" data source is intentionally omitted here. The
// underlying API's List endpoint returns IPv6 utilization_v6 values in a shape the current
// SDK cannot unmarshal, so a grid with any pre-existing IPv6 Federated Blocks will cause an
// unfiltered list/datasource call to fail with a JSON unmarshal error. Always scope
// FederatedBlock data sources with filters or tag_filters until that SDK issue is fixed.
