// List specific Internal Domain Lists using filters
list "infoblox_internal_domain_list" "list_internal_domain_list_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_list"
    }
  }
  limit = 10
}

// List specific Internal Domain Lists using Tags
list "infoblox_internal_domain_list" "list_internal_domain_list_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Internal Domain Lists with resource details included
list "infoblox_internal_domain_list" "list_internal_domain_list_with_resource" {
  provider         = infoblox
  include_resource = true
}
