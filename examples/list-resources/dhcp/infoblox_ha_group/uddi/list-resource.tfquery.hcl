// List specific HA Groups using filters
list "infoblox_ha_group" "list_ha_group_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-ha-group"
    }
  }
  limit = 10
}

// List specific HA Groups using tags
list "infoblox_ha_group" "list_ha_group_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      location = "site-1"
    }
  }
}

// List HA Groups with resource details included
list "infoblox_ha_group" "list_ha_group_with_resource" {
  provider         = infoblox
  include_resource = true
}
