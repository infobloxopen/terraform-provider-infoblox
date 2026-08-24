// Retrieve a specific Forward NSG using filters
data "infoblox_forward_nsg" "get_forward_nsg_using_filters" {
  filters = {
    name = "example-forward-nsg"
  }
}

// Retrieve specific Forward NSGs using tags
data "infoblox_forward_nsg" "get_forward_nsg_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Forward NSGs
data "infoblox_forward_nsg" "all" {}
