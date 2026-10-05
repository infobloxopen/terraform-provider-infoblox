// Get internal domain lists filtered by an attribute
data "infoblox_internal_domain_list" "example_by_attribute" {
  filters = {
    name = "example_list"
  }
}

// Get internal domain lists filtered by tag
data "infoblox_internal_domain_list" "example_by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Get all internal domain lists
data "infoblox_internal_domain_list" "example_all" {}
