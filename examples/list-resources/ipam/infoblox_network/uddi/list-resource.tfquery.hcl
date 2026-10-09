// List specific Networks using filters
list "infoblox_network" "list_network_using_filters" {
  provider = infoblox
  config {
    filters = {
      address = "10.0.0.0"
    }
  }
  limit = 10
}

// List specific Networks using Tags
list "infoblox_network" "list_network_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Networks with resource details included
list "infoblox_network" "list_network_with_resource" {
  provider         = infoblox
  include_resource = true
}
