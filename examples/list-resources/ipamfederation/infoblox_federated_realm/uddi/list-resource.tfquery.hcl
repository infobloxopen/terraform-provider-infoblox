// List specific Federated Realms using filters
list "infoblox_federated_realm" "list_federated_realm_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_federated_realm"
    }
  }
  limit = 10
}

// List specific Federated Realms using tags
list "infoblox_federated_realm" "list_federated_realm_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      site = "Site A"
    }
  }
}

// List Federated Realms with resource details included
list "infoblox_federated_realm" "list_federated_realm_with_resource" {
  provider         = infoblox
  include_resource = true
}
