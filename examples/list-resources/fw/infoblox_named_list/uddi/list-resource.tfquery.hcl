// List specific Named Lists using filters
list "infoblox_named_list" "list_named_list_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Named Lists using Tags
list "infoblox_named_list" "list_named_list_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Named Lists with resource details included
list "infoblox_named_list" "list_named_list_with_resource" {
  provider         = infoblox
  include_resource = true
}
