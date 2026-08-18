// Retrieve Forward NSGs filtered by an attribute
data "infoblox_forward_nsg" "by_name" {
  filters = {
    name = "example-forward-nsg"
  }
}

// Retrieve Forward NSGs filtered by tag
data "infoblox_forward_nsg" "by_tag" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Forward NSGs
data "infoblox_forward_nsg" "all" {}
