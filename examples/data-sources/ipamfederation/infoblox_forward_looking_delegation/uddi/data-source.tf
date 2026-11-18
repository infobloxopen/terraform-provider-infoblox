// Retrieve Forward Looking Delegations filtered by an attribute
data "infoblox_forward_looking_delegation" "get_forward_looking_delegation_using_filters" {
  filters = {
    name = "example_forward_looking_delegation"
  }
}

// Retrieve Forward Looking Delegations filtered by tag
data "infoblox_forward_looking_delegation" "get_forward_looking_delegation_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Forward Looking Delegations
data "infoblox_forward_looking_delegation" "get_all_forward_looking_delegations" {}
