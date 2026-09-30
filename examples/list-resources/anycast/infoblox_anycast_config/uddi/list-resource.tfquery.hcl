// List specific Anycast Configs using Filters
list "infoblox_anycast_config" "list_anycast_config_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "anycast_config_example"
    }
  }
}

// List specific Anycast Configs using Tags
list "infoblox_anycast_config" "list_anycast_config_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Anycast Configs with resource details included
list "infoblox_anycast_config" "list_anycast_config_with_resource" {
  provider         = infoblox
  include_resource = true
}
