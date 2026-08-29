// Retrieve Federated Realms filtered by an attribute
data "infoblox_federated_realm" "get_federated_realm_using_filters" {
  filters = {
    name = "example_federated_realm"
  }
}

// Retrieve Federated Realms filtered by tag
data "infoblox_federated_realm" "get_federated_realm_using_tag_filters" {
  tag_filters = {
    site = "Site A"
  }
}

// Retrieve all Federated Realms
data "infoblox_federated_realm" "get_all_federated_realms" {}
