// List specific Ranges using filters
list "infoblox_range" "list_range_using_filters" {
  provider = infoblox
  config {
    filters = {
      start = "192.168.1.15"
    }
  }
  limit = 10
}

// List specific Ranges using Tags
list "infoblox_range" "list_range_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ranges with resource details included
list "infoblox_range" "list_range_with_resource" {
  provider         = infoblox
  include_resource = true
}
