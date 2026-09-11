// Retrieve specific Named Lists by filters
data "infoblox_named_list" "get_named_list_using_filters" {
  filters = {
    type = "custom_list"
  }
}

// Retrieve specific Named Lists using Tags
data "infoblox_named_list" "get_named_list_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Named Lists
data "infoblox_named_list" "get_all_named_lists" {}
