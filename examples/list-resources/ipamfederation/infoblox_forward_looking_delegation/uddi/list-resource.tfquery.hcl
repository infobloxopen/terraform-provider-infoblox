// List specific Forward Looking Delegations using filters
list "infoblox_forward_looking_delegation" "list_forward_looking_delegation_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Forward Looking Delegations using Tags
list "infoblox_forward_looking_delegation" "list_forward_looking_delegation_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Forward Looking Delegations with resource details included
list "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource" {
  provider         = infoblox
  include_resource = true
}
