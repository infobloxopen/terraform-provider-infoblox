// List Forward NSGs using filters
list "infoblox_forward_nsg" "by_name" {
  provider = infoblox
  config {
    filters = {
      name = "example-forward-nsg"
    }
  }
  limit = 10
}

// List Forward NSGs using tags
list "infoblox_forward_nsg" "by_tag" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Forward NSGs with resource details included
list "infoblox_forward_nsg" "with_resource" {
  provider         = infoblox
  include_resource = true
}
