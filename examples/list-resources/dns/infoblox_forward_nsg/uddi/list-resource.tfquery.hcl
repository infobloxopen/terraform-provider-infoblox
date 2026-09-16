// List specific Forward NSGs using filters
list "infoblox_forward_nsg" "list_forward_nsg_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-forward-nsg"
    }
  }
  limit = 10
}

// List specific Forward NSGs using tags
list "infoblox_forward_nsg" "list_forward_nsg_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Forward NSGs with resource details included
list "infoblox_forward_nsg" "list_forward_nsg_with_resource" {
  provider         = infoblox
  include_resource = true
}
