// List specific Join Tokens using filters
list "infoblox_join_token" "list_join_token_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_join_token"
    }
  }
  limit = 10
}

// List specific Join Tokens using Tags
list "infoblox_join_token" "list_join_token_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Join Tokens with resource details included
list "infoblox_join_token" "list_join_token_with_resource" {
  provider         = infoblox
  include_resource = true
}
