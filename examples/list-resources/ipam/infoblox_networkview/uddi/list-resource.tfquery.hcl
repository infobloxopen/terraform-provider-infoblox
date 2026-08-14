// List specific Networkviews using filters
list "infoblox_networkview" "list_networkview_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ip_space"
    }
  }
  limit = 10
}

// List specific Networkviews using Tags
list "infoblox_networkview" "list_networkview_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Networkviews with resource details included
list "infoblox_networkview" "list_networkview_with_resource" {
  provider         = infoblox
  include_resource = true
}
